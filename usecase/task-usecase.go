package usecase

import (
	"fmt"
	"todo-app/models"
	"todo-app/repository"
)

type TaskUseCase interface {
	RegisterNewTask(payload models.Task) (models.Task, error)
	FindTaskByAuthor(id string) ([]models.Task, error)
	FindAllTasks() ([]models.Task, error)
}

type taskUseCase struct {
	taskRepository repository.TaskRepository
	authorUC       AuthorUseCase
}

func (t *taskUseCase) FindAllTasks() ([]models.Task, error) {
	return t.taskRepository.List()
}

func (t *taskUseCase) FindTaskByAuthor(authorID string) ([]models.Task, error) {
	// Periksa apakah penulis ada
	_, err := t.authorUC.FindAuthorById(authorID)
	if err != nil {
		return nil, fmt.Errorf("Gagal menemukan tugas: penulis tidak ditemukan")
	}

	// Penulis ada, lanjutkan untuk menemukan tugas oleh penulis
	tasks, err := t.taskRepository.GetTasksByAuthor(authorID)
	if err != nil {
		return nil, fmt.Errorf("Gagal menemukan tugas: %v", err)
	}

	return tasks, nil
}

func (t *taskUseCase) RegisterNewTask(payload models.Task) (models.Task, error) {
	// Validasi authorId
	_, err := t.authorUC.FindAuthorById(payload.AuthorId)
	if err != nil {
		return models.Task{}, fmt.Errorf("Gagal membuat tugas: penulis tidak ditemukan")
	}

	// Validasi judul dan konten
	if payload.Title == "" {
		return models.Task{}, fmt.Errorf("Judul tidak boleh kosong")
	}
	// Anda dapat menambahkan validasi tambahan untuk konten jika diperlukan

	// Buat tugas
	task, err := t.taskRepository.Create(payload)
	if err != nil {
		return models.Task{}, fmt.Errorf("Gagal membuat tugas: %v", err)
	}

	return task, nil
}

func NewTaskUseCase(taskRepository repository.TaskRepository, authorUC AuthorUseCase) TaskUseCase {
	return &taskUseCase{
		taskRepository: taskRepository,
		authorUC:       authorUC,
	}
}
