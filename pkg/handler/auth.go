package handler

import (
	"net/http"
	"ozinshe/schemas"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

// @Summary Sign Up
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schemas.UserInput true "body json"
// @Success 201 {object} schemas.UserCreatedSwaggerResponse "Successful sign up"
// @Failure 400 {string} Invalid input
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input schemas.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorizhation.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{"id": id})
}

type signInInput struct {
	Email    string
	Password string
}

// @Summary Sign In
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schemas.UserSignIn true "body json"
// @Success 200 {object} schemas.UserTokenSwaggerResponse "Successful"
// @Failure 400 {string} Invalid input
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorizhation.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{"token": token})
}
