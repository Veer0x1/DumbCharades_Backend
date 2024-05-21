package model

type Movie struct {
    ID          int    `db:"id" json:"id"`
    Title       string `db:"title" json:"title"`
    ReleaseYear int    `db:"release_year" json:"release_year"`
    Genre       string `db:"genre" json:"genre"`
    Director    string `db:"director" json:"director"`
    Cast        string `db:"cast" json:"cast"`
    PosterURL   string `db:"poster_url" json:"poster_url"`
}
