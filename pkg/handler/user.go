package handler

import (
	"net/http"
	"ozinshe/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input schemas.UpdateUserInput
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
