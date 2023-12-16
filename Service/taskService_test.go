package Service

import (
	"Task-scheduler-App/Constant"
	"Task-scheduler-App/models"
	"context"
	"github.com/stretchr/testify/assert"
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

func parseStringToTimeLocal(dateTime string) *time.Time {
	parsedDate, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return nil
	}
	local := parsedDate.Local()
	return &local
}

func parseStringToTimeUtc(dateTime string) *time.Time {
	parsedDate, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return nil
	}
	utc := parsedDate.UTC()
	return &utc
}

func TestTaskService_GetTaskById(t *testing.T) {
	tests := []struct {
		desc     string
		id       string
		expected models.Response
	}{
		{
			desc: "get task by id",
			id:   "30",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Entry for this id not found",
				HttpStatus: 200,
			},
		},
		{
			desc: "get task by non-existent id",
			id:   "31",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Entry for this id not found",
				HttpStatus: 200,
			},
		}, {
			desc: "testing for empty value",
			id:   "",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Valid id not found",
				HttpStatus: 200,
			},
		}, {
			desc: "testing for non numeric value",
			id:   "s",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Id can only be a numeric value",
				HttpStatus: 200,
			},
		},
	}

	ctx := intializeCtx(t)
	taskService := TaskService{}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			resp := taskService.GetTaskById(tc.id, ctx)
			assert.Equal(t, tc.expected, resp, "Test case failed: %s", tc.desc)
		})
	}
}

func TestTaskService_DeleteTaskById(t *testing.T) {
	tests := []struct {
		desc     string
		id       string
		expected models.Response
	}{
		{
			desc: "delete task by id",
			id:   "30",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Id does not exist or already deleted",
				HttpStatus: 200,
			},
		},
		{
			desc: "delete task by non-existent id",
			id:   "31",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Id does not exist or already deleted",
				HttpStatus: 200,
			},
		}, {
			desc: "testing for empty value",
			id:   "",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Valid id not found",
				HttpStatus: 200,
			},
		}, {
			desc: "testing for non numeric value",
			id:   "s",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Id can only be a numeric value",
				HttpStatus: 200,
			},
		}, {
			desc: "testing for non exisiting id",
			id:   "1000",
			expected: models.Response{
				Status:     "FAILURE",
				Message:    "Id does not exist or already deleted",
				HttpStatus: 200,
			},
		},
	}

	ctx := intializeCtx(t)
	taskService := TaskService{}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			resp := taskService.DeleteTaskById(tc.id, ctx)
			assert.Equal(t, tc.expected, resp, "Test case failed: %s", tc.desc)
		})
	}
}

func TestTaskService_CreateTaskById(t *testing.T) {
	tests := []struct {
		desc     string
		input    models.Request
		expected models.Response
	}{
		{
			desc: "create task with valid input",
			input: models.Request{
				Title:       "New title",
				Description: "Descp of new title",
				Priority:    2,
				DueDate:     "2023-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:      "SUCCESS",
				Message:     "Task Creation OK!!",
				HttpStatus:  200,
				ID:          107,
				Title:       "New title",
				Description: "Descp of new title",
				Priority:    2,
				DueDate:     parseStringToTimeUtc("2023-12-19T15:04:08Z"),
			},
		},
		{
			desc: "create task for empty title",
			input: models.Request{
				Title:       "",
				Description: "Descp of new title",
				Priority:    2,
				DueDate:     "2023-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Title Cannot be Empty",
				HttpStatus: 200,
			},
		}, {
			desc: "create task for empty title",
			input: models.Request{
				Title:       "test 123",
				Description: "Descp of new title",
				Priority:    8,
				DueDate:     "2023-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Priority Set",
				HttpStatus: 200,
			},
		},
		{
			desc: "create task for empty title",
			input: models.Request{
				Title:       "test 123",
				Description: "Descp of new title",
				Priority:    8,
				DueDate:     "",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Due-Date",
				HttpStatus: 200,
			},
		}, {
			desc: "create task for empty title",
			input: models.Request{
				Title:       "test 123",
				Description: "Descp of new title",
				Priority:    0,
				DueDate:     "2023-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Priority Set",
				HttpStatus: 200,
			},
		}, {
			desc: "create task with valid input",
			input: models.Request{
				Title:       "New title",
				Description: "Descp of new title",
				Priority:    2,
				DueDate:     "2022-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Due-Date",
				HttpStatus: 200,
			},
		},
		// Add more test cases as needed
	}

	ctx := intializeCtx(t)
	taskService := TaskService{}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			resp, err := taskService.CreateTaskById(tc.input, ctx)

			assert.Equal(t, tc.expected, resp, "Test case failed: %s", tc.desc)
			assert.NoError(t, err, "Unexpected error for test case: %s", tc.desc)
		})
	}
}

func TestTaskService_UpdateTaskById(t *testing.T) {
	tests := []struct {
		desc     string
		id       string
		input    models.Request
		expected models.Response
	}{
		{
			desc: "update task with valid input",
			id:   "1",
			input: models.Request{
				ID:          18,
				Title:       "New update title",
				Description: "Descp of new title",
			},
			expected: models.Response{
				Status:      "SUCCESS",
				Message:     "Task Updation OK!!",
				HttpStatus:  200,
				ID:          18,
				Title:       "New update title",
				Description: "Descp of new title",
				Priority:    2,
				DueDate:     parseStringToTimeLocal("2023-12-19T20:34:08+05:30"),
			},
		}, {
			desc: "create task with valid input",
			id:   "1",
			input: models.Request{
				ID:          1,
				Title:       "New title",
				Description: "Descp of new title",
				Priority:    2,
				DueDate:     "2022-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Due-Date",
				HttpStatus: 200,
			},
		}, {
			desc: "create task for empty title",
			input: models.Request{
				ID:          1,
				Title:       "test 123",
				Description: "Descp of new title",
				Priority:    8,
				DueDate:     "2023-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Priority Set",
				HttpStatus: 200,
			},
		},
		{
			desc: "create task for empty title",
			id:   "1",
			input: models.Request{
				ID:          1,
				Title:       "test 123",
				Description: "Descp of new title",
				Priority:    0,
				DueDate:     "2023-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Priority Set",
				HttpStatus: 200,
			},
		},
		{
			desc: "create task for empty title",
			id:   "",
			input: models.Request{
				ID:          1,
				Title:       "test 123",
				Description: "Descp of new title",
				Priority:    0,
				DueDate:     "2023-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Priority Set",
				HttpStatus: 200,
			},
		},
		{
			desc: "create task for empty title",
			id:   "0",
			input: models.Request{
				ID:          0,
				Title:       "test 123",
				Description: "Descp of new title",
				Priority:    0,
				DueDate:     "2023-12-19T15:04:08Z",
			},
			expected: models.Response{
				Status:     Constant.FAILURE,
				Message:    "Invalid Priority Set",
				HttpStatus: 200,
			},
		},

		// Add more test cases as needed
	}

	ctx := intializeCtx(t)
	taskService := TaskService{}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			resp := taskService.UpdateTaskById(tc.input, ctx)

			assert.Equal(t, tc.expected, resp, "Test case failed: %s", tc.desc)
		})
	}
}
