package handler

import (
	"net/http"
	"ozinshe/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {array} schemas.UserResponse "Successful"
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /admin/users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.User.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get user
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} schemas.UserResponse "Successful"
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /admin/users/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
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
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} schemas.UserInput "Successful"
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /admin/users/{id} [put]
func (h *Handler) updateUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
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
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 {string} Successful
// @Failure 400 {string} Invalid input
// @Security ApiKeyAuth
// @Router /admin/users/{id} [delete]
func (h *Handler) deleteUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.User.DeleteUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{Status: "deleted"})
}
