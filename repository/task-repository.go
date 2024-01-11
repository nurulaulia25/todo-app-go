package repository

import (
	"database/sql"
	"log"
	"todo-app/models"
)

type TaskRepository interface {
	Create(payload models.Task) (models.Task, error)
	List() ([]models.Task, error)
	GetTasksByAuthor(authorID string) ([]models.Task, error)
}

func NewTaskRepository(db *sql.DB) TaskRepository {
    return &taskRepository{db: db}
}

type taskRepository struct {
	db *sql.DB
}

func (t *taskRepository) Create(payload models.Task) (models.Task, error) {
	var task models.Task
	err := t.db.QueryRow("INSERT INTO tasks (title, content, author_id) VALUES ($1, $2, $3) RETURNING id, created_at",
		payload.Title, payload.Content, payload.AuthorId).Scan(&task.Id, &task.CreatedAt)
	if err != nil {
		log.Printf("taskRepository.Create: %v\n", err)
		return models.Task{}, err
	}
	task.Title = payload.Title
	task.Content = payload.Content
	task.AuthorId = payload.AuthorId
	return task, nil
}

func (t *taskRepository) List() ([]models.Task, error) {
	sql := "SELECT id, title, content, author_id, created_at FROM tasks"

	var tasks []models.Task
	rows, err := t.db.Query(sql)
	if err != nil {
		log.Printf("taskRepository.List: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Content, &task.AuthorId, &task.CreatedAt)
		if err != nil {
			log.Printf("taskRepository.List: %v\n", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepository) GetTasksByAuthor(authorID string) ([]models.Task, error) {
	sql := "SELECT id, title, content, author_id, created_at FROM tasks WHERE author_id = $1"

	var tasks []models.Task
	rows, err := t.db.Query(sql, authorID)
	if err != nil {
		log.Printf("taskRepository.GetTasksByAuthor: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Content, &task.AuthorId, &task.CreatedAt)
		if err != nil {
			log.Printf("taskRepository.GetTasksByAuthor: %v\n", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
