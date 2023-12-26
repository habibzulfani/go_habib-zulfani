package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mini_project/internal/delivery/http"
	"mini_project/internal/repository"
	"mini_project/internal/usecase"
	"mini_project/internal/domain"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := gorm.Open(mysql.Open("library.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database")
	}
	db.AutoMigrate(&domain.Book{})

	bookRepository := repository.NewBookRepository(db)
	bookUseCase := usecase.NewBookUseCase(bookRepository)
	bookHandler := http.NewBookHandler(bookUseCase)

	e := echo.New()

	http.SetMiddleware(e)
	http.SetJWTMiddleware(e) // Add JWT middleware

	// Admin group with a custom middleware to check for admin role
	adminGroup := e.Group("/admin")
	adminGroup.Use(http.AdminMiddleware)
	adminGroup.GET("/books", bookHandler.GetBooks)
	adminGroup.GET("/books/:id", bookHandler.GetBookByID)
	adminGroup.POST("/books", bookHandler.CreateBook)
	adminGroup.DELETE("/books/:id", bookHandler.DeleteBook)
	adminGroup.POST("/admin", bookHandler.CreateAdmin)


	// User group
	userGroup := e.Group("/user")
	userGroup.POST("/register", bookHandler.RegisterUser)
	userGroup.POST("/login", bookHandler.LoginUser)
	userGroup.GET("/books", bookHandler.GetBooks)

	e.Logger.Fatal(e.Start(":8080"))
}
