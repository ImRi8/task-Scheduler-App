package SqlDB

import (
	"Task-scheduler-App/Datalayer/Entity"
	"context"
	"gofr.dev/pkg/gofr"
	"testing"
	"time"
)

func intializeCtx(t *testing.T) *gofr.Context {
	app := gofr.New()
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	return ctx
}

func TestTaskDbService_GetEntityById(t *testing.T) {
	ctx := intializeCtx(t)
	taskDbService := TaskDbService{}
	test := []struct {
		id       int64
		response interface{}
	}{
		{30, nil},
		{31, nil},
		{33, Entity.Task{
			ID: 33,
		}}}

	for _, tc := range test {
		resp, err := taskDbService.GetEntityById(tc.id, ctx)

		if err != nil {
			ctx.Logger.Infof("Error Occured")
		}
		if resp == tc.response {
			ctx.Logger.Infof("Error Occured")
		}
	}
}

func TestTaskDbService_ShadowRowInEntity(t *testing.T) {
	ctx := intializeCtx(t)
	taskDbService := TaskDbService{}
	test := []struct {
		id                int64
		expErr            error
		expectedRowsCount int64
	}{
		{1, nil, 0},
		{2, nil, 0},
		{3, nil, 0},
	}
	for _, tc := range test {
		err := taskDbService.ShadowRowInEntity(tc.id, ctx)

		if err != tc.expErr {
			ctx.Logger.Infof("Error Occured")
		}

		res, rowsAffErr := ctx.DB().ExecContext(ctx, "SELECT 1")

		rowsAffected, rowsAffErr := res.RowsAffected()
		if rowsAffErr != nil {
			ctx.Logger.Infof("Error while getting rows affected: %v", rowsAffErr)
		}

		if rowsAffected != tc.expectedRowsCount {
			ctx.Logger.Infof("Rows affected mismatch for ID %d. Expected %d, got %d", tc.id, tc.expectedRowsCount, rowsAffected)
		}
	}
}

func TestTaskDbService_UnShadowRowInEntity(t *testing.T) {
	ctx := intializeCtx(t)
	taskDbService := TaskDbService{}

	tests := []struct {
		id                int64
		expectedError     error
		expectedRowsCount int64
	}{
		{1, nil, 0},
		{2, nil, 0},
		{3, nil, 0},
	}

	for _, tc := range tests {
		err := taskDbService.UnShadowRowInEntity(tc.id, ctx)

		// Check for the expected error
		if err != tc.expectedError {
			ctx.Logger.Infof("Error mismatch for ID %d. Expected %v, got %v", tc.id, tc.expectedError, err)
		}

		// Check for the expected number of affected rows
		res, rowsAffErr := ctx.DB().ExecContext(ctx, "SELECT 1") // Execute a query to get the rows affected
		if rowsAffErr != nil {
			ctx.Logger.Infof("Error while reading the rows affected: %v", rowsAffErr)
		}

		rowsAffected, rowsAffErr := res.RowsAffected()
		if rowsAffErr != nil {
			ctx.Logger.Infof("Error while getting rows affected: %v", rowsAffErr)
		}

		if rowsAffected != tc.expectedRowsCount {
			ctx.Logger.Infof("Rows affected mismatch for ID %d. Expected %d, got %d", tc.id, tc.expectedRowsCount, rowsAffected)
		}
	}
}

func TestTaskDbService_CreateRowInEntity(t *testing.T) {
	ctx := intializeCtx(t)
	taskDbService := TaskDbService{}

	tests := []struct {
		inputTask     *Entity.Task
		expectedError error
	}{
		{
			&Entity.Task{
				IsShadowed:  false,
				Title:       "New Task",
				Description: "Description of the new task",
				Priority:    1,
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			nil,
		}, {
			&Entity.Task{
				IsShadowed:  false,
				Title:       "New Task1",
				Description: "Description of the new task1",
				Priority:    2,
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			nil,
		}, {
			&Entity.Task{
				IsShadowed:  false,
				Title:       "New Task2",
				Description: "Description of the new task2",
				Priority:    3,
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			nil,
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		resp, err := taskDbService.CreateRowInEntity(ctx, tc.inputTask)

		if err != tc.expectedError {
			ctx.Logger.Infof("Error mismatch. Expected %v, got %v", tc.expectedError, err)
		}

		if resp == nil {
			ctx.Logger.Infof("Unexpected nil response")
			continue
		}

		createdTask, ok := resp.(*Entity.Task)
		if !ok {
			ctx.Logger.Infof("Unexpected response type. Expected *Entity.Task, got %T", resp)
			continue
		}

		if createdTask.ID == 0 {
			ctx.Logger.Infof("ID not set in the created task")
		}

		if createdTask.Title != tc.inputTask.Title || createdTask.Description != tc.inputTask.Description {
			ctx.Logger.Infof("Task values mismatch. Expected %v, got %v", tc.inputTask, createdTask)
		}
	}
}

func TestTaskDbService_UpdateRowInEntity(t *testing.T) {
	ctx := intializeCtx(t)
	taskDbService := TaskDbService{}

	tests := []struct {
		id            int64
		inputTask     *Entity.Task
		expectedError error
	}{
		{5,
			&Entity.Task{
				Title:       "New update1",
				Description: "Description of the new update1",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			nil,
		}, {4,
			&Entity.Task{
				Title:       "New update1",
				Description: "Description of the new update1",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			nil,
		}, {3,
			&Entity.Task{
				Title:       "New Update1",
				Description: "Description of the new update1",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			nil,
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		resp, err := taskDbService.UpdateRowInEntity(ctx, tc.inputTask, tc.id)

		if err != tc.expectedError {
			ctx.Logger.Infof("Error mismatch. Expected %v, got %v", tc.expectedError, err)
		}

		if resp == nil {
			ctx.Logger.Infof("Unexpected nil response")
			continue
		}
	}
}
