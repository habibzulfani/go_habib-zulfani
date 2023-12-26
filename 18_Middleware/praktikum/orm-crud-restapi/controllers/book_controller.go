package controllers

import (
	"net/http"
	"strconv"

	"orm-crud-restapi/config"
	"orm-crud-restapi/models"

	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
)

func GetBooksController(c echo.Context) error {

	var books []models.Books

	if err := config.DB.Find(&books).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success get all Books",

		"books": books,
	})

}

// get book by id

func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var book models.Books
	if err := config.DB.First(&book, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "book not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// create new book

func CreateBookController(c echo.Context) error {

	book := models.Books{}

	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success create new book",

		"book": book,
	})

}

// delete book by id

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var book models.Books
	if err := config.DB.First(&book, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Book not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var books []models.Books
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book deleted",
		"books":   books,
	})
}

// update book by id

func UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var bookToUpdate models.Books
	if err := config.DB.First(&bookToUpdate, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Book not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&bookToUpdate)

	if err := config.DB.Save(&bookToUpdate).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var updatedBooks []models.Books
	if err := config.DB.Find(&updatedBooks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book updated",
		"books":   updatedBooks,
	})
}
