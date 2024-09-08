package schemas

type Genre struct {
	Id    int    `json:"id" db:"id"`
	Genre string `json:"genre" db:"genre" binding:"required"`
}
