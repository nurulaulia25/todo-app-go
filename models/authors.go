package models

import (
	"time"
)

type Author struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
	Role string `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Tasks []Task `json:"task"`
}