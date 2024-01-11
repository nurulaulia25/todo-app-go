package usecase

import (
	"todo-app/models"
	"todo-app/repository"
)

type AuthorUseCase interface {
	FindAllAuthor() ([]models.Author, error)
	FindAuthorById(id string) (models.Author, error)
	FindAuthorByEmail(email string) (models.Author, error)
}

type authorUseCase struct {
	authorRepository repository.AuthorRepository
}
func (a *authorUseCase) FindAllAuthor() ([]models.Author, error) {
	return a.authorRepository.List()
}

func (a *authorUseCase) FindAuthorByEmail (email string) (models.Author, error) {
	return a.authorRepository.GetByEmail(email)
}

func(a *authorUseCase) FindAuthorById(id string) (models.Author, error) {
	return a.authorRepository.Get(id)
}

func NewAuthorUseCase(authorRepository repository.AuthorRepository) AuthorUseCase {
	return &authorUseCase{
		authorRepository: authorRepository,
	}
}

