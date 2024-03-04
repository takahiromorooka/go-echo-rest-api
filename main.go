package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-echo-rest-api/controller"
	"go-echo-rest-api/model"
)

func connect(c echo.Context) error {
	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "database can't connect!!!!!!")
	} else {
		return c.String(http.StatusOK, "database connected")
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}