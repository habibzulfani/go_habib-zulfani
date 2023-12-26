package route

import (
	"orm-crud-restapi/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance

	e := echo.New()

	// Route / to handler function

	// Users Route

	e.GET("/users", controller.GetUsersController)

	e.GET("/users/:id", controller.GetUserController)

	e.POST("/users", controller.CreateUserController)

	e.DELETE("/users/:id", controller.DeleteUserController)

	e.PUT("/users/:id", controller.UpdateUserController)

	// Books Route

	e.GET("/books", controller.GetBooksController)

	e.GET("/books/:id", controller.GetBookController)

	e.POST("/books", controller.CreateBookController)

	e.DELETE("/books/:id", controller.DeleteBookController)

	e.PUT("/books/:id", controller.UpdateBookController)

	// Blog Route

	e.GET("/blogs", controller.GetBlogsController)

	e.GET("/blogs/:id", controller.GetBlogController)

	e.POST("/blogs", controller.CreateBlogController)

	e.DELETE("/blogs/:id", controller.DeleteBlogController)

	e.PUT("/blogs/:id", controller.UpdateBlogController)

	return e
}
