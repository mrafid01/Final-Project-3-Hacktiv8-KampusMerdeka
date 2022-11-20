package main

import (
	"os"
	"project3/config"
	"project3/controller"
	"project3/middleware"
	"project3/repository"
	"project3/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	db := config.InitDB(dbUsername, dbPassword, dbHost, dbPort, dbName)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskController := controller.NewTaskController(taskService)

	router := gin.Default()

	// User
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.PUT("/users/update-account", middleware.AuthMiddleware, userController.UpdateUser)
	router.DELETE("/users/delete-account", middleware.AuthMiddleware, userController.DeleteUser)

	// Category
	router.POST("/categories", middleware.AuthMiddleware, categoryController.CreateCategory)
	router.GET("/categories", middleware.AuthMiddleware, categoryController.GetCategory)
	router.PATCH("/categories/:id", middleware.AuthMiddleware, categoryController.PatchCategory)
	router.DELETE("/categories/:id", middleware.AuthMiddleware, categoryController.DeleteCategory)

	// Task
	router.POST("/tasks", middleware.AuthMiddleware, taskController.CreateTask)
	router.GET("/tasks", middleware.AuthMiddleware, taskController.GetTask)
	router.PUT("/tasks/:id", middleware.AuthMiddleware, taskController.UpdateTask)
	router.PATCH("/tasks/update-status/:id", middleware.AuthMiddleware, taskController.PatchStatusTask)
	router.PATCH("/tasks/update-category/:id", taskController.PatchCategoryTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware, taskController.DeleteTask)

	router.Run(":" + port)
}
