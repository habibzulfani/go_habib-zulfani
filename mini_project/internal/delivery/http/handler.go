package http

import (
	"mini_project/internal/domain"
	"mini_project/internal/usecase"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// BookHandler handles HTTP requests for books
type BookHandler struct {
	bookUseCase usecase.BookUseCase
	userUseCase usecase.UserUseCase
}

// NewBookHandler creates a new instance of BookHandler
func NewBookHandler(bookUseCase usecase.BookUseCase) *BookHandler {
	return &BookHandler{bookUseCase: bookUseCase}
}

// GetBooks handles the request to get all books
func (h *BookHandler) GetBooks(c echo.Context) error {
	// Access user claims to check the role
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role, ok := claims["role"].(string)

	var books []domain.Book
	var err error

	// Differentiate logic based on user role
	if ok && role == "admin" {
		books, err = h.bookUseCase.GetAllBooks()
	} else {
		books, err = h.bookUseCase.GetUserBooks()
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}

// GetBookByID handles the request to get a book by ID
func (h *BookHandler) GetBookByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	book, err := h.bookUseCase.GetBookByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, book)
}

// CreateBook handles the request to create a new book
func (h *BookHandler) CreateBook(c echo.Context) error {
	book := new(domain.Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err := h.bookUseCase.CreateBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	err = h.bookUseCase.DeleteBook(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *BookHandler) RegisterUser(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err := h.userUseCase.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *BookHandler) LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.userUseCase.Login(username, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token, err := generateToken(user.IsAdmin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func generateToken(isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"role": "user",
	}
	if isAdmin {
		claims["role"] = "admin"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
