package handlers

import (
	"context"
	"net/http"

	"github.com/Athlevo/Api_booking_Athlevo/genproto/booking"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// AccessBetaHandler handles requests related to Access Beta.
type AccessBetaHandler struct {
	serviceBeta booking.AccessServiceBetaClient
}

// NewAccessBetaHandler creates a new AccessBetaHandler.
func NewAccessBetaHandler(grpcConn *grpc.ClientConn) *AccessBetaHandler {
	return &AccessBetaHandler{
		serviceBeta: booking.NewAccessServiceBetaClient(grpcConn),
	}
}

// CheckUserAccessBetaPersonal godoc
// @Summary     Check User Access (Beta) for Personal Subscription
// @Description Checks if a user has access to a specific sport hall for a personal subscription.
// @Tags        Access Beta
// @Accept      json
// @Produce     json
// @Param       sport_hall_id path     string true "Sport Hall ID"
// @Security    ApiKeyAuth
// @Success     200 {object} booking.AccessBetaPersonalResponse
// @Failure     400 {object} map[string]interface{}
// @Failure     500 {object} map[string]interface{}
// @Router      /v1/access-beta/personal/{sport_hall_id} [get]
func (h *AccessBetaHandler) CheckUserAccessBetaPersonal(c *gin.Context) {
	sportHallID := c.Param("sport_hall_id")

	// Get user ID from Gin context
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	req := &booking.AccessBetaPersonalRequest{
		UserId:      userID.(string),
		SportHallId: sportHallID,
	}

	// Use gRPC to check user access
	grpcResponse, err := h.serviceBeta.CheckUserAccess(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user access: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
