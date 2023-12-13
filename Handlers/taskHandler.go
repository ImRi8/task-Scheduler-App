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
