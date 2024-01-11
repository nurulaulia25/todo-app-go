package repository

import (
	"database/sql"
	"log"
	"todo-app/models"
)

type AuthorRepository interface {
	GetByEmail(email string) (models.Author, error)
	Get(id string) (models.Author, error)
	List() ([]models.Author, error)
	Create(payload models.Author) (models.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

// Get implements AuthorRepository.
func (*authorRepository) Get(id string) (models.Author, error) {
	panic("unimplemented")
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

func (r *authorRepository) GetByEmail(email string) (models.Author, error) {
	sql := "SELECT id, name, email, role, created_at FROM authors WHERE email = $1"

	var author models.Author
	err := r.db.QueryRow(sql, email).Scan(&author.Id, &author.Name, &author.Email, &author.Role, &author.CreatedAt)
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}

func (r *authorRepository) Create(payload models.Author) (models.Author, error) {
	var author models.Author
	err := r.db.QueryRow("INSERT INTO authors (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at", payload.Name, payload.Email, payload.Password, payload.Role).Scan(&author.Id, &author.CreatedAt)
	if err != nil {
		log.Println("authorRepository.QueryRow", err.Error())
		return models.Author{}, err
	}
	author.Name = payload.Name
	author.Email = payload.Email
	author.Role = payload.Role
	return author, nil
}

func (r *authorRepository) List() ([]models.Author, error) {
	sql := "SELECT id, name, email, role, created_at FROM authors"

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.Id, &author.Name, &author.Email, &author.Role, &author.CreatedAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}
