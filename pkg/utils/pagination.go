package utils

import (
	"fmt"
	"ozinshe/schemas"
)

type PaginationMovieResponse struct {
	Total      int             `json:"total"`
	PageNUm    int             `json:"page_num"`
	PerPage    int             `json:"per_page"`
	Pagination map[string]any  `json:"pagination"`
	Data       []schemas.Movie `json:"data"`
}

type Pagination struct {
	PageNum     int
	PerPage     int
	Total       int
	UrlPath     string
	OtherParams string
}

func (p *Pagination) Paginate() map[string]any {
	start := (p.PageNum - 1) * p.PerPage
	end := start + p.PerPage

	pagination := make(map[string]any, 2)
	pagination["next"] = nil
	pagination["prev"] = nil

	if end >= p.Total {
		if p.PageNum > 1 {
			pagination["prev"] = fmt.Sprintf(
				"%s?page_num=%d&per_page=%d&total=%d%s",
				p.UrlPath, p.PageNum-1, p.PerPage, p.Total, p.OtherParams,
			)
		}
	} else {
		if p.PageNum > 1 {
			pagination["prev"] = fmt.Sprintf(
				"%s?page_num=%d&per_page=%d&total=%d%s",
				p.UrlPath, p.PageNum-1, p.PerPage, p.Total, p.OtherParams,
			)
		}

		pagination["next"] = fmt.Sprintf(
			"%s?page_num=%d&per_page=%d&total=%d%s",
			p.UrlPath, p.PageNum+1, p.PerPage, p.Total, p.OtherParams,
		)
	}

	return pagination
}
