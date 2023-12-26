package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int    `json:"id" form:"id"`
	Age   int    `json:"age" form:"age"`
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form:"name"`
}

func main() {
	e := echo.New()

	e.GET("/user", UserController)
	e.POST("/user", RegisterController)

	e.Start(":8000")
}

// response "Hello World"
func UserController(e echo.Context) error {
	id, _ := strconv.Atoi(e.QueryParam("id"))   // param = url
	age, _ := strconv.Atoi(e.QueryParam("age")) // queryParam =  queryparam
	search := e.QueryParam("search")
	short := e.QueryParam("short")
	user := User{
		Id:    id,
		Age:   age,
		Email: "Habib@email.co.id",
		Name:  "habib",
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"search": search,
		"short":  short,
	})
}

func RegisterController(input echo.Context) error {

	user := User{}
	input.Bind(&user)

	return input.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
