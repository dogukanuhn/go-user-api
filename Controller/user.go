package controller

import (
	"net/http"
	"user-basic/service"
	"github.com/labstack/echo"
	"user-basic/models"

)
 
 
type UserController struct {
		service service.UserService
}

  

// func (c *UserController) Login(ctx eco.Context) error  {
// 	user := service.GetUser(ctx)
// 	return ctx.JSON(http.StatusO, user)
// }

func(c *UserController) Register(ctx echo.Context) (err error) {

	u := &models.User{}

	if err = ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := service.Register(u)
  
	if (err != nil) {
		return ctx.JSON(http.StatusBadRequest, result)
	}

	return ctx.JSON(http.StatusOK, u)

}

var json map[string]interface{} = map[string]interface{}{}

func(c *UserController) Login(ctx echo.Context) (err error){



	if err := ctx.Bind(&json); err != nil {
		return err
	}
	result, err := service.Login(json["email"].(string))
	return ctx.JSON(http.StatusOK, result)
}

func(c *UserController) Authenticate(ctx echo.Context) (err error){


	auth := &models.Authenticate{}

	if err = ctx.Bind(auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := service.Authenticate(auth)

	return ctx.JSON(http.StatusOK, result)
}