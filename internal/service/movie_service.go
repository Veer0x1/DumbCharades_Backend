package service

import (
	"github.com/Veer0x1/DumbCharades/internal/model"
	"github.com/Veer0x1/DumbCharades/internal/repository"
	"github.com/jmoiron/sqlx"
)


type MovieService struct {
	movieRepo *repository.MovieRepository
}

func NewMovieService(db *sqlx.DB) *MovieService {
	return &MovieService{movieRepo: repository.NewMovieRepository(db)}
}

func (s *MovieService) GetRandomMovie() (*model.Movie,error){
	return s.movieRepo.GetRandomMovie()
}