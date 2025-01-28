package memberships

import (
	"net/http"

	"github.com/gbrn7/music_catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var request memberships.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.Service.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberships.LoginResponse{
		AccessToken: accessToken,
	})
}
