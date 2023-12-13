package models

import (
	"Task-scheduler-App/Datalayer/Entity"
	"time"
)

type Response struct {
	Status      string    `json:"status,omitempty"`
	Message     string    `json:"message,omitempty"`
	HttpStatus  int       `json:"httpStatus,omitempty"`
	ID          int64     `json:"id"`
	IsShadowed  bool      `json:"isShadowed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int64     `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
}

func ResponseMsg(status string, message string, httpStatus int) Response {
	resp := Response{
		Status:     status,
		Message:    message,
		HttpStatus: httpStatus,
	}
	return resp
}

func TaskResponseMsg(status string, message string, httpStatus int, task Entity.Task) Response {
	resp := Response{
		Status:      status,
		Message:     message,
		HttpStatus:  httpStatus,
		ID:          task.ID,
		IsShadowed:  task.IsShadowed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		Title:       task.Title,
		Description: task.Title,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
	}

	return resp
}
