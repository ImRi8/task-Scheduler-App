package models

import "time"

type Request struct {
	ID          int64     `json:"id"`
	IsShadowed  bool      `json:"isShadowed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int64     `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
}
