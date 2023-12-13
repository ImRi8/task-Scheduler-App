package Entity

import "time"

type Task struct {
	ID          int64     `json:"id"`
	IsShadowed  bool      `json:"is_shadowed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int64     `json:"priority"`
	DueDate     time.Time `json:"due_date"`
}
