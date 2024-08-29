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

// SubscriptionCoachHandler handles requests related to Subscription Coach.
type SubscriptionCoachHandler struct {
	service booking.SubscriptionCoachServiceClient
}

// NewSubscriptionCoachHandler creates a new SubscriptionCoachHandler.
func NewSubscriptionCoachHandler(grpcConn *grpc.ClientConn) *SubscriptionCoachHandler {
	return &SubscriptionCoachHandler{
		service: booking.NewSubscriptionCoachServiceClient(grpcConn),
	}
}

// CreateSubscriptionCoach godoc
// @Summary     Create a new Subscription Coach
// @Description Create a new subscription coach with the provided details.
// @Tags        SubscriptionCoach
// @Accept      json
// @Produce     json
// @Param       subscription_coach body     booking.SubscriptionCoach true "Subscription Coach details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.SubscriptionCoach
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-coach [post]
func (h *SubscriptionCoachHandler) CreateSubscriptionCoach(c *gin.Context) {
	var subscriptionCoach booking.SubscriptionCoach
	if err := c.ShouldBindJSON(&subscriptionCoach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the subscription coach
	grpcResponse, err := h.service.CreateSubscriptionCoach(context.Background(), &booking.CreateSubscriptionCoachRequest{SubscriptionCoach: &subscriptionCoach})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription coach " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// GetSubscriptionCoach godoc
// @Summary     Get Subscription Coach by ID
// @Description Get details of a subscription coach by its ID.
// @Tags        SubscriptionCoach
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Subscription Coach ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.SubscriptionCoach
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-coach/{id} [get]
func (h *SubscriptionCoachHandler) GetSubscriptionCoach(c *gin.Context) {
	subscriptionCoachID := c.Param("id")

	// Use gRPC to get the subscription coach from the service
	grpcResponse, err := h.service.GetSubscriptionCoach(context.Background(), &booking.GetSubscriptionCoachRequest{Id: subscriptionCoachID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Subscription coach not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription coach " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateSubscriptionCoach godoc
// @Summary     Update Subscription Coach
// @Description Update an existing subscription coach.
// @Tags        SubscriptionCoach
// @Accept      json
// @Produce     json
// @Param       id                   path     string                             true "Subscription Coach ID"
// @Param       subscription_coach body     booking.SubscriptionCoach true "Updated subscription coach"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.SubscriptionCoach
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-coach/{id} [put]
func (h *SubscriptionCoachHandler) UpdateSubscriptionCoach(c *gin.Context) {
	subscriptionCoachID := c.Param("id")
	var subscriptionCoach booking.SubscriptionCoach
	if err := c.ShouldBindJSON(&subscriptionCoach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if subscriptionCoach.Id != subscriptionCoachID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Use gRPC to update the subscription coach
	grpcResponse, err := h.service.UpdateSubscriptionCoach(context.Background(), &booking.UpdateSubscriptionCoachRequest{SubscriptionCoach: &subscriptionCoach})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subscription coach " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// DeleteSubscriptionCoach godoc
// @Summary     Delete Subscription Coach
// @Description Delete a subscription coach by its ID.
// @Tags        SubscriptionCoach
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Subscription Coach ID"
// @Security    ApiKeyAuth
// @Success     204     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-coach/{id} [delete]
func (h *SubscriptionCoachHandler) DeleteSubscriptionCoach(c *gin.Context) {
	subscriptionCoachID := c.Param("id")

	// Call gRPC service to delete subscription coach
	_, err := h.service.DeleteSubscriptionCoach(context.Background(), &booking.DeleteSubscriptionCoachRequest{Id: subscriptionCoachID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Subscription coach not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subscription coach " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Subscription coach deleted successfully"})
}

// ListSubscriptionCoach godoc
// @Summary     List Subscription Coach
// @Description Get a list of subscription coach records.
// @Tags        SubscriptionCoach
// @Accept      json
// @Produce     json
// @Param        gym_id             query    string false  "Filter by gym ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListSubscriptionCoachResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-coach [get]
func (h *SubscriptionCoachHandler) ListSubscriptionCoach(c *gin.Context) {
	// Get query parameters for filtering
	gymID := c.Query("gym_id")

	// Use gRPC to get the subscription coach records from the service
	grpcResponse, err := h.service.ListSubscriptionCoach(context.Background(), &booking.ListSubscriptionCoachRequest{
		GymId: gymID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription coach records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
