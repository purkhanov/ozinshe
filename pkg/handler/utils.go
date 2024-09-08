package handler

import (
	"ozinshe/pkg/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getPaginationParams(c *gin.Context) (*utils.Pagination, error) {
	otherParams := ""

	movieNmae := c.Query("search")
	if movieNmae != "" {
		otherParams += "&search=" + movieNmae
	}

	genre := c.Query("genre")
	if genre != "" {
		otherParams += "&genre=" + genre

	}

	urlPath := strings.Split(c.Request.RequestURI, "?")

	pageNum := 1
	perPage := 20
	total := 0

	pageNumQuery := c.Query("page_num")
	if pageNumQuery != "" {
		pageN, err := strconv.Atoi(pageNumQuery)
		if err != nil {
			return nil, err
		}
		pageNum = pageN
	}

	perPageQuery := c.Query("per_page")
	if perPageQuery != "" {
		perPageN, err := strconv.Atoi(perPageQuery)
		if err != nil {
			return nil, err
		}
		perPage = perPageN
	}

	totalQuery := c.Query("total")
	if totalQuery != "" {
		totalNum, err := strconv.Atoi(totalQuery)
		if err != nil {
			return nil, err
		}
		total = totalNum
	}

	pag := &utils.Pagination{
		PageNum:     pageNum,
		PerPage:     perPage,
		Total:       total,
		UrlPath:     urlPath[0],
		OtherParams: otherParams,
	}

	return pag, nil
}
