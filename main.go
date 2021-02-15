package main

import (
	"user-basic/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

)

func main() {
	e := echo.New()


	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userC := controller.UserController{}

	// e.GET("/", userC.GetUser)
	e.POST("/register", userC.Register)
	e.POST("/login", userC.Login)


	e.Logger.Fatal(e.Start(":4000"))
}
