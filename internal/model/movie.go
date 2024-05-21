package model

type Movie struct {
	Imdb_ID       string  `db:"imdb_id" json:"imdb_id"`
	Title         string  `db:"title" json:"title"`
	PosterURL     string  `db:"poster_path" json:"poster_path"`
	Wikipedia_URL string  `db:"wiki_link" json:"wiki_link"`
	ReleaseYear   string  `db:"year_of_release" json:"year_of_release"`
	Genre         string  `db:"genres" json:"genres"`
	Imdb_rating   float64 `db:"imdb_rating" json:"imdb_rating"`
	Imdb_votes    float64 `db:"imdb_votes" json:"imdb_votes"`
	Summary       string  `db:"summary" json:"summary"`
	Actors        string  `db:"actors" json:"actors"`
}
