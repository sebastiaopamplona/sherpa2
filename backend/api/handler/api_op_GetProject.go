package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetProject(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
