package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetSprint(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
