package tracks

import (
	"net/http"

	trackactivities "github.com/gbrn7/music_catalog/internal/models/trackactivities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertTrackActivities(c *gin.Context) {
	ctx := c.Request.Context()

	var request trackactivities.TrackActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("userID")
	err := h.Service.UpsertTrackActivities(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
	}

	c.Status(http.StatusOK)
}
