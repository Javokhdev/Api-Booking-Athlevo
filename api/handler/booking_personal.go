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

// BookingPersonalHandler handles requests related to Booking Personal.
type BookingPersonalHandler struct {
	service booking.BookingPersonalServiceClient
}

// NewBookingPersonalHandler creates a new BookingPersonalHandler.
func NewBookingPersonalHandler(grpcConn *grpc.ClientConn) *BookingPersonalHandler {
	return &BookingPersonalHandler{
		service: booking.NewBookingPersonalServiceClient(grpcConn),
	}
}

// CreateBookingPersonal godoc
// @Summary     Create a new Booking Personal
// @Description Create a new booking personal with the provided details.
// @Tags        BookingPersonal
// @Accept      json
// @Produce     json
// @Param       booking_personal body     booking.BookingPersonal true "Booking Personal details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.BookingPersonal
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-personal [post]
func (h *BookingPersonalHandler) CreateBookingPersonal(c *gin.Context) {
	var bookingPersonal booking.BookingPersonal
	if err := c.ShouldBindJSON(&bookingPersonal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the booking personal
	grpcResponse, err := h.service.CreateBookingPersonal(context.Background(), &booking.CreateBookingPersonalRequest{BookingPersonal: &bookingPersonal})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking personal " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// GetBookingPersonal godoc
// @Summary     Get Booking Personal by ID
// @Description Get details of a booking personal by its ID.
// @Tags        BookingPersonal
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Booking Personal ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.BookingPersonal
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-personal/{id} [get]
func (h *BookingPersonalHandler) GetBookingPersonal(c *gin.Context) {
	bookingPersonalID := c.Param("id")

	// Use gRPC to get the booking personal from the service
	grpcResponse, err := h.service.GetBookingPersonal(context.Background(), &booking.GetBookingPersonalRequest{Id: bookingPersonalID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Booking personal not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get booking personal " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateBookingPersonal godoc
// @Summary     Update Booking Personal
// @Description Update an existing booking personal.
// @Tags        BookingPersonal
// @Accept      json
// @Produce     json
// @Param       id                   path     string                             true "Booking Personal ID"
// @Param       booking_personal body     booking.BookingPersonal true "Updated booking personal"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.BookingPersonal
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-personal/{id} [put]
func (h *BookingPersonalHandler) UpdateBookingPersonal(c *gin.Context) {
	bookingPersonalID := c.Param("id")
	var bookingPersonal booking.BookingPersonal
	if err := c.ShouldBindJSON(&bookingPersonal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if bookingPersonal.Id != bookingPersonalID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Use gRPC to update the booking personal
	grpcResponse, err := h.service.UpdateBookingPersonal(context.Background(), &booking.UpdateBookingPersonalRequest{BookingPersonal: &bookingPersonal})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking personal " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// DeleteBookingPersonal godoc
// @Summary     Delete Booking Personal
// @Description Delete a booking personal by its ID.
// @Tags        BookingPersonal
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Booking Personal ID"
// @Security    ApiKeyAuth
// @Success     204     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-personal/{id} [delete]
func (h *BookingPersonalHandler) DeleteBookingPersonal(c *gin.Context) {
	bookingPersonalID := c.Param("id")

	// Call gRPC service to delete booking personal
	_, err := h.service.DeleteBookingPersonal(context.Background(), &booking.DeleteBookingPersonalRequest{Id: bookingPersonalID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Booking personal not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking personal " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Booking personal deleted successfully"})
}

// ListBookingPersonal godoc
// @Summary     List Booking Personal
// @Description Get a list of booking personal records.
// @Tags        BookingPersonal
// @Accept      json
// @Produce     json
// @Param        user_id             query    string false  "Filter by user ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListBookingPersonalResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/booking-personal [get]
func (h *BookingPersonalHandler) ListBookingPersonal(c *gin.Context) {
	// Get query parameters for filtering
	userID := c.Query("user_id")

	// Use gRPC to get the booking personal records from the service
	grpcResponse, err := h.service.ListBookingPersonal(context.Background(), &booking.ListBookingPersonalRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get booking personal records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
