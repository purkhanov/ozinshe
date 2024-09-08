package handler

import (
	"net/http"
	"os"
	"ozinshe/schemas"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getMovies(c *gin.Context) {
	movieNmae := c.Query("search")
	if movieNmae != "" {
		h.searchByName(c)
		return
	}

	genre := c.Query("genre")
	if genre != "" {
		h.searchByGenre(c)
		return
	}

	h.getAllMovies(c)
}

func (h *Handler) getAllMovies(c *gin.Context) {
	paginParams, err := getPaginationParams(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid query params")
	}

	movies, err := h.services.Movie.GetAll(paginParams)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) searchByName(c *gin.Context) {
	paginParams, err := getPaginationParams(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid query params")
	}

	movieNmae := c.Query("search")

	movies, err := h.services.Movie.SearchByName(movieNmae, paginParams)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) searchByGenre(c *gin.Context) {
	paginParams, err := getPaginationParams(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid query params")
	}
	genre := c.Query("genre")

	movies, err := h.services.Movie.SearchByGenre(genre, paginParams)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) getMovieById(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	movie, err := h.services.GetById(movieId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (h *Handler) addMovieInfo(c *gin.Context) {
	var movieInfo schemas.AddMovieInfo

	if err := c.BindJSON(&movieInfo); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Movie.AddMovie(movieInfo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{"id": id})
}

func (h *Handler) uploadMovie(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	url := "/movies/" + file.Filename

	err = h.services.Movie.UploadMovie(movieId, url)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.SaveUploadedFile(file, "./data"+url); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{"id": movieId})
}

func (h *Handler) getScreenshots(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	screens, err := h.services.Movie.GetScreenshots(movieId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, screens)
}

func (h *Handler) uploadScreenshot(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	url := "/screenshots/" + file.Filename

	_, err = h.services.Movie.AddScreenshot(movieId, url)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.SaveUploadedFile(file, "./data"+url); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Movie.AddScreenshot(movieId, url)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{"id": id})
}

func (h *Handler) updateMovie(c *gin.Context) {
	var movieInput schemas.UpdateMovieInfo

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.BindJSON(&movieInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Movie.UpdateMovie(movieId, movieInput); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, statusResponse{Status: "updated"})
}

func (h *Handler) deleteMovie(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Movie.DeleteMovie(movieId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{Status: "deleted"})
}

func (h *Handler) deleteScreenshot(c *gin.Context) {
	screenIdStr := c.Query("screenshot_id")
	if screenIdStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "param screenshot_id is required")
		return
	}

	screenId, err := strconv.Atoi(screenIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid screenshot_id param")
		return
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	filePath, err := h.services.Movie.DeleteScreenshot(movieId, screenId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := os.Remove("./data" + filePath); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{Status: "deleted"})
}
