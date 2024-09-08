package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizhationHeader = "Authorization"
	userCtxId            = "userId"
	userCtxIsAdmin       = "isAdmin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizhationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "not authenticated")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	userMap, err := h.services.Authorizhation.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtxId, userMap[userCtxId])
}

func (h *Handler) adminIdentity(c *gin.Context) {
	header := c.GetHeader(authorizhationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "not authenticated")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	userMap, err := h.services.Authorizhation.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if userMap[userCtxIsAdmin] != true {
		newErrorResponse(c, http.StatusUnauthorized, "you are not admin")
		return
	}

	c.Set(userCtxId, userMap[userCtxId])
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtxId)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("uer id not found")
	}

	return idInt, nil
}
