/*
The graph api puts together multiple parts of the back end. Due to using different frameworks and middleware, not all of the error messages are consistent.
The UI shouldn't have to handle validation erros differently based on which layer it came from.

TODO: graph API errors are inconsistent at the moment. consolidate!
- protobuf validation errors
- service layer validation errors
*/

import { ErrorResponse } from "../types/errors";

interface ValidationFn {
  (errors: string[]): void;
}

interface ErrorFn {
  (error: string): void;
}

/**
 * Handle a common error response from the GraphQL API.
 * @param err
 * @param validationFn Callback to handle validation errors
 * @param errorFn Callback to handle an unknown error
 * @returns
 */
export function handle(
  err: unknown,
  validationFn: ValidationFn,
  errorFn: ErrorFn,
) {
  if ((err as ErrorResponse).response?.errors) {
    const errs: string[] = [];
    for (const error of (err as ErrorResponse).response.errors) {
      errs.push(error.message);
    }
    validationFn(errs);
    return;
  }
  errorFn(`Error: ${err}`); // TODO: don't show raw errors to users
}
