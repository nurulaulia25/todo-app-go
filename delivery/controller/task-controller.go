package controller 

import (
	"net/http"
	"strconv"
	"todo-app/models"
	"todo-app/usecase"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUC usecase.TaskUseCase
}

func NewTaskController(taskUC usecase.TaskUseCase, rg *gin.RouterGroup) *TaskController {
	return &TaskController{
		taskUC: taskUC,
		rg: rg,
	}
}

func (t *TaskController) Route() {
	t.rg.POST("/tasks/create", t.CreateHandler)
	t.rg.GET("/tasks/list", t.ListHandler)
	t.rg.GET("tasks/get/:id", t.GetByAuthorHandler)
}

func (t *TaskController) CreateHandler(ctx *gin.Context) {
	var payload models.Task 
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	task, err := t.taskUC.RegisterNewTask(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": gin.H {
			"code": http.StatusCreated,
			"message": "Created",
		}, 
		"data": task,
	})
}

func (t *TaskController) ListHandler (ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Queery("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	tasks, paging, err := t.taskUC.FindAllTasks(page, sizee)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	var response []interface {
		for _, v := range tasks {
			response := append(response, v)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"code": http.StatusOK,
				"message": "Ok",
			},
			"data": response,
			"pagging": pagging,
		})
	}

func (t *TaskController) GetByAuthorHandler(ctx *gin.Context) {
	author := ctx.Param("id")
	tasks, err := t.taskUC.FindTaskByAuthor(author)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusNotFound,
			"message": "task with author ID" + author + "not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code": http.StatusOK,
			"message": "Oke"
		}, 
		"data": tasks,
	})
}