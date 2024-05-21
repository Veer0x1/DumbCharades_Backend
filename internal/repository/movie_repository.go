package repository

import (
	"github.com/Veer0x1/DumbCharades/internal/model"
	"github.com/jmoiron/sqlx"
)

type MovieRepository struct {
	db *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) *MovieRepository{
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetRandomMovie() (*model.Movie, error) {
	var movie model.Movie
	err := r.db.Get(&movie, "SELECT * FROM movies ORDER BY random() LIMIT 1")
	if err != nil {
		return nil, err
	}
	return &movie, nil
}