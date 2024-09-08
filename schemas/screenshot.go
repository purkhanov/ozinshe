package schemas

type Screenshot struct {
	Id      int    `json:"id" binding:"required"`
	Link    string `json:"link"`
	MovieId int    `json:"movie_id" db:"movie_id"`
}
