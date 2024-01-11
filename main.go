package main

import "todo-app/delivery"

//"database/sql"
//"fmt"
//"todo-app/delivery"
//"todo-app/handler"
//"todo-app/models"
//"todo-app/repository"
//"todo-app/usecase"

//"github.com/gin-gonic/gin"
//_ "github.com/lib/pq"


func main() {
	delivery.NewServer().Run()
}
