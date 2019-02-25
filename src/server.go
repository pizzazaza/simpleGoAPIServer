package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	//_ "github.com/go-sql-driver/mysql"
	"time"

	//"fmt"

	//"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

type (
	user struct {
		UID  int    `json:"uid"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err := xorm.NewEngine("mysql", "root:Btpwns123@/simpleApi")

	e := echo.New()

	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := &user{
		ID: id,
	}

	has, _ := engine.Get(&u)

	fmt.Println(has)

	return c.JSON(http.StatusOK, users[id])
}

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}

	engine.Insert(u)

	fmt.Println("Insert ID", u.uid)

	return c.JSON(http.StatusCreated, result)
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
