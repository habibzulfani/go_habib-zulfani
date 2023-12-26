package middlewares

import (
	"orm-crud-restapi/config"
	"orm-crud-restapi/models"

	"github.com/labstack/echo"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User

	err := db.Where("email = ? AND password = ?", username, password).First(&user). Error

	if err != nil {
		return false, err
	}

	return true, nil
}