package main

import (
	"bwa_startup/auth"
	"bwa_startup/handler"
	"bwa_startup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect database mysql with GO
	dsn := "root:root@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Handle error 
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.rbY_zh7v_we9_S3bMNX3V-rp7YLS944LOYZd524aKSE")
	if err != nil {
		fmt.Println("Error")
	}

	if token.Valid {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailable)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}