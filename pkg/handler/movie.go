package handler

import (
	"net/http"
	"os"
	"ozinshe/schemas"

	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get movies
// @Tags movie
// @Accept json
// @Produce json
// @Param search query string false "найдет фильм по имени если не указан найдет все"
// @Param genre query string false "найдет фильм по жанру если не указан найдет все"
// @Param page_num query int false "номер страницы по умолчнию 1"
// @Param per_page query int false "количество фильмов в одном ответе по умолчанию 20"
// @Success 200 {object} schemas.SwaggerPaginMovieResponse "Successful"
// @Failure 400 {string} Invalid input
// @Router /movies [get]
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

// @Summary Get movie by ID
// @Tags movie
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} schemas.SwaggerMovieResponse "Successful"
// @Failure 400 {string} Invalid input
// @Router /movies/{id} [get]
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

// @Summary Add movie
// @Tags admin
// @Accept json
// @Produce json
// @Param request body schemas.AddMovieInfo true "body json"
// @Success 201 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /admin/movies [post]
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

// @Summary Upload movie
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Movie ID"
// @Param file formData file true "Movie file"
// @Success 201 {string} Successful
// @Failure 400 {string} Invalid input
// @Failure 500 {string} Server error
// @Security ApiKeyAuth
// @Router /admin/movies/{id} [post]
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

// @Summary Get movie screenshots
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} schemas.Screenshot "Successful"
// @Failure 400 {string} Invalid input
// @Failure 500 {string} Server error
// @Security ApiKeyAuth
// @Router /movies/screenshots/{id} [get]
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

// @Summary Upload screenshot
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Movie ID"
// @Param file formData file true "Movie screenshot"
// @Success 201 {string} Successful
// @Failure 400 {string} Invalid input
// @Failure 500 {string} Server error
// @Security ApiKeyAuth
// @Router /admin/screenshots/{id} [post]
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

// @Summary Update movie
// @Tags admin
// @Accept json
// @Produce json
// @Param request body schemas.UpdateMovieInfo true "body json"
// @Success 201 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /admin/movies [put]
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

// @Summary Delete movie
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 204 {string} Successful
// @Failure 400 {string} Invalid input
// @Failure 500 {string} Server error
// @Security ApiKeyAuth
// @Router /admin/movie/{id} [delete]
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

// @Summary Delete screenshot
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Screenshot ID"
// @Success 204 {string} Successful
// @Failure 400 {string} Invalid input
// @Failure 500 {string} Server error
// @Security ApiKeyAuth
// @Router /admin/screenshots/{id} [delete]
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
