package main

import (
	"LifeScribe_Backend/internal/api/controller"
	"LifeScribe_Backend/internal/api/repository"
	"LifeScribe_Backend/internal/api/usecase"
	"LifeScribe_Backend/internal/db"
	"LifeScribe_Backend/pkg/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	diaryRepository := repository.NewDiaryRepository(db)
	diaryUsecase := usecase.NewDiaryUsecase(diaryRepository)
	diaryController := controller.NewDiaryController(diaryUsecase)

	routes := routes.NewRouter(userController, diaryController)
	routes.Logger.Fatal(routes.Start(":8080"))
}
