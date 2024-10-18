package main

import (
	"fmt"
	"go_project/internal/handlers"
	"go_project/middlewares"
	"go_project/pkg/db"
	"go_project/pkg/router"
	
)

func main() {
	// Import DB and TestConnection
	db.InitDB()
	r, err := router.InitRouter()

	if err != nil {
		fmt.Errorf(" Error init router %w ", err.Error())
		return
	}

	r.LoadHTMLGlob("static/web/*")
	r.Static("/static", "./static/web")



	r.Handle("GET", "/", handlers.Home)

	userRepo := handlers.NewUserRepository(db.DB)

	r.Handle("GET", "/user", middlewares.IsAuthorized(), middlewares.IsAdmin(), userRepo.GetAllUsers)
	r.Handle("POST", "/signup", handlers.Signup)
	r.Handle("GET", "/user/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), userRepo.GetUserByID)
	r.Handle("DELETE", "/user/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), userRepo.DeleteUser)
	r.Handle("PUT", "/user/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(),  userRepo.UpdateUser)
	r.Handle("POST", "/login", handlers.Login)
	r.Handle("GET", "/logout", handlers.Logout)
	
	itemRepo := handlers.NewItemRepository(db.DB)
	r.Handle("GET", "/products", middlewares.IsAuthorized(), itemRepo.GetAllItems)
	r.Handle("GET", "/products/:id", middlewares.IsAuthorized(), itemRepo.GetItemByID)
	r.Handle("POST", "/products", middlewares.IsAuthorized(), middlewares.IsAdmin(), itemRepo.CreateItem)
	r.Handle("PUT", "/products/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), itemRepo.UpdateItem)
	r.Handle("DELETE", "/products/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), itemRepo.DeleteItem)

	categoryRepo := handlers.NewCategoryRepository(db.DB)
	r.Handle("GET", "/categories", middlewares.IsAuthorized(), categoryRepo.GetAllCategories)
	r.Handle("POST", "/categories", middlewares.IsAuthorized(), middlewares.IsAdmin(), categoryRepo.CreateCategory)
	r.Handle("GET", "/categories/:id", middlewares.IsAuthorized(), categoryRepo.GetCategoryByID)
	r.Handle("PUT", "/categories/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), categoryRepo.UpdateCategory)
	r.Handle("DELETE", "/categories/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), categoryRepo.DeleteCategory)

	orderRepo := handlers.NewOrderRepository(db.DB)
	r.Handle("POST", "/orders", middlewares.IsAuthorized(), orderRepo.CreateOrder)
	r.Handle("PUT", "/orders/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), orderRepo.UpdateOrder)
	r.Handle("GET", "/orders/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), orderRepo.GetOrderByID)
	r.Handle("GET", "/orders", middlewares.IsAuthorized(), middlewares.IsAdmin(), orderRepo.GetAllOrders)
	r.Handle("DELETE", "/orders/:id", middlewares.IsAuthorized(), middlewares.IsAdmin(), orderRepo.DeleteOrder)







	



	err = r.Run()
	if err != nil {
		return
	}
	fmt.Println(db.DB.Config)

}
