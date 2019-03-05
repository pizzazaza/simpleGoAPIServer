package main

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"net/http"
	//"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type userinfo struct {
	uid  int    `json:"uid"`
	id   int    `json:"id"`
	name string `json:"name"`
}

var (
	engine *xorm.Engine
)

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:Btpwns123@/simple?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	engine.SetConnMaxLifetime(60 * time.Second)
	engine.SetMapper(core.SameMapper{})

	// echo 프레임워크 객체 생성
	e := echo.New()

	e.POST("/users", createUser)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)

	log.Fatal(e.Start(":1323"))
}

//func getUser(c echo.Context) error {
//	//fmt.Println("hi")
//	//id, _ := strconv.Atoi(c.Param("id"))
//	//fmt.Println("hi")
//	//user := new(userinfo)
//	////userinfo.id = id
//	//fmt.Println(user)
//	//has, err := engine.Get(&user)
//	//if err != nil {
//	//	fmt.Errorf("error: %s \n", err)
//	//}
//	//fmt.Println(has)
//	//
//	return c.JSON(http.StatusOK, c)
//}

func createUser(c echo.Context) error {
	user := new(userinfo)
	user.id = 3
	user.name = "myname"
	affected, err := engine.Insert(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Insert ID", user.uid)

	return c.JSON(http.StatusCreated, affected)
}

//func updateUser(c echo.Context) error {
//	u := new(user)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//	id, _ := strconv.Atoi(c.Param("id"))
//	users[id].Name = u.Name
//	return c.JSON(http.StatusOK, users[id])
//}
//
//func deleteUser(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//	delete(users, id)
//	return c.NoContent(http.StatusNoContent)
//}
