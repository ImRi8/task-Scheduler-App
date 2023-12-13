package Service

import (
	"Task-scheduler-App/Constant"
	"Task-scheduler-App/Datalayer/Entity"
	"Task-scheduler-App/Datalayer/SqlDB"
	"Task-scheduler-App/models"
	"gofr.dev/pkg/gofr"
	"net/http"
	"strconv"
)

type TaskService struct {
}

var (
	sqlDb = &SqlDB.TaskDbService{}
)

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func convertStringtoInt64(id string) int64 {
	result, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// Handle the error if the conversion fails.
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

	return models.TaskResponseMsg(Constant.SUCCESS, "ALL OK!!", http.StatusOK, task)

}

func (taskService *TaskService) validateGetTaskById(id string) models.Response {
	if id == "" {
		return models.ResponseMsg(Constant.FAILURE, "Valid id not found", http.StatusOK)
	}
	if !isNumeric(id) {
		return models.ResponseMsg(Constant.FAILURE, "Id can only be a numeric value", http.StatusOK)
	}
	return models.ResponseMsg(Constant.SUCCESS, "All ok", http.StatusOK)
}
