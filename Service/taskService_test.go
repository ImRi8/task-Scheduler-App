package Service

import (
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

func ParseStringToTime1(dateTime string) *time.Time {
	parsedDate, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return nil
	}
	utc := parsedDate.Local()
	return &utc
}

func ParseStringToTime2(dateTime string) *time.Time {
	parsedDate, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return nil
	}
	utc := parsedDate.UTC()
	return &utc
}

//func TestTaskService_GetTaskById(t *testing.T) {
//	ctx := intializeCtx(t)
//	taskService := TaskService{}
//	//taskService.GetTaskById("30", ctx)
//	response := []struct {
//		id       string
//		response interface{}
//	}{
//		{"30", models.Response{
//			Status:     "FAILURE",
//			Message:    "Entry for this id not found",
//			HttpStatus: 200,
//		}}, {"31",
//			models.Response{
//				Status:     "FAILURE",
//				Message:    "Entry for this id not found",
//				HttpStatus: 200,
//			}},
//	}
//
//	for _, tc := range response {
//		resp := taskService.GetTaskById(tc.id, ctx)
//
//		if resp == tc.response {
//			ctx.Logger.Infof("Error Occured")
//		}
//	}
//}

//func TestTaskService_DeleteTaskById(t *testing.T) {
//	ctx := intializeCtx(t)
//	taskService := TaskService{}
//	//taskService.GetTaskById("30", ctx)
//	response := []struct {
//		id       string
//		response interface{}
//	}{
//		{"30", models.Response{
//			Status:     "SUCCESS",
//			Message:    "Delete ALL OK!!",
//			HttpStatus: 200,
//		}}, {"31",
//			models.Response{
//				Status:     "SUCCESS",
//				Message:    "Delete ALL OK!!",
//				HttpStatus: 200,
//			}},
//	}
//
//	for _, tc := range response {
//		resp := taskService.DeleteTaskById(tc.id, ctx)
//
//		if resp == tc.response {
//			ctx.Logger.Infof("Error Occured")
//		}
//	}
//}

//func TestTaskService_CreateTaskById(t *testing.T) {
//	ctx := intializeCtx(t)
//	taskService := TaskService{}
//
//	response := []struct {
//		input    models.Request
//		response models.Response
//	}{{models.Request{
//		Title:       "New title",
//		Description: "Descp of new title",
//		Priority:    2,
//		DueDate:     "2023-12-19T15:04:08Z",
//	}, models.Response{
//		Status:      "SUCCESS",
//		Message:     "Task Creation OK!!",
//		HttpStatus:  200,
//		ID:          18,
//		Title:       "newtask1",
//		Description: "sleep1 at",
//		Priority:    2,
//		DueDate:     ParseStringToTime1("2023-12-19T15:04:08Z"),
//	}}, {models.Request{
//		Title:       "New title",
//		Description: "Descp of new title",
//		Priority:    2,
//		DueDate:     "2023-12-19T15:04:08Z",
//	}, models.Response{
//		Status:      "SUCCESS",
//		Message:     "Task Creation OK!!",
//		HttpStatus:  200,
//		ID:          18,
//		Title:       "newtask1",
//		Description: "sleep1 at",
//		Priority:    2,
//		DueDate:     ParseStringToTime1("2023-12-19T15:04:08Z"),
//	}}, {models.Request{
//		Title:       "New title1",
//		Description: "Descp of new title1",
//		Priority:    2,
//		DueDate:     "2023-12-19T15:04:08Z",
//	}, models.Response{
//		Status:      "SUCCESS",
//		Message:     "Task Creation OK!!",
//		HttpStatus:  200,
//		ID:          18,
//		Title:       "newtask1",
//		Description: "sleep1 at",
//		Priority:    2,
//		DueDate:     ParseStringToTime1("2023-12-19T15:04:08Z"),
//	}}}
//
//	for _, tc := range response {
//		resp, err := taskService.CreateTaskById(tc.input, ctx)
//
//		if resp == tc.response {
//			ctx.Logger.Infof("Error Occured")
//		}
//
//		if err != nil {
//			ctx.Logger.Infof("Error occured")
//		}
//	}
//
//}

//func TestTaskService_UpdateTaskById(t *testing.T) {
//	ctx := intializeCtx(t)
//	taskService := TaskService{}
//
//	response := []struct {
//		id       string
//		input    models.Request
//		response models.Response
//	}{{"1", models.Request{
//		Title:       "New update title",
//		Description: "Descp of new title",
//	}, models.Response{
//		Status:      "SUCCESS",
//		Message:     "Task Creation OK!!",
//		HttpStatus:  200,
//		ID:          18,
//		Title:       "newtask1",
//		Description: "sleep1 at",
//		Priority:    2,
//		DueDate:     ParseStringToTime1("2023-12-19T15:04:08Z"),
//	}}, {"2", models.Request{
//		Title:       "New title",
//		Description: "Descp of new update",
//		Priority:    2,
//		DueDate:     "2023-12-19T15:04:08Z",
//	}, models.Response{
//		Status:      "SUCCESS",
//		Message:     "Task Creation OK!!",
//		HttpStatus:  200,
//		ID:          18,
//		Title:       "newtask1",
//		Description: "sleep1 at",
//		Priority:    2,
//		DueDate:     ParseStringToTime1("2023-12-19T15:04:08Z"),
//	}}, {"3", models.Request{
//		Title:       "New title1",
//		Description: "Descp of new title1",
//		Priority:    2,
//		DueDate:     "2023-12-19T15:04:08Z",
//	}, models.Response{
//		Status:      "SUCCESS",
//		Message:     "Task Creation OK!!",
//		HttpStatus:  200,
//		ID:          18,
//		Title:       "newtask1",
//		Description: "sleep1 at",
//		Priority:    2,
//		DueDate:     ParseStringToTime1("2023-12-19T15:04:08Z"),
//	}}}
//
//	for _, tc := range response {
//		resp := taskService.UpdateTaskById(tc.input, ctx)
//
//		if resp != tc.response {
//			ctx.Logger.Infof("Error Occured")
//		}
//	}
//
//}

///////////////////////////////////////Test????????????????????????????????????????????????????????????///////////

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
				ID:          73,
				Title:       "New title",
				Description: "Descp of new title",
				Priority:    2,
				DueDate:     ParseStringToTime2("2023-12-19T15:04:08Z"),
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
				DueDate:     ParseStringToTime1("2023-12-19T20:34:08+05:30"),
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
