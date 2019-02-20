package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

func main() {
	e := echo.New()
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	user1 := user{}
	user1.ID = 1
	user1.Name = "hi"

	users[1] = &user1

	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func saveUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
