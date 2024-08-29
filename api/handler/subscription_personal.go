package handlers

import (
	"context"
	"net/http"

	"github.com/Athlevo/Api_booking_Athlevo/genproto/booking"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SubscriptionPersonalHandler handles requests related to Subscription Personal.
type SubscriptionPersonalHandler struct {
	service booking.SubscriptionPersonalServiceClient
}

// NewSubscriptionPersonalHandler creates a new SubscriptionPersonalHandler.
func NewSubscriptionPersonalHandler(grpcConn *grpc.ClientConn) *SubscriptionPersonalHandler {
	return &SubscriptionPersonalHandler{
		service: booking.NewSubscriptionPersonalServiceClient(grpcConn),
	}
}

// CreateSubscriptionPersonal godoc
// @Summary     Create a new Subscription Personal
// @Description Create a new subscription personal with the provided details.
// @Tags        SubscriptionPersonal
// @Accept      json
// @Produce     json
// @Param       subscription_personal body     booking.SubscriptionPersonal true "Subscription Personal details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.SubscriptionPersonal
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-personal [post]
func (h *SubscriptionPersonalHandler) CreateSubscriptionPersonal(c *gin.Context) {
	var subscriptionPersonal booking.SubscriptionPersonal
	if err := c.ShouldBindJSON(&subscriptionPersonal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the subscription personal
	grpcResponse, err := h.service.CreateSubscriptionPersonal(context.Background(), &booking.CreateSubscriptionPersonalRequest{SubscriptionPersonal: &subscriptionPersonal})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription personal " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// GetSubscriptionPersonal godoc
// @Summary     Get Subscription Personal by ID
// @Description Get details of a subscription personal by its ID.
// @Tags        SubscriptionPersonal
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Subscription Personal ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.SubscriptionPersonal
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-personal/{id} [get]
func (h *SubscriptionPersonalHandler) GetSubscriptionPersonal(c *gin.Context) {
	subscriptionPersonalID := c.Param("id")

	// Use gRPC to get the subscription personal from the service
	grpcResponse, err := h.service.GetSubscriptionPersonal(context.Background(), &booking.GetSubscriptionPersonalRequest{Id: subscriptionPersonalID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Subscription personal not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription personal " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateSubscriptionPersonal godoc
// @Summary     Update Subscription Personal
// @Description Update an existing subscription personal.
// @Tags        SubscriptionPersonal
// @Accept      json
// @Produce     json
// @Param       id                   path     string                             true "Subscription Personal ID"
// @Param       subscription_personal body     booking.SubscriptionPersonal true "Updated subscription personal"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.SubscriptionPersonal
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-personal/{id} [put]
func (h *SubscriptionPersonalHandler) UpdateSubscriptionPersonal(c *gin.Context) {
	subscriptionPersonalID := c.Param("id")
	var subscriptionPersonal booking.SubscriptionPersonal
	if err := c.ShouldBindJSON(&subscriptionPersonal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if subscriptionPersonal.Id != subscriptionPersonalID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Use gRPC to update the subscription personal
	grpcResponse, err := h.service.UpdateSubscriptionPersonal(context.Background(), &booking.UpdateSubscriptionPersonalRequest{SubscriptionPersonal: &subscriptionPersonal})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subscription personal " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// DeleteSubscriptionPersonal godoc
// @Summary     Delete Subscription Personal
// @Description Delete a subscription personal by its ID.
// @Tags        SubscriptionPersonal
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Subscription Personal ID"
// @Security    ApiKeyAuth
// @Success     204     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-personal/{id} [delete]
func (h *SubscriptionPersonalHandler) DeleteSubscriptionPersonal(c *gin.Context) {
	subscriptionPersonalID := c.Param("id")

	// Call gRPC service to delete subscription personal
	_, err := h.service.DeleteSubscriptionPersonal(context.Background(), &booking.DeleteSubscriptionPersonalRequest{Id: subscriptionPersonalID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Subscription personal not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subscription personal " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Subscription personal deleted successfully"})
}

// ListSubscriptionPersonal godoc
// @Summary     List Subscription Personal
// @Description Get a list of subscription personal records.
// @Tags        SubscriptionPersonal
// @Accept      json
// @Produce     json
// @Param        gym_id             query    string false  "Filter by gym ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListSubscriptionPersonalResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-personal [get]
func (h *SubscriptionPersonalHandler) ListSubscriptionPersonal(c *gin.Context) {
	// Get query parameters for filtering
	gymID := c.Query("gym_id")

	// Use gRPC to get the subscription personal records from the service
	grpcResponse, err := h.service.ListSubscriptionPersonal(context.Background(), &booking.ListSubscriptionPersonalRequest{
		GymId: gymID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription personal records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
