package tasks

import (
	"TaskManager/models"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateTask(conn *pgx.Conn, task *models.Task) error {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO tasks (title, description, created_at, owner, status) VALUES ($1, $2, $3, $4, $5)",
		task.Title, task.Description, time.Now(), task.Owner, task.Status)
	if err != nil {
		log.Println("Failed to insert task:", err)
		return err
	}

	return nil
}
