package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"
	sdk "go.temporal.io/sdk/client"
	sdkmocks "go.temporal.io/sdk/mocks"
)

func TestRunScheduleCreatesMissingSchedule(t *testing.T) {
	scheduleClient := &sdkmocks.ScheduleClient{}
	handle := &sdkmocks.ScheduleHandle{}
	handle.On("Describe", mock.Anything).Return(nil, &serviceerror.NotFound{})
	handle.On("GetID").Return("fetch-schedule")
	scheduleClient.On("GetHandle", mock.Anything, "fetch-schedule").Return(handle)
	scheduleClient.On("Create", mock.Anything, mock.MatchedBy(func(options sdk.ScheduleOptions) bool {
		return options.ID == "fetch-schedule" && options.Overlap == enums.SCHEDULE_OVERLAP_POLICY_SKIP
	})).Return(handle, nil)

	client := &sdkmocks.Client{}
	client.On("ScheduleClient").Return(scheduleClient)

	err := runSchedule(context.Background(), &Config{
		ScheduleID: "fetch-schedule",
		WorkflowID: "fetch-workflow",
		TaskQueue:  "feed-fetch-queue",
	}, client)
	require.NoError(t, err)
	client.AssertExpectations(t)
	scheduleClient.AssertExpectations(t)
	handle.AssertExpectations(t)
}

func TestRunScheduleUpdatesExistingScheduleWithoutDeletingIt(t *testing.T) {
	scheduleClient := &sdkmocks.ScheduleClient{}
	handle := &sdkmocks.ScheduleHandle{}
	handle.On("Describe", mock.Anything).Return(&sdk.ScheduleDescription{
		Schedule: sdk.Schedule{
			State: &sdk.ScheduleState{Paused: true},
		},
	}, nil)
	handle.On("Update", mock.Anything, mock.Anything).Return(nil)
	scheduleClient.On("GetHandle", mock.Anything, "fetch-schedule").Return(handle)

	client := &sdkmocks.Client{}
	client.On("ScheduleClient").Return(scheduleClient)

	err := runSchedule(context.Background(), &Config{
		ScheduleID: "fetch-schedule",
		WorkflowID: "fetch-workflow",
		TaskQueue:  "feed-fetch-queue",
	}, client)
	require.NoError(t, err)
	client.AssertExpectations(t)
	scheduleClient.AssertExpectations(t)
	handle.AssertExpectations(t)
}
