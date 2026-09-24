# Temporal workflow alignment

## Objective

Align feed fetching, feed adding, task execution, scheduling, status, retries,
and worker operations with Temporal's durable-execution model while preserving
the existing feed behavior and task queues.

The target design is:

```text
GraphQL/API or Schedule
  -> start the owning Workflow on its owning Task Queue
  -> Workflow orchestrates Activities and, where justified, Child Workflows
  -> caller receives Workflow ID and Run ID without waiting unless it asks for a result
```

## Async task-queue intent

`feed_tasks` is the application's attempt to provide the same developer
experience as Resque or Sidekiq:

```text
web request -> enqueue work -> return a job ID
```

That intent is valid and should remain a first-class design goal. Temporal's
equivalent of enqueueing a durable background job is starting a Workflow with a
Temporal Client. `ExecuteWorkflow` persists the Workflow Execution and returns
a handle containing the Workflow ID and Run ID; the web request should return
those identifiers without calling `Get` unless it explicitly wants to wait for
completion.

The important distinction is that a Temporal Task Queue is worker routing
infrastructure, not the public job abstraction. The application-facing task
dispatcher may map a task type to a typed Workflow and Task Queue, but it should
start the owning Workflow directly. A proxy Workflow whose Activity starts
another Workflow adds an execution, obscures the real job ID, and makes
completion semantics harder to reason about.

The intended mapping is:

```text
enqueue add feed     -> start AddFeedWorkflow on feed-add-queue
enqueue feed refresh -> start FetchFeedsWorkflow on feed-fetch-queue
enqueue generation   -> start GenerateFeedsWorkflow on feed-tasks-queue
```

Each Workflow may contain one Activity for a simple job or several Activities
and Child Workflows for a real business process. This preserves the easy
"please do this now" API while using Temporal for durable execution, retries,
results, cancellation, visibility, and recovery.

Temporal references:

