package handler

import (
	"net/http"
	"todo-app/models"
	"todo-app/usecase"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskUC usecase.TaskUseCase
}

func NewTaskHandler(taskUC usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUC: taskUC,
	}
}

func (a *TaskHandler) CreateHandler(c *gin.Context) {
	var paload models.Task
	err := c.Bind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest), gin.Context{
			"Code": http.StatusBadRequest,
			"Message": "Payload not found"
		}
	}
}

func (a *TaskHandler) FindTaskByAuthor(c *gin.Context) {
	author := c.Param("id")
	tasks, err := a.taskUC.FindTaskByAuthor(author)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "task with ID " + author + "not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "success",
		"data": tasks,
	})
}