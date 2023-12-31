package SqlDB

import (
	"Task-scheduler-App/Datalayer/Entity"
	"database/sql"
	"gofr.dev/pkg/gofr"
)

type TaskDbService struct {
}

const (
	readTaskById   = "SELECT id,is_shadowed,created_at,updated_at,title,description,priority,due_date FROM task WHERE id = ? and is_shadowed = ?;"
	updateTaskById = "UPDATE task SET title= ?,description =?,priority=?,due_date=? WHERE id = ?;"
	shadowTaskById = "UPDATE task SET is_shadowed = ? WHERE id = ?;"
	createTaskById = "INSERT INTO task (is_shadowed, title, description, priority, due_date) VALUES (?, ?, ?, ?, ?);"
)

func (taskDbService *TaskDbService) ShadowRowInEntity(Id int64, ctx *gofr.Context) error {
	res, err := ctx.DB().ExecContext(ctx, shadowTaskById, true, Id)

	if err != nil {
		ctx.Logger.Error("Error while Updating the Data", err.Error())
		return err
	}

	rowsAffected, rowsAffErr := res.RowsAffected()

	if rowsAffErr != nil {
		ctx.Logger.Error("Error while reading the rows ", err)
		return rowsAffErr
	}

	if rowsAffected == 0 {
		ctx.Logger.Error("Unable to update the row")
		//return error.Error("Db failed to updated")
	}
	return nil
}

func (taskDbService *TaskDbService) UnShadowRowInEntity(Id int64, ctx *gofr.Context) error {
	res, err := ctx.DB().ExecContext(ctx, shadowTaskById, false, Id)

	if err != nil {
		ctx.Logger.Error("Error while Updating the Data", err)
		return err
	}

	rowsAffected, rowsAffErr := res.RowsAffected()

	if rowsAffErr != nil {
		ctx.Logger.Error("Error while reading the rows ", err)
		return rowsAffErr
	}

	if rowsAffected == 0 {
		ctx.Logger.Error("Unable to update the row")
		//return error.Error("Db failed to updated")
	}
	return nil
}

func (taskDbService *TaskDbService) CreateRowInEntity(ctx *gofr.Context, obj interface{}) (interface{}, error) {
	task := obj.(*Entity.Task)

	res, err := ctx.DB().ExecContext(ctx, createTaskById, task.IsShadowed, task.Title, task.Description, task.Priority, task.DueDate)

	if err != nil {
		ctx.Logger.Error("Error while creating the record", err)
		return nil, err
	}

	rowsAffected, rowsAffErr := res.RowsAffected()

	resId, _ := res.LastInsertId()
	task.ID = resId

	if rowsAffErr != nil {
		ctx.Logger.Error("Error while creating a row", err)
		return nil, rowsAffErr
	}

	if rowsAffected == 0 {
		ctx.Logger.Error("Error while creating a new row")
	}

	return task, nil
}

func (taskDbService *TaskDbService) UpdateRowInEntity(ctx *gofr.Context, obj interface{}, id int64) (interface{}, error) {
	task := obj.(*Entity.Task)

	res, err := ctx.DB().ExecContext(ctx, updateTaskById, task.Title, task.Description, task.Priority, task.DueDate, id)

	if err != nil {
		ctx.Logger.Error("Error while Updating the Data", err)
		return nil, err
	}

	rowsAffected, rowsAffErr := res.RowsAffected()

	if rowsAffErr != nil {
		ctx.Logger.Error("Error while reading the rows ", err)
		return nil, rowsAffErr
	}

	if rowsAffected == 0 {
		ctx.Logger.Error("Unable to update the row")
		return nil, nil
	}
	return task, nil
}

func (taskDbService *TaskDbService) GetEntityById(Id int64, ctx *gofr.Context) (interface{}, error) {
	var task Entity.Task
	rows, err := ctx.DB().QueryContext(ctx, readTaskById, Id, false)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Logger.Debug("No rows found for this id", Id)
			return nil, sql.ErrNoRows
		}
		ctx.Logger.Error("Error while reading from db", err)
		return nil, err
	}

	for rows.Next() {
		scnErr := rows.Scan(&task.ID,
			&task.IsShadowed,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.DueDate)
		if scnErr != nil {
			ctx.Logger.Error("Error while scanning the rows", scnErr.Error())
			return nil, scnErr
		}
	}

	if task.ID == 0 {
		return nil, nil
	}

	return task, nil
}
