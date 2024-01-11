package controller

import (
	"net/http"
	"todo-app/shared/common"
	"todo-app/usecase"
	"todo-app/shared/shared-models"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorUC usecase.AuthorUseCase
	rg *gin.RouterGroup
}

func (a *AuthorController) listHandler (ctx *gin.Context) {
	author := ctx.MustGet("author").(string)
	authors , err := a.authorUC.FindAllAuthor(author)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range authors {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, shared-model.Paging{}, "Ok")
}

func (a *AuthorController) getHandler (ctx *gin.Context) {
	id := ctx.Param("id")
	author, err := a.authorUC.FindAuthorById(id)
	if err != nil {
		common.SendErrrorResponse(ctx, http.StatusFound, "author with ID" + id + "not found")
		return
	}
	common.SendSingleResponse(ctx, author, "Ok")
}

func (a *AuthorController) Route() {
	a.rg.GET("/authors/list", a.listHandler)
	a.rg.GET("/authors/get/:id", a.getHandler)
}

func NewAuthorController(authorUC usecase.AuthorUseCase, rg *gin.Router) *AuthorController {
	return &AuthorController{
		authorUC: authorUC,
		rg: rg,
	}
}