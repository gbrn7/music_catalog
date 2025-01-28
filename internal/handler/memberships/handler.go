package memberships

import (
	"github.com/gbrn7/music_catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go  -destination=handler_mock_test.go -package=memberships
type service interface {
	SignUp(request memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine
	Service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.POST("/sign-up", h.SignUp)
}
