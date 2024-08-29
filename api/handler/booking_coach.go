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

// BookingCoachHandler handles requests related to Booking Coach.
type BookingCoachHandler struct {
	service booking.BookingCoachServiceClient
}

// NewBookingCoachHandler creates a new BookingCoachHandler.
func NewBookingCoachHandler(grpcConn *grpc.ClientConn) *BookingCoachHandler {
	return &BookingCoachHandler{
		service: booking.NewBookingCoachServiceClient(grpcConn),
	}
}

// CreateBookingCoach godoc
// @Summary     Create a new Booking Coach
// @Description Create a new booking coach with the provided details.
// @Tags        BookingCoach
// @Accept      json
// @Produce     json
// @Param       booking_coach body     booking.BookingCoach true "Booking Coach details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.BookingCoach
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-coach [post]
func (h *BookingCoachHandler) CreateBookingCoach(c *gin.Context) {
	var bookingCoach booking.BookingCoach
	if err := c.ShouldBindJSON(&bookingCoach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the booking coach
	grpcResponse, err := h.service.CreateBookingCoach(context.Background(), &booking.CreateBookingCoachRequest{BookingCoach: &bookingCoach})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking coach " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// GetBookingCoach godoc
// @Summary     Get Booking Coach by ID
// @Description Get details of a booking coach by its ID.
// @Tags        BookingCoach
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Booking Coach ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.BookingCoach
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-coach/{id} [get]
func (h *BookingCoachHandler) GetBookingCoach(c *gin.Context) {
	bookingCoachID := c.Param("id")

	// Use gRPC to get the booking coach from the service
	grpcResponse, err := h.service.GetBookingCoach(context.Background(), &booking.GetBookingCoachRequest{Id: bookingCoachID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Booking coach not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get booking coach " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateBookingCoach godoc
// @Summary     Update Booking Coach
// @Description Update an existing booking coach.
// @Tags        BookingCoach
// @Accept      json
// @Produce     json
// @Param       id                   path     string                             true "Booking Coach ID"
// @Param       booking_coach body     booking.BookingCoach true "Updated booking coach"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.BookingCoach
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-coach/{id} [put]
func (h *BookingCoachHandler) UpdateBookingCoach(c *gin.Context) {
	bookingCoachID := c.Param("id")
	var bookingCoach booking.BookingCoach
	if err := c.ShouldBindJSON(&bookingCoach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if bookingCoach.Id != bookingCoachID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Use gRPC to update the booking coach
	grpcResponse, err := h.service.UpdateBookingCoach(context.Background(), &booking.UpdateBookingCoachRequest{BookingCoach: &bookingCoach})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking coach " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// DeleteBookingCoach godoc
// @Summary     Delete Booking Coach
// @Description Delete a booking coach by its ID.
// @Tags        BookingCoach
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Booking Coach ID"
// @Security    ApiKeyAuth
// @Success     204     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-coach/{id} [delete]
func (h *BookingCoachHandler) DeleteBookingCoach(c *gin.Context) {
	bookingCoachID := c.Param("id")

	// Call gRPC service to delete booking coach
	_, err := h.service.DeleteBookingCoach(context.Background(), &booking.DeleteBookingCoachRequest{Id: bookingCoachID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Booking coach not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking coach " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Booking coach deleted successfully"})
}

// ListBookingCoach godoc
// @Summary     List Booking Coach
// @Description Get a list of booking coach records.
// @Tags        BookingCoach
// @Accept      json
// @Produce     json
// @Param        user_id             query    string false  "Filter by user ID"
// @Param        subscription_id    query    string false  "Filter by subscription ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListBookingCoachResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-coach [get]
func (h *BookingCoachHandler) ListBookingCoach(c *gin.Context) {
	// Get query parameters for filtering
	userID := c.Query("user_id")
	subscriptionID := c.Query("subscription_id")

	// Use gRPC to get the booking coach records from the service
	grpcResponse, err := h.service.ListBookingCoach(context.Background(), &booking.ListBookingCoachRequest{
		UserId:         userID,
		SubscriptionId: subscriptionID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get booking coach records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
