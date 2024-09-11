package schemas

import "github.com/lib/pq"

type Movie struct {
	Id            int            `json:"id" db:"id"`
	Name          string         `json:"name" db:"name" binding:"required"`
	Description   *string        `json:"description" db:"description"`
	Director      *string        `json:"director" db:"director"`
	Producer      *string        `json:"producer" db:"producer"`
	Runtime       *int           `json:"runtime" db:"runtime"`
	Year          *int           `json:"year" db:"year"`
	Stars         *int8          `json:"stars" db:"stars"`
	Series        *int16         `json:"series" db:"series"`
	Seasons       *int16         `json:"seasons" db:"seasons"`
	Image         *string        `json:"image" db:"image"`
	VideoUrl      *string        `json:"video_url" db:"video_url"`
	Genres        pq.StringArray `json:"genres" db:"genres"`
	Screenshots   pq.StringArray `json:"screenshots" db:"screenshots"`
	SimilarMovies []Movie        `json:"similar_movies" db:"similar_movies"`
	WatchedAt     *string        `json:"watched_at" db:"watched_at"`
}

type AddMovieInfo struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Director    *string `json:"director"`
	Producer    *string `json:"producer"`
	Runtime     *int    `json:"runtime"`
	Year        *int    `json:"year"`
	Stars       *int8   `json:"stars"`
	Series      *int16  `json:"series"`
	Seasons     *int16  `json:"seasons"`
	Genres      []int   `json:"genres" binding:"required"`
}

type UpdateMovieInfo struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Director    *string `json:"director"`
	Producer    *string `json:"producer"`
	Runtime     *int    `json:"runtime"`
	Year        *int    `json:"year"`
	Stars       *int8   `json:"stars"`
	Series      *int16  `json:"series"`
	Seasons     *int16  `json:"seasons"`
}

func (m *UpdateMovieInfo) ToMap() map[string]any {
	movie := make(map[string]any, 9)

	if m.Name != nil {
		movie["name"] = m.Name
	}

	if m.Description != nil {
		movie["description"] = m.Description
	}

	if m.Director != nil {
		movie["director"] = m.Director
	}

	if m.Producer != nil {
		movie["producer"] = m.Producer
	}

	if m.Runtime != nil {
		movie["runtime"] = m.Runtime
	}

	if m.Year != nil {
		movie["year"] = m.Year
	}

	if m.Stars != nil {
		movie["stars"] = m.Stars
	}

	if m.Series != nil {
		movie["series"] = m.Series
	}

	if m.Seasons != nil {
		movie["seasons"] = m.Seasons
	}

	return movie
}

// swagger
type SwaggerMovieResponse struct {
	Id            int                    `json:"id" db:"id"`
	Name          string                 `json:"name" db:"name" binding:"required"`
	Description   *string                `json:"description" db:"description"`
	Director      *string                `json:"director" db:"director"`
	Producer      *string                `json:"producer" db:"producer"`
	Runtime       *int                   `json:"runtime" db:"runtime"`
	Year          *int                   `json:"year" db:"year"`
	Stars         *int8                  `json:"stars" db:"stars"`
	Series        *int16                 `json:"series" db:"series"`
	Seasons       *int16                 `json:"seasons" db:"seasons"`
	Image         *string                `json:"image" db:"image"`
	VideoUrl      *string                `json:"video_url" db:"video_url"`
	Genres        []string               `json:"genres" db:"genres"`
	Screenshots   []string               `json:"screenshots" db:"screenshots"`
	SimilarMovies []SwaggerMovieResponse `json:"similar_movies" db:"similar_movies"`
	WatchedAt     *string                `json:"watched_at" db:"watched_at"`
}

type Pagination struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type SwaggerPaginMovieResponse struct {
	Total      int                    `json:"total"`
	PageNum    int                    `json:"page_num"`
	PerPage    int                    `json:"per_page"`
	Pagination Pagination             `json:"pagination"`
	Data       []SwaggerMovieResponse `json:"data"`
}

type MovieCreatedResponse struct {
	Id int `json:"id"`
}
