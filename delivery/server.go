package delivery

import (
	"database/sql"
	"fmt"
	//"todo-app/config"
	"todo-app/delivery/controller"
	"todo-app/repository"
	"todo-app/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	authorUC usecase.AuthorUseCase
	taskUC usecase.TaskUseCase
	engine *gin.Engine
	host string
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api/v1")
	controller.NewAuthorController(s.authorUC, rg).Route()
	controller.NewTaskController(s.taskUC, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(":8080"); err != nil {
		panic(fmt.Errorf("server not running on host:8080, because error %v", err.Error()))
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(fmt.Errorf("config error: %v", err))
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Datababase)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(fmt.Errorf("connection error: %v", err))
	}
	// Inject DB ke repository 
	authorRepo := repository.NewAuthorRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	// Inject Repo ke useCase
	authorUC := usecase.NewAuthorUseCase(authorRepo)
	taskUC := usecase.NewTaskUseCase(taskRepo, authorUC)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cgf.ApiPort)

	return &Server {
		authorUC: authorUC,
		taskUC: taskUC, 
		engine: engine,
		host: host,
	}
}