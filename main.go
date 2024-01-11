package main

import (
	"database/sql"
	"fmt"
	"todo-app/handler"
	//"todo-app/models"
	"todo-app/repository"
	"todo-app/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" 
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Tingkatkanvalue25"
	dbname   = "todo_app_db"
)

func main() {
	fmt.Println("Welcome Todo App!")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database")

	// Inject DB ke Repository
	taskRepository := repository.NewTaskRepository(db)
	authorRepository := repository.NewAuthorRepository(db)

	// Inject Repository ke Usecase
	authorUseCase := usecase.NewAuthorUseCase(authorRepository)
	taskUseCase := usecase.NewTaskUseCase(taskRepository, authorUseCase)

	// Inject Usecase ke CHandler
	authorHandler := handler.NewAuthorHandler(authorUseCase)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	engine := gin.Default() 

	author := engine.Group("/api/v1/authors")
	{
		author.GET("/list", authorHandler.ListAuthor)
		author.GET("/get/:id", authorHandler.GetAuthor)
	}

	task := engine.Group("/api/v1/tasks")
	{
		task.GET("/list", taskHandler.FindTaskByAuthor)
		task.POST("/create", taskHandler.CreateHandler)
	}
	if err := engine.Run(":8080");err != nil {
		panic(fmt.Errorf("Failed to start server: %v", err))
	}

	/*// Membuat contoh task
	task := models.Task{
		Title:    "task 1",
		Content:  "Task 1 content",
		AuthorId: "id yang sesuai", // Gantilah dengan ID penulis yang valid
	}

	// Mendaftarkan task baru
	newTask, err := taskUseCase.RegisterNewTask(task)
	if err != nil {
		panic(err)
	}

	fmt.Println("New Task:", newTask)*/
}
