package Service

import (
	"Task-scheduler-App/Constant"
	"Task-scheduler-App/Datalayer/Entity"
	"Task-scheduler-App/Datalayer/SqlDB"
	"Task-scheduler-App/models"
	"fmt"
	"gofr.dev/pkg/gofr"
	"net/http"
	"strconv"
	"time"
)

type TaskService struct {
}

var (
	sqlDb  = &SqlDB.TaskDbService{}
	loc, _ = time.LoadLocation(Constant.IndiaTimeZone)
)

const (
	dateTimeFormat = "2006-12-02T15:04:05Z"
)

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func DateTimeToTime(dateTime string) *time.Time {
	parsedDate, err := time.ParseInLocation(dateTimeFormat, dateTime, time.UTC)
	if err != nil {
		fmt.Println("error while parsing datetime", err.Error())
		return nil
	}

	return &parsedDate
}

func ParseStringToTime(dateTime string) (*time.Time, error) {
	parsedDate, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return nil, err
	}
	return &parsedDate, nil
}

func convertStringtoInt64(id string) int64 {
	result, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

func (taskService *TaskService) GetTaskById(id string, ctx *gofr.Context) models.Response {
	valResp := taskService.validateGetTaskById(id)
	if valResp.Status == Constant.FAILURE {
		return valResp
	}
	intId := convertStringtoInt64(id)
	resp, err := sqlDb.GetEntityById(intId, ctx)

	if err != nil {
		return models.ResponseMsg(Constant.FAILURE, "Entry for this id not found", http.StatusOK)
	}

	if resp == nil {
		return models.ResponseMsg(Constant.FAILURE, "Entry for this id not found", http.StatusOK)
	}

	task := resp.(Entity.Task)

	return models.TaskResponseMsg(Constant.SUCCESS, "Get ALL OK!!", http.StatusOK, task)

}

func (taskService *TaskService) validateGetTaskById(id string) models.Response {
	if id == "" {
		return models.ResponseMsg(Constant.FAILURE, "Valid id not found", http.StatusOK)
	}
	if !isNumeric(id) {
		return models.ResponseMsg(Constant.FAILURE, "Id can only be a numeric value", http.StatusOK)
	}
	return models.ResponseMsg(Constant.SUCCESS, "Validate Get Task All ok", http.StatusOK)
}

func (taskService *TaskService) DeleteTaskById(id string, ctx *gofr.Context) models.Response {
	valResp := taskService.validateDeleteTaskById(id)
	if valResp.Status == Constant.FAILURE {
		return valResp
	}
	intId := convertStringtoInt64(id)
	resp, getErr := sqlDb.GetEntityById(intId, ctx)

	if getErr != nil || resp == nil {
		return models.ResponseMsg(Constant.FAILURE, "Id does not exist or already deleted", http.StatusOK)
	}

	err := sqlDb.ShadowRowInEntity(intId, ctx)

	if err != nil {
		return models.ResponseMsg(Constant.FAILURE, "Entry for this id not found", http.StatusOK)
	}

	return models.ResponseMsg(Constant.SUCCESS, "Delete ALL OK!!", http.StatusOK)

}

func (taskService *TaskService) validateDeleteTaskById(id string) models.Response {
	if id == "" {
		return models.ResponseMsg(Constant.FAILURE, "Valid Id not found", http.StatusOK)
	}

	if !isNumeric(id) {
		return models.ResponseMsg(Constant.FAILURE, "Id can only be a numeric value", http.StatusOK)
	}

	return models.ResponseMsg(Constant.SUCCESS, "Validate Delete Task All OK!!", http.StatusOK)
}

func (taskService *TaskService) CreateTaskById(taskRequest models.Request, ctx *gofr.Context) (models.Response, error) {
	valResp := taskService.validateCreateTaskById(taskRequest)
	if valResp.Status == Constant.FAILURE {
		return valResp, nil
	}

	task := taskService.convertRequestToTaskEntity(taskRequest)

	resp, err := sqlDb.CreateRowInEntity(ctx, &task)

	if err != nil {
		ctx.Logger.Error("Error in the CreateTaskById", err.Error())
		return models.ResponseMsg(Constant.FAILURE, "Internal Server Error", http.StatusOK), nil
	}

	if resp == nil {
		ctx.Logger.Error("Unable to create entry", err.Error())
		return models.ResponseMsg(Constant.FAILURE, "Internal Server Error in resp", http.StatusOK), nil
	}

	resTask := resp.(*Entity.Task)

	return models.TaskResponseMsg(Constant.SUCCESS, "Task Creation OK!!", http.StatusOK, *resTask), nil
}

func (taskService *TaskService) convertRequestToTaskEntity(taskRequest models.Request) Entity.Task {
	dateTime, _ := ParseStringToTime(taskRequest.DueDate)
	task := Entity.Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Priority:    taskRequest.Priority,
		DueDate:     *dateTime,
		IsShadowed:  false,
	}
	return task
}
func (taskService *TaskService) validateCreateTaskById(taskRequest models.Request) models.Response {

	currentTime := time.Now()

	if taskRequest.Title == "" {
		return models.ResponseMsg(Constant.FAILURE, "Title Cannot be Empty", http.StatusOK)
	}
	if taskRequest.Priority <= 1 || taskRequest.Priority > 5 {
		return models.ResponseMsg(Constant.FAILURE, "Invalid Priority Set", http.StatusOK)
	}

	dateTime, _ := ParseStringToTime(taskRequest.DueDate)

	if dateTime == nil || currentTime.After(*dateTime) {
		return models.ResponseMsg(Constant.FAILURE, "Invalid Due-Date", http.StatusOK)
	}

	return models.ResponseMsg(Constant.SUCCESS, "Validate Task All Ok", http.StatusOK)
}

