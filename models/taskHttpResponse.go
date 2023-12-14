package models

import (
	"Task-scheduler-App/Datalayer/Entity"
	"time"
)

type Response struct {
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	HttpStatus int    `json:"httpStatus,omitempty"`
	ID         int64  `json:"id,omitempty"`
	IsShadowed bool   `json:"isShadowed,omitempty"`
	//CreatedAt   time.Time `json:"createdAt,omitempty"`
	//UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int64     `json:"priority,omitempty"`
	DueDate     time.Time `json:"dueDate,omitempty"`
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
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
	}

	return resp
}
