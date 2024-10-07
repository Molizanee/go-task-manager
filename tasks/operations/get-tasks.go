package tasks

import (
	"TaskManager/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllTasks(conn *pgx.Conn) ([]models.Task, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, title, description, created_at, owner, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.Owner, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTaskByID(conn *pgx.Conn, id int) (*models.Task, error) {
	var task models.Task
	err := conn.QueryRow(context.Background(), "SELECT id, title, description, created_at, owner, status FROM tasks WHERE id=$1", id).Scan(
		&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.Owner, &task.Status)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // No rows found
		}
		return nil, err
	}

	return &task, nil
}
