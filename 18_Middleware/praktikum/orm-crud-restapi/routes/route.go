package routes

import (
	"orm-crud-restapi/constants"
	"orm-crud-restapi/controllers"

	m "orm-crud-restapi/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// create a new echo instance

	e := echo.New()

	// logger
	m.LogMiddlewares(e)

	// auth middlewares
	eAuth := e.Group("/auth")
	eAuth.Use(mid.BasicAuth(m.BasicAuthDB))
	e.POST("/login", controllers.LoginUsersController)

	// auth - users
	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.GET("/users/:id", controllers.GetUserController)
	eAuth.POST("/users", controllers.CreateUserController)
	eAuth.DELETE("/users/:id", controllers.DeleteUserController)
	eAuth.PUT("/users/:id", controllers.UpdateUserController)

	// auth - books
	eAuth.GET("/books", controllers.GetBooksController)
	eAuth.GET("/books/:id", controllers.GetBookController)
	eAuth.POST("/books", controllers.CreateBookController)
	eAuth.DELETE("/books/:id", controllers.DeleteBookController)
	eAuth.PUT("/books/:id", controllers.UpdateBookController)

	// auth - blogs
	eAuth.GET("/blogs", controllers.GetBlogsController)
	eAuth.GET("/blogs/:id", controllers.GetBlogController)
	eAuth.POST("/blogs", controllers.CreateBlogController)
	eAuth.DELETE("/blogs/:id", controllers.DeleteBlogController)
	eAuth.PUT("/blogs/:id", controllers.UpdateBlogController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUserDetailControllers)

	return e
}
