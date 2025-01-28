package memberships

import (
	"net/http"

	"github.com/gbrn7/music_catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {

	var request memberships.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.Service.SignUp(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
