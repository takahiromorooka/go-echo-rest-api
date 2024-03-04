package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"go-echo-rest-api/model"
)

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	users := []model.User{}
	model.DB.Find(&users)
	fmt.Println(users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := model.User{}
	id := c.Param("id")
	if err := c.Bind(&user); err != nil {
		return err
	}
	fmt.Println(id)
	fmt.Println(user)
	model.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	fmt.Println(user)
	model.DB.Model(&user).Updates(&user)
	fmt.Println(user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	user := model.User{}
	id := c.Param("id")
	model.DB.Delete(&user, id)
	return c.NoContent(http.StatusNoContent)
}