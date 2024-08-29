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

// SubscriptionGroupHandler handles requests related to Subscription Group.
type SubscriptionGroupHandler struct {
	service booking.SubscriptionGroupServiceClient
}

// NewSubscriptionGroupHandler creates a new SubscriptionGroupHandler.
func NewSubscriptionGroupHandler(grpcConn *grpc.ClientConn) *SubscriptionGroupHandler {
	return &SubscriptionGroupHandler{
		service: booking.NewSubscriptionGroupServiceClient(grpcConn),
	}
}

// CreateSubscriptionGroup godoc
// @Summary     Create a new Subscription Group
// @Description Create a new subscription group with the provided details.
// @Tags        SubscriptionGroup
// @Accept      json
// @Produce     json
// @Param       subscription_group body     booking.SubscriptionGroup true "Subscription Group details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.SubscriptionGroup
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-group [post]
func (h *SubscriptionGroupHandler) CreateSubscriptionGroup(c *gin.Context) {
	var subscriptionGroup booking.SubscriptionGroup
	if err := c.ShouldBindJSON(&subscriptionGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the subscription group
	grpcResponse, err := h.service.CreateSubscriptionGroup(context.Background(), &booking.CreateSubscriptionGroupRequest{SubscriptionGroup: &subscriptionGroup})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription group " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// GetSubscriptionGroup godoc
// @Summary     Get Subscription Group by ID
// @Description Get details of a subscription group by its ID.
// @Tags        SubscriptionGroup
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Subscription Group ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.SubscriptionGroup
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-group/{id} [get]
func (h *SubscriptionGroupHandler) GetSubscriptionGroup(c *gin.Context) {
	subscriptionGroupID := c.Param("id")

	// Use gRPC to get the subscription group from the service
	grpcResponse, err := h.service.GetSubscriptionGroup(context.Background(), &booking.GetSubscriptionGroupRequest{Id: subscriptionGroupID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Subscription group not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription group " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateSubscriptionGroup godoc
// @Summary     Update Subscription Group
// @Description Update an existing subscription group.
// @Tags        SubscriptionGroup
// @Accept      json
// @Produce     json
// @Param       id                   path     string                             true "Subscription Group ID"
// @Param       subscription_group body     booking.SubscriptionGroup true "Updated subscription group"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.SubscriptionGroup
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-group/{id} [put]
func (h *SubscriptionGroupHandler) UpdateSubscriptionGroup(c *gin.Context) {
	subscriptionGroupID := c.Param("id")
	var subscriptionGroup booking.SubscriptionGroup
	if err := c.ShouldBindJSON(&subscriptionGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if subscriptionGroup.Id != subscriptionGroupID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Use gRPC to update the subscription group
	grpcResponse, err := h.service.UpdateSubscriptionGroup(context.Background(), &booking.UpdateSubscriptionGroupRequest{SubscriptionGroup: &subscriptionGroup})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subscription group " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// DeleteSubscriptionGroup godoc
// @Summary     Delete Subscription Group
// @Description Delete a subscription group by its ID.
// @Tags        SubscriptionGroup
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Subscription Group ID"
// @Security    ApiKeyAuth
// @Success     204     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-group/{id} [delete]
func (h *SubscriptionGroupHandler) DeleteSubscriptionGroup(c *gin.Context) {
	subscriptionGroupID := c.Param("id")

	// Call gRPC service to delete subscription group
	_, err := h.service.DeleteSubscriptionGroup(context.Background(), &booking.DeleteSubscriptionGroupRequest{Id: subscriptionGroupID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Subscription group not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subscription group " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Subscription group deleted successfully"})
}

// ListSubscriptionGroup godoc
// @Summary     List Subscription Group
// @Description Get a list of subscription group records.
// @Tags        SubscriptionGroup
// @Accept      json
// @Produce     json
// @Param        gym_id             query    string false  "Filter by gym ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListSubscriptionGroupResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/subscription-group [get]
func (h *SubscriptionGroupHandler) ListSubscriptionGroup(c *gin.Context) {
	// Get query parameters for filtering
	gymID := c.Query("gym_id")

	// Use gRPC to get the subscription group records from the service
	grpcResponse, err := h.service.ListSubscriptionGroup(context.Background(), &booking.ListSubscriptionGroupRequest{
		GymId: gymID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription group records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
