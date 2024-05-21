package handler

import (
	"github.com/Veer0x1/DumbCharades/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	movieService := service.NewMovieService(db)
	movieHandler := NewMovieHandler(movieService)

	app.Get("/api/movies/random", movieHandler.GetRandomMovie)
}

type MovieHandler struct {
	movieService *service.MovieService
}

func NewMovieHandler(movieService *service.MovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

func (h *MovieHandler) GetRandomMovie(c *fiber.Ctx) error {
	movie, err := h.movieService.GetRandomMovie()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(movie)
}
