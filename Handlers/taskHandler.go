package Handlers

import (
	"Task-scheduler-App/Constant"
	"Task-scheduler-App/Service"
	"Task-scheduler-App/models"
	"gofr.dev/pkg/gofr"
	"net/http"
)

var (
	taskService = &Service.TaskService{}
)

func GetTaskByIdHandler(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")

	if id == "" {
		return models.ResponseMsg(Constant.FAILURE, "Valid id not found", http.StatusOK), nil
	}
	return taskService.GetTaskById(id, ctx), nil
}

func DeleteTaskByIdHandler(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")

	if id == "" {
		return models.ResponseMsg(Constant.FAILURE, "Valid id not found", http.StatusOK), nil
	}

	return taskService.DeleteTaskById(id, ctx), nil
}

func CreateTaskHandler(ctx *gofr.Context) (interface{}, error) {
	var taskHttpRequest models.Request
	bindErr := ctx.Bind(&taskHttpRequest)
	if bindErr != nil {
		return models.ResponseMsg(Constant.FAILURE, "Bad request", http.StatusOK), nil
	}

	resp, _ := taskService.CreateTaskById(taskHttpRequest, ctx)

	return resp, nil
}

func UpdateTaskHandler(ctx *gofr.Context) (interface{}, error) {
	var taskHttpRequest models.Request
	bindErr := ctx.Bind(&taskHttpRequest)
	if bindErr != nil {
		return models.ResponseMsg(Constant.FAILURE, "Bad request", http.StatusOK), nil
	}
	resp := taskService.UpdateTaskById(taskHttpRequest, ctx)

	return resp, nil

}
