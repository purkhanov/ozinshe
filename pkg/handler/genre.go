package handler

import (
	"net/http"
	"ozinshe/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllGenres(c *gin.Context) {
	genres, err := h.services.Genre.GetAllGenre()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, genres)
}

func (h *Handler) addGenre(c *gin.Context) {
	var genre schemas.Genre

	if err := c.BindJSON(&genre); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	genreId, err := h.services.Genre.AddGenre(genre.Genre)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{"id": genreId})
}

func (h *Handler) updateGenre(c *gin.Context) {
	var input schemas.Genre

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	genreId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Genre.UpdateGenre(genreId, input.Genre)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, statusResponse{Status: "accepted"})
}

func (h *Handler) deleteGenre(c *gin.Context) {
	genreId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Genre.DeleteGenre(genreId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{Status: "deleted"})
}
