package main

import (
	"user-basic/controller"
	"user-basic/models"


	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"

)

func restricted(c echo.Context) error {
	if temp := c.Get("user"); temp != nil {
		user := temp .(*jwt.Token)
		claims := user.Claims.(*models.JwtCustomClaims)
		name := claims.Email
	return c.String(http.StatusOK, "asdf "+name)
	}
	return c.String(http.StatusOK, "Welcome ")
}

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	Claims:     &models.JwtCustomClaims{},
    SigningKey: []byte("secret"),
})


func main() {
	e := echo.New()


	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	


	userC := controller.UserController{}

	// e.GET("/", userC.GetUser)
	e.POST("/register", userC.Register)
	e.POST("/login", userC.Login)
	e.POST("/auth", userC.Authenticate)
	e.GET("/test",restricted, IsLoggedIn)




	e.Logger.Fatal(e.Start(":4000"))
}
