package main

import (
	"os"

	"github.com/TiagoAmaralFerreira/go-echo/config"
	"github.com/TiagoAmaralFerreira/go-echo/internal/handle"
	jwt "github.com/TiagoAmaralFerreira/go-echo/internal/middlewares"
	"github.com/TiagoAmaralFerreira/go-echo/internal/repository"
	"github.com/TiagoAmaralFerreira/go-echo/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := config.ConnectDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handle.NewUserHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	secret := os.Getenv("JWT_SECRET")
	e.POST("/users", userHandler.CreateUser, jwt.JWTMiddleware(secret))
	// e.GET("/users", userHandler.ListUsers, jwt.JWTMiddleware(secret))
	// e.GET("/users/:id", userHandler.GetUser, jwt.JWTMiddleware(secret))
	// e.PUT("/users/:id", userHandler.UpdateUser, jwt.JWTMiddleware(secret))
	// e.DELETE("/users/:id", userHandler.DeleteUser, jwt.JWTMiddleware(secret))

	e.Logger.Fatal(e.Start(":8080"))
}
