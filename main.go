package main

import (
	"fmt"
	"go_project/internal/handlers"
	"go_project/internal/models"
	"go_project/pkg/db"
	"go_project/pkg/router"
)

func main() {
	// Import DB and TestConnection
	db := db.ConnectDB()
	db.AutoMigrate(&models.User{})
	r, err := router.InitRouter()

	if err != nil {
		fmt.Errorf(" Error init router %w ", err.Error())
		return
	}

	userRepo := handlers.NewUserRepository(db)

	r.Handle("GET", "/user", userRepo.GetAllUsers)
	r.Handle("POST", "/user", userRepo.CreateUser)
	r.Handle("GET", "/user/:id", userRepo.GetUserByID)
	r.Handle("DELETE", "/user/:id", userRepo.DeleteUser)
	r.Handle("PUT", "/user/:id", userRepo.UpdateUser)

	err = r.Run()
	if err != nil {
		return
	}
	fmt.Println(db.Config)

}