- [Go workflow basics](https://docs.temporal.io/develop/go/workflows/basics)
- [Go child workflows](https://docs.temporal.io/develop/go/workflows/child-workflows)
- [Go error handling](https://docs.temporal.io/develop/go/best-practices/error-handling)
- [Go Temporal Client](https://docs.temporal.io/develop/go/client/temporal-client)
- [Go Schedules](https://docs.temporal.io/develop/go/workflows/schedules)
- [Go Continue-As-New](https://docs.temporal.io/develop/go/workflows/continue-as-new)

## Scope

In scope:

- `internal/temporal/feed_fetch`
- `internal/temporal/feed_add`
- `internal/temporal/feed_tasks`
- `internal/tasks`
- GraphQL workflow start/status integration
- Temporal worker setup, task queues, schedules, and Helm/Tilt configuration
- Workflow/activity tests and operational documentation

Out of scope:

- Replacing the existing gRPC feed service or GraphQL API
- Changing feed/article business rules beyond what is required for durable
  execution and idempotency
- Introducing a separate job database unless Temporal visibility and workflow
  results cannot satisfy a demonstrated product requirement

## Current gaps to preserve as implementation context

- `internal/tasks/temporal.go` is the Sidekiq-like enqueue/dispatch boundary;
  its new path starts owning workflows directly and returns their handles.
- The legacy `feed_tasks` proxy workflows and activities remain registered for
  compatibility and still need a retirement/migration plan.
- Typed entrypoints use versioned Temporal workflow names (`*WorkflowV2`);
  original positional workflow names remain registered for in-flight history.
- `feed_tasks.RefreshFeeds` starts `FetchFeedsWorkflow` from an Activity but does
  not wait for that workflow to finish.
- `feed_tasks.AddFeed` starts the real add workflow from an Activity and waits for
  it, creating two workflow executions for one add request.
- `FetchFeedsWorkflow` now waits for all started children and returns a
  best-effort per-feed outcome.
- Workflow logging and the content-not-changed signal now use Temporal-safe
  workflow logging and typed Application Errors.
- Most policies use `MaximumAttempts: 1`; several workflows rely on implicit
  Activity retry behavior.
- Fetch and save activities do not expose progress heartbeats or per-run object
  identities.
- Task status is now read with `DescribeWorkflowExecution`, but result retrieval
  and authorization are still open.
- `SaveResults` now affects the feed workflow outcome when partial persistence
  occurs.
- `WorkflowID` configuration is defined but not consistently applied.
- Schedule startup now reconciles an existing schedule instead of deleting it.
- The task worker now uses the shared worker observability/interceptor setup.

## 1. Establish workflow ownership and typed contracts

Owner: Temporal/application architecture

- [x] Define the owning workflow and task queue for every operation:
      `AddFeedWorkflow -> feed-add-queue`,
      `FetchFeedsWorkflow -> feed-fetch-queue`, and
      `GenerateFeedsWorkflow -> feed-tasks-queue`.
- [x] Decide that `GenerateFeedsWorkflow` is an enqueue-only bulk operation:
      it creates add-feed jobs and does not wait for each independent add
      workflow to complete.
- [x] Decide and document whether each public operation is asynchronous or
      synchronous. Default API behavior should be asynchronous and return a
      workflow handle.
- [x] Replace `[]any` task arguments with serializable input structs for
      workflows that accept inputs, including `AddFeedInput` and
      `GenerateFeedsInput`; the fetch workflow intentionally has no input.
- [x] Replace `TaskResult.Result any` and `TaskStatusResult.Status any` with
      typed result/status fields.
- [ ] Centralize workflow type names and task queue names so client and worker
      registration cannot silently diverge.
- [x] Preserve compatibility for already-running workflow types while changing
      new client call paths.

Acceptance criteria:

- [ ] Every workflow has one documented owner, one primary task queue, typed
      serializable inputs, and a typed result contract.
- [ ] The API contract clearly states whether it returns a workflow handle or a
      completed business result.

## 2. Preserve the enqueue API while starting real workflows directly

Owner: `internal/tasks`, GraphQL, workflow entrypoints

- [x] Keep an application-facing task dispatcher/service, but make it a thin
      typed Temporal Client adapter rather than a worker-side proxy workflow.
- [x] Change the dispatcher to start the owning workflow directly on its
      owning queue instead of starting a `feed_tasks` proxy workflow.
- [x] Keep the task-type-to-workflow mapping in one place and make it explicit
      that this is a job-enqueue registry, not business orchestration.
- [x] Document that returning `WorkflowID`/`RunID` is the successful enqueue
      response, analogous to returning a Sidekiq/Resque job ID.
- [ ] Remove the `feed_tasks.AddFeed` Activity that starts
      `feed_add.AddFeedWorkflow`, unless a separate orchestration requirement
      proves the proxy is needed.
- [ ] Remove the `feed_tasks.RefreshFeeds` Activity that starts
      `FetchFeedsWorkflow`, unless it is deliberately changed to await and
      propagate the child workflow result.
- [ ] If a parent workflow must own another workflow, replace client-based
      `ExecuteWorkflow` calls with `workflow.ExecuteChildWorkflow` and configure
      `ChildWorkflowOptions` explicitly.
- [x] For asynchronous callers, return immediately after workflow start with
      `WorkflowID` and `RunID`; do not call `Get` in the start path.
- [ ] For synchronous callers, call `Get` with a typed result pointer and
      propagate the workflow error.
- [x] Ensure GraphQL `addFeed(url, name)` either passes `name` into the owning
      workflow or explicitly removes it from the public contract.
- [x] Decide that `GenerateFeedsWorkflow` merely enqueues each generated add;
      test and document the selected semantics.

Acceptance criteria:

- [x] One user feed-add request produces one owning add workflow execution.
- [x] Treat refresh as an explicit fire-and-forget enqueue operation whose API
      response returns the owning fetch workflow ID.
- [x] API requests do not remain open while an asynchronous workflow performs
      remote fetches or database writes.

## 3. Implement reliable workflow status and results

Owner: `internal/tasks`, GraphQL

- [x] Replace history scraping in `Temporal.Status` with
      `DescribeWorkflowExecution`.
- [x] Return the workflow ID in `TaskStatusResult` and GraphQL
      `FeedTaskStatusResponse`.
- [x] Map Temporal execution states to a stable application enum or documented
      string set: `RUNNING`, `COMPLETED`, `FAILED`, `CANCELED`, `TIMED_OUT`, and
      `TERMINATED`.
- [ ] Preserve failure type/message in an internal result while avoiding
      leaking sensitive Activity details through GraphQL.
- [x] Add a typed `GetResult` path using `GetWorkflow(...).Get(ctx, &result)`.
- [ ] Add status/result tests for running, completed, failed, canceled, and
      unknown workflow IDs.
- [ ] Add authorization checks so a user cannot inspect another user's add
      workflow by ID.

Acceptance criteria:

- [x] `feedTaskStatus` returns a meaningful state while a workflow is running.
- [x] Completed workflow results and failures are retrievable without scanning
      event history manually.
- [x] The returned task ID is the ID of the workflow whose status is being
      reported.

## 4. Correct fetch fan-out and lifecycle behavior

Owner: `internal/temporal/feed_fetch`

- [x] Replace standard `slog` calls in workflow code with
      `workflow.GetLogger(ctx)`.
- [x] Decide whether fetch batches are best-effort or fail-fast.
- [x] For best-effort processing, wait for all started children and return a
      serializable per-feed outcome containing feed ID, stage, error type, and
      retryability.
- [ ] For fail-fast processing, explicitly cancel or terminate sibling child
      workflows and test that behavior.
- [x] Add explicit `ChildWorkflowOptions`, including task queue, workflow ID
      strategy, retry policy if applicable, and parent close policy.
- [x] Use stable child IDs that include the fetch execution/run identity and
      feed ID, unless duplicate suppression by feed ID is an intentional
      requirement.
- [x] Ensure the parent waits for `ChildWorkflowExecutionStarted` when using
      asynchronous child behavior.
- [x] Document the fixed five-minute parent await as the current batch SLA;
      replace it with a batch-size policy when feed pagination is introduced.
- [ ] Page or batch feed IDs rather than returning an unbounded feed slice into
      workflow history.
- [ ] Add Continue-As-New when event history or batch size approaches the
      selected operational limit.
- [ ] Add tests for multiple children, one child failure, all children failing,
      parent timeout, cancellation, and retry/replay behavior.

Acceptance criteria:

- [x] A single feed failure does not unexpectedly terminate unrelated feed work.
- [ ] Parent and child lifecycle behavior is explicit and covered by tests.
- [ ] Fetch behavior remains bounded as the number of feeds grows.

## 5. Make Activities durable, idempotent, and observable

Owner: feed fetch/add Activity implementations

- [ ] Define retry classification for HTTP, gRPC, storage, validation, rate
      limit, duplicate, and not-found errors.
- [x] Replace plain sentinel errors crossing Activity boundaries with typed
      non-retryable Temporal application errors or serializable result statuses.
- [ ] Set Activity options deliberately per operation, including
      `StartToCloseTimeout`, `ScheduleToCloseTimeout`, retry policy, and
      `HeartbeatTimeout` where progress can be reported.
- [ ] Add heartbeats to long-running article-save or large-feed operations,
      including useful progress details such as article count and last URL.
- [ ] Make article writes idempotent under Activity retry and worker crash
      recovery.
- [x] Decide whether partial article failures fail the Activity. If partial
      success is allowed, return and persist the complete `SaveResults` value
      and make the workflow outcome reflect it.
- [ ] Add run-specific or content-addressed object keys for raw RSS and parsed
      article artifacts, or enforce a documented single-writer policy.
- [ ] Bound remote feed size, parsing work, and article count before consuming
      unbounded resources.
- [x] Use `activity.GetLogger(ctx)` consistently in Activities and include
      workflow/run/feed identifiers in structured fields.
- [ ] Add Activity tests for retryable errors, non-retryable errors, timeout,
      cancellation, heartbeat progress, duplicate writes, and partial saves.

Acceptance criteria:

- [ ] Retrying any Activity cannot create duplicate feed/article records.
- [ ] A worker crash resumes from a durable checkpoint or safely repeats an
      idempotent operation.
- [ ] Operators can distinguish transient failures from invalid feed content.

## 6. Make schedules safe and operationally owned

Owner: schedule provisioning and deployment configuration

- [ ] Separate schedule provisioning from the long-running fetch worker.
- [x] Implement idempotent create-or-update behavior instead of deleting the
      schedule before creation.
- [x] Set an explicit schedule overlap policy, probably `Skip` or `BufferOne`,
      based on the desired behavior when a fetch exceeds one interval.
- [ ] Set an explicit scheduled workflow ID strategy and reuse policy.
- [ ] Add schedule description, pause, trigger, and update commands or an
      equivalent deployment job.
- [ ] Reconcile code/configuration names: `SCHEDULE_ID`, `WORKFLOW_ID`,
      `FEED_SCHEDULE_ID`, and `FEED_WORKFLOW_ID`.
- [ ] Decide whether `cmd/feed-fetch` remains an ad-hoc launcher or is split
      into a worker binary and a schedule-admin command.
- [x] Add tests proving that running schedule
      provisioning twice does not remove or duplicate the schedule.

Acceptance criteria:

- [ ] A worker restart cannot silently remove the production fetch schedule.
- [ ] At most the intended number of fetch workflows run concurrently.
- [ ] Schedule ownership and recovery steps are documented.

## 7. Standardize worker setup and deployment

Owner: worker platform and Helm/Tilt

- [x] Make `feed-tasks-worker` use the shared Temporal client, metrics, OTEL,
      tracing interceptor, and shutdown helper used by fetch/add workers.
- [ ] Create one worker bootstrap path that validates task queue and Temporal
      connection configuration consistently.
- [ ] Configure worker concurrency deliberately for fetch, add, and task queues;
      do not rely only on the fetch workflow semaphore.
- [ ] Add readiness/health behavior that distinguishes process health from
      successful Temporal task-queue polling where practical.
- [ ] Add task queue backlog, Activity failure, workflow failure, and schedule
      overlap metrics/alerts.
- [ ] Ensure Temporal client creation happens once per process and clients are
      closed during shutdown.
- [ ] Document namespace, TLS, credentials, and server persistence requirements
      for non-local environments.
- [ ] Replace the Kubernetes `temporal server start-dev` deployment with a
      supported persistent Temporal deployment or explicitly mark the chart as
      local/demo-only.
- [ ] Add Helm/Tilt checks for all three task queues and required environment
      variables.

Acceptance criteria:

- [ ] All worker types emit consistent traces and Temporal metrics.
- [ ] Scaling a worker deployment changes capacity without changing workflow
      semantics.
- [ ] Non-local deployments do not depend on an ephemeral development server.

## 8. Testing, migration, and rollout

Owner: application and release engineering

- [ ] Add deterministic workflow tests for every changed workflow contract.
- [ ] Add integration coverage that starts each workflow on its real task queue
      and verifies the owning worker completes it.
- [ ] Add tests that kill/restart a worker during an Activity and verify resume
      or safe retry behavior.
- [ ] Add tests for duplicate workflow starts using the selected Workflow ID
      reuse/conflict policy.
- [ ] Add a replay test or workflow-history compatibility test before changing
      existing workflow definitions.
- [ ] Verify GraphQL add-feed, task-start, task-status, and result behavior
      against the local Temporal deployment.
- [ ] Verify schedule create/update/trigger/pause behavior.
- [ ] Run focused checks:
      `go test ./internal/temporal/... ./internal/tasks ./internal/graph`.
- [ ] Run repository checks:
      `mise run go-test`, `mise run go-lint`, `mise run buf-lint`, and relevant
      Helm/Tilt validation.
- [ ] Update `docs/architecture.md`, `docs/services.md`, and
      `docs/testing.md` with the final workflow ownership and operational model.
- [ ] Add a rollback plan for workflow type changes, task queue changes, and
      schedule changes.
- [ ] Keep this specification open until all behavior, operations, and release
      checks are complete.

## Completion gate

- [ ] Feed add starts exactly one owning workflow and reports its real ID.
- [ ] Feed fetch has explicit fan-out, child lifecycle, failure, retry, and
      batch-history behavior.
- [ ] Task status and result APIs accurately describe the underlying workflow.
- [ ] Activity retry/idempotency/heartbeat behavior is documented and tested.
- [ ] Schedule provisioning is idempotent and has an explicit overlap policy.
- [ ] All workers use consistent observability and shutdown behavior.
- [ ] Local and non-local Temporal deployment assumptions are documented.
- [ ] Focused, repository, integration, and deployment checks pass.
