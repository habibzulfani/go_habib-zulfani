package controller

import (
	"net/http"
	"strconv"

	"orm-crud-restapi/config"
	"orm-crud-restapi/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {

	var users []model.User

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

	var user model.User
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

	user := model.User{}

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

	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var users []model.User
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

	var userToUpdate model.User
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

	var updatedUsers []model.User
	if err := config.DB.Find(&updatedUsers).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated",
		"users":   updatedUsers,
	})
}
