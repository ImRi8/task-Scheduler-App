package models

import "time"

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
