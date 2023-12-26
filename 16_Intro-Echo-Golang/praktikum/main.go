package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user",
				"user":     user,
			})
		}
	}

	return c.String(http.StatusNotFound, "User not found")
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	indexToDelete := -1
	for i, user := range users {
		if user.Id == id {
			indexToDelete = i
			break
		}
	}

	// If the user was found, remove them from the slice
	if indexToDelete != -1 {
		users = append(users[:indexToDelete], users[indexToDelete+1:]...)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "User deleted",
			"user":     users,
		})
	}

	return c.String(http.StatusNotFound, "User not found")

}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	for i, user := range users {
		if user.Id == id {
			// Bind the request data to the existing user
			if err := c.Bind(&users[i]); err != nil {
				return c.String(http.StatusInternalServerError, "Failed to update user")
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "User updated",
				"user":     users[i],
			})
		}
	}

	return c.String(http.StatusNotFound, "User not found")

}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
