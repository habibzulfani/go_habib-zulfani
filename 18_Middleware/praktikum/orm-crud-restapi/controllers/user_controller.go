package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"orm-crud-restapi/config"
	"orm-crud-restapi/lib/database"
	"orm-crud-restapi/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {

	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	if err := config.DB.Preload("Blogs").Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success get all users",

		"users": users,
	})

}

// get user by id

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Preload("Blogs").Find(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// create new user

func CreateUserController(c echo.Context) error {

	user := models.User{}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success create new user",

		"user": user,
	})

}

// delete user by id

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted",
		"users":   users,
	})
}

// update user by id

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var userToUpdate models.User
	if err := config.DB.First(&userToUpdate, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&userToUpdate)

	if err := config.DB.Save(&userToUpdate).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var updatedUsers []models.User
	if err := config.DB.Find(&updatedUsers).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated",
		"users":   updatedUsers,
	})
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	fmt.Printf("Email from request: %s\n", user.Email)
	fmt.Printf("Password from request: %s\n", user.Password)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logged in successfully!",
		"users":   users,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := database.GetDetailUsers((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"users":   users,
	})
}
