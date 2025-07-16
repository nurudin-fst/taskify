package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nurudin-fst/taskify/config"
	"github.com/nurudin-fst/taskify/db"
	"github.com/nurudin-fst/taskify/internal/app"
	"github.com/nurudin-fst/taskify/internal/handler"
	"github.com/nurudin-fst/taskify/internal/repository"
	"github.com/nurudin-fst/taskify/internal/usecase"
)

func main() {
	config.LoadEnv()
	db := db.Init()
	fmt.Println(db)

	userRepo := repository.NewUserRepo(db)
	projectRepo := repository.NewProjectRepo(db)
	taskRepo := repository.NewTaskRepo(db)

	userUc := usecase.NewUserUC(userRepo)
	projectUC := usecase.NewProjectUC(projectRepo)
	taskUC := usecase.NewTaskUC(taskRepo)

	userHandler := handler.NewUserHandler(userUc)
	projectHandler := handler.NewProjectHandler(projectUC)
	taskHandler := handler.NewTaskHandler(taskUC)

	app := app.InitFiberApp()
	userHandler.Router(app)
	projectHandler.Router(app)
	taskHandler.Router(app)

	port := os.Getenv("APP_PORT")
	err := app.Listen(":" + port)
	log.Fatal(err)

}
