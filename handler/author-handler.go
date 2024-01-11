package handler

import (
	"todo-app/usecase"
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	usecase.AuthorUseCase
}

func NewAuthorHandler (authorUC usecase.AuthorUseCase) *AuthorHandler {
	return 
}

func (a *AuthorHandler) ListAuthor(*gin.Context) {
	author, err := a.AuthorUseCase.FindAllAuthor()
}

func (a *AuthorHandler) GetAuthor(c *gin.Context) {

}