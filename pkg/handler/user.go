package handler

import (
	"net/http"
	"ozinshe/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} schemas.UserResponse "Successful"
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user [get]
func (h *Handler) getUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	user, err := h.services.User.GetUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Update user
// @Tags user
// @Accept json
// @Produce json
// @Param request body schemas.UserInput true "body json"
// @Success 201 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user [put]
func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input schemas.UserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.UpdateUser(userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, statusResponse{Status: "accepted"})
}

// @Summary Delete user
// @Tags user
// @Accept json
// @Produce json
// @Success 204 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	err = h.services.User.DeleteUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{Status: "deleted"})
}

// @Summary Get user favorite movies
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} schemas.SwaggerMovieResponse "Successful"
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user/favorites [get]
func (h *Handler) getFavoriteMovies(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	movies, err := h.services.User.GetFavoriteMovies(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movies)
}

// @Summary Add user favorite movie
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 201 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user/favorites/{id} [post]
func (h *Handler) addFavoriteMovie(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	id, err := h.services.User.AddFavoriteMovie(userId, movieId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, map[string]int{"id": id})
}

// @Summary Delete user favorite movie
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 201 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user/favorites/{id} [delete]
func (h *Handler) deleteFavoriteMovie(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.User.DeleteFavoriteMovie(userId, movieId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{Status: "deleted"})
}

// @Summary Get user watched movies
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} schemas.SwaggerMovieResponse "Successful"
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user/watched-movies [get]
func (h *Handler) getWtchedMovies(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	movies, err := h.services.User.GetWatchedMovies(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, movies)
}

// @Summary Add user watched movie
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 201 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user/watched-movies/{id} [post]
func (h *Handler) addWtchedMovies(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	id, err := h.services.User.AddWatchedMovie(userId, movieId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, map[string]int{"id": id})
}

// @Summary Delete user watched movie
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 204 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /user/watched-movies/{id} [delete]
func (h *Handler) deleteWtchedMovies(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.User.DeleteWatchedMovie(userId, movieId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{Status: "deleted"})
}
