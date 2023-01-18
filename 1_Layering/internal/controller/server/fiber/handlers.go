package fiber

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
	"go.uber.org/zap/zapcore"
)

//	@Summary		Get Book
//	@Description	Get Book by ID.
//	@Tags			Book
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	domain.Book
//	@Failure		400	{object}	fiber.errorMessage
//	@Router			/api/books/{id} [get]
func (s *Server) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		logger.Logger().Warn("fiber: Book id is empty",
			zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetBook"})
		return c.Status(422).JSON(errorMessage{"Book ID is required"})
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		logger.Logger().Warn("fiber: Book id is not a number")
		return c.Status(422).JSON(errorMessage{"Book ID is not a number"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	book, err := s.Service.GetBook(ctx, intId)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(book)
}

//	@Summary		Get Books
//	@Description	Get Books with count and offset as optional params.
//	@Description	Count is the number of books to return.
//	@Description	Offset is the number of books to skip.
//	@Tags			Book
//	@Accept			application/json
//	@Produce		json
//	@Param			count	query		int	false	"Number of books to return"
//	@Param			offset	query		int	false	"Offset"
//	@Success		200		{object}	[]domain.Book
//	@Failure		400		{object}	fiber.errorMessage
//	@Router			/api/books [get]
func (s *Server) GetBooks(c *fiber.Ctx) error {
	var intCount, intOffset int
	var err error
	count := c.Query("count")
	if count != "" {
		intCount, err = strconv.Atoi(count)
		if err != nil {
			logger.Logger().Warn("fiber: Count is not a number",
				zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetBooks"})
			return c.Status(422).JSON(errorMessage{"Count is not a number"})
		}
	}
	offset := c.Query("offset")
	if offset != "" {
		intOffset, err = strconv.Atoi(offset)
		if err != nil {
			logger.Logger().Warn("fiber: Offset is not a number",
				zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetBooks"})
			return c.Status(422).JSON(errorMessage{"Offset is not a number"})
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	books, err := s.Service.GetBooks(ctx, intCount, intOffset)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(books)
}

//	@Summary		Create new book
//	@Description	Create new book.
//	@Tags			Book
//	@Accept			application/json
//	@Produce		json
//	@Param			book	body		domain.Book	true	"Book"
//	@Success		200		{object}	domain.Book
//	@Failure		400		{object}	fiber.errorMessage
//	@Router			/api/books [post]
func (s *Server) CreateBook(c *fiber.Ctx) error {
	var book domain.Book
	if err := c.BodyParser(&book); err != nil {
		logger.Logger().Warn("fiber: Invalid book",
			zapcore.Field{Key: "method", Type: zapcore.StringType, String: "CreateBook"})
		return c.Status(422).JSON(errorMessage{"Invalid book"})
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	book, err := s.Service.CreateBook(ctx, book)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(201).JSON(book)
}

//	@Summary		Update book
//	@Description	Update book with new values.
//	@Tags			Book
//	@Accept			application/json
//	@Produce		json
//	@Param			book	body		domain.Book	true	"Book"
//	@Success		200		{object}	domain.Book
//	@Failure		400		{object}	fiber.errorMessage
//	@Router			/api/books [patch]
func (s *Server) UpdateBook(c *fiber.Ctx) error {
	var book domain.Book
	if err := c.BodyParser(&book); err != nil {
		logger.Logger().Warn("fiber: Invalid book",
			zapcore.Field{Key: "method", Type: zapcore.StringType, String: "UpdateBook"})
		return c.Status(422).JSON(errorMessage{"Invalid book"})
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	book, err := s.Service.UpdateBook(ctx, book)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(book)
}

//	@Summary		Delete Book
//	@Description	Delete Book by ID.
//	@Tags			Book
//	@Accept			application/json
//	@Produce		json
//	@Param			id	query		int	true	"Book ID"
//	@Success		201	{object}	nil
//	@Failure		400	{object}	fiber.errorMessage
//	@Router			/api/books [delete]
func (s *Server) DeleteBook(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		logger.Logger().Warn("fiber: Book id is empty",
			zapcore.Field{Key: "method", Type: zapcore.StringType, String: "DeleteBook"})
		return c.Status(422).JSON(errorMessage{"Book ID is required"})
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		logger.Logger().Warn("fiber: Book id is not a number",
			zapcore.Field{Key: "method", Type: zapcore.StringType, String: "DeleteBook"})
		return c.Status(422).JSON(errorMessage{"Book ID is not a number"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = s.Service.DeleteBook(ctx, intId)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(204)
}
