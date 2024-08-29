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

// BookingGroupHandler handles requests related to Booking Group.
type BookingGroupHandler struct {
	service booking.BookingGroupServiceClient
}

// NewBookingGroupHandler creates a new BookingGroupHandler.
func NewBookingGroupHandler(grpcConn *grpc.ClientConn) *BookingGroupHandler {
	return &BookingGroupHandler{
		service: booking.NewBookingGroupServiceClient(grpcConn),
	}
}

// CreateBookingGroup godoc
// @Summary     Create a new Booking Group
// @Description Create a new booking group with the provided details.
// @Tags        BookingGroup
// @Accept      json
// @Produce     json
// @Param       booking_group body     booking.BookingGroup true "Booking Group details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.BookingGroup
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-group [post]
func (h *BookingGroupHandler) CreateBookingGroup(c *gin.Context) {
	var bookingGroup booking.BookingGroup
	if err := c.ShouldBindJSON(&bookingGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the booking group
	grpcResponse, err := h.service.CreateBookingGroup(context.Background(), &booking.CreateBookingGroupRequest{BookingGroup: &bookingGroup})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking group " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// GetBookingGroup godoc
// @Summary     Get Booking Group by ID
// @Description Get details of a booking group by its ID.
// @Tags        BookingGroup
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Booking Group ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.BookingGroup
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-group/{id} [get]
func (h *BookingGroupHandler) GetBookingGroup(c *gin.Context) {
	bookingGroupID := c.Param("id")

	// Use gRPC to get the booking group from the service
	grpcResponse, err := h.service.GetBookingGroup(context.Background(), &booking.GetBookingGroupRequest{Id: bookingGroupID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Booking group not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get booking group " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateBookingGroup godoc
// @Summary     Update Booking Group
// @Description Update an existing booking group.
// @Tags        BookingGroup
// @Accept      json
// @Produce     json
// @Param       id                   path     string                             true "Booking Group ID"
// @Param       booking_group body     booking.BookingGroup true "Updated booking group"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.BookingGroup
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-group/{id} [put]
func (h *BookingGroupHandler) UpdateBookingGroup(c *gin.Context) {
	bookingGroupID := c.Param("id")
	var bookingGroup booking.BookingGroup
	if err := c.ShouldBindJSON(&bookingGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if bookingGroup.Id != bookingGroupID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Use gRPC to update the booking group
	grpcResponse, err := h.service.UpdateBookingGroup(context.Background(), &booking.UpdateBookingGroupRequest{BookingGroup: &bookingGroup})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking group " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// DeleteBookingGroup godoc
// @Summary     Delete Booking Group
// @Description Delete a booking group by its ID.
// @Tags        BookingGroup
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Booking Group ID"
// @Security    ApiKeyAuth
// @Success     204     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-group/{id} [delete]
func (h *BookingGroupHandler) DeleteBookingGroup(c *gin.Context) {
	bookingGroupID := c.Param("id")

	// Call gRPC service to delete booking group
	_, err := h.service.DeleteBookingGroup(context.Background(), &booking.DeleteBookingGroupRequest{Id: bookingGroupID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Booking group not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking group " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Booking group deleted successfully"})
}

// ListBookingGroup godoc
// @Summary     List Booking Group
// @Description Get a list of booking group records.
// @Tags        BookingGroup
// @Accept      json
// @Produce     json
// @Param        user_id             query    string false  "Filter by user ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListBookingGroupResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-group [get]
func (h *BookingGroupHandler) ListBookingGroup(c *gin.Context) {
	// Get query parameters for filtering
	userID := c.Query("user_id")

	// Use gRPC to get the booking group records from the service
	grpcResponse, err := h.service.ListBookingGroup(context.Background(), &booking.ListBookingGroupRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get booking group records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