func (taskService *TaskService) UpdateTaskById(taskRequest models.Request, ctx *gofr.Context) models.Response {
	valResp := taskService.validateUpdateTaskById(taskRequest)
	if valResp.Status == Constant.FAILURE {
		return valResp
	}

	resp, err := sqlDb.GetEntityById(taskRequest.ID, ctx)

	if err != nil {
		return models.ResponseMsg(Constant.FAILURE, "Entry for this id not found", http.StatusOK)
	}

	if resp == nil {
		return models.ResponseMsg(Constant.FAILURE, "Entry for this id not found", http.StatusOK)
	}

	task := resp.(Entity.Task)

	taskService.createUpdateTaskEntity(taskRequest, &task)

	updResp, updErr := sqlDb.UpdateRowInEntity(ctx, &task, taskRequest.ID)

	if updErr != nil {
		ctx.Logger.Error("Error in the updateTaskById", updErr.Error())
		return models.ResponseMsg(Constant.FAILURE, "Internal Server Error", http.StatusOK)
	}

	if updResp == nil {
		ctx.Logger.Error("Unable to update entry")
		return models.ResponseMsg(Constant.FAILURE, "Internal Server Error in resp", http.StatusOK)
	}

	resTask := updResp.(*Entity.Task)

	return models.TaskResponseMsg(Constant.SUCCESS, "Task Updation OK!!", http.StatusOK, *resTask)

}

func (taskService *TaskService) createUpdateTaskEntity(taskRequest models.Request, task *Entity.Task) *Entity.Task {
	//fmt.Println("task id ", task.ID)
	if taskRequest.Title != "" {
		task.Title = taskRequest.Title
	}

	if taskRequest.DueDate != "" {
		dueTime, _ := ParseStringToTime(taskRequest.DueDate)
		if dueTime != nil {
			task.DueDate = *dueTime
		}
	}

	if taskRequest.Description != "" {
		task.Description = taskRequest.Description
	}

	if taskRequest.Priority != 0 {
		task.Priority = taskRequest.Priority
	}

	return task
}

func (taskService *TaskService) validateUpdateTaskById(taskRequest models.Request) models.Response {
	if taskRequest.ID == 0 {
		return models.ResponseMsg(Constant.FAILURE, "Valid id not found", http.StatusOK)
	}
	if taskRequest.Priority != 0 {
		if taskRequest.Priority <= 1 || taskRequest.Priority > 5 {
			return models.ResponseMsg(Constant.FAILURE, "Invalid Priority Set", http.StatusOK)
		}
	}
	if taskRequest.DueDate != "" {
		currentTime := time.Now()
		dateTime, _ := ParseStringToTime(taskRequest.DueDate)
		if dateTime == nil || currentTime.After(*dateTime) {
			return models.ResponseMsg(Constant.FAILURE, "Invalid Due-Date", http.StatusOK)
		}
	}

	return models.ResponseMsg(Constant.SUCCESS, "Validate Task All Ok", http.StatusOK)
}
