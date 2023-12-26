package controller

import (
	"errors"
	"net/http"
	"strconv"

	"orm-crud-restapi/config"
	"orm-crud-restapi/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetBlogsController(c echo.Context) error {

	var blogs []model.Blog
	if err := config.DB.Find(&blogs).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success get all Blogs",

		"blogs": blogs,
	})

}

// get blog by id

func GetBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var blog model.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "blog not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"blog":    blog,
	})
}

// create new blog

func CreateBlogController(c echo.Context) error {

	blog := model.Blog{}

	c.Bind(&blog)
	user := model.User{}

	if err := config.DB.Where("id = ?", blog.UserID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error: "+err.Error())
	}

	if err := config.DB.Create(&blog).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success create new blog",

		"blog": blog,
	})

}

// delete blog by id

func DeleteBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var blog model.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var blogs []model.Blog
	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blog deleted",
		"blogs":   blogs,
	})
}

// update blog by id

func UpdateBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var blogToUpdate model.Blog
	if err := config.DB.First(&blogToUpdate, id).Error; err != nil {
		if config.DB.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&blogToUpdate)

	if err := config.DB.Save(&blogToUpdate).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var updatedBlogs []model.Blog
	if err := config.DB.Find(&updatedBlogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blog updated",
		"users":   updatedBlogs,
	})
}
