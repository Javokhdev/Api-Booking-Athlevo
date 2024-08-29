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

// AccessHandler handles requests related to Access records.
type AccessHandler struct {
	service booking.AccessServiceClient
}

// NewAccessHandler creates a new AccessHandler.
func NewAccessHandler(grpcConn *grpc.ClientConn) *AccessHandler {
	return &AccessHandler{
		service: booking.NewAccessServiceClient(grpcConn),
	}
}

// CreateAccessPersonal godoc
// @Summary     Create a new Access Personal record
// @Description Create a new access personal record with the provided details.
// @Tags        Access
// @Accept      json
// @Produce     json
// @Param       access_personal body     booking.CreateAccessPersonalRequest true "Access Personal details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.AccessPersonal
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/access/personal [post]
func (h *AccessHandler) CreateAccessPersonal(c *gin.Context) {
	var req booking.CreateAccessPersonalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the access personal record
	grpcResponse, err := h.service.CreateAccessPersonal(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access personal record " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// ListAccessPersonal godoc
// @Summary     List Access Personal records
// @Description Get a list of access personal records for a specific booking.
// @Tags        Access
// @Accept      json
// @Produce     json
// @Param       booking_personal_id   path     string true "Booking Personal ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListAccessPersonalResponse
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/access/personal/{booking_personal_id} [get]
func (h *AccessHandler) ListAccessPersonal(c *gin.Context) {
	bookingPersonalID := c.Param("booking_personal_id")

	// Use gRPC to get the access personal records from the service
	grpcResponse, err := h.service.ListAccessPersonal(context.Background(), &booking.ListAccessPersonalRequest{
		BookingPersonalId: bookingPersonalID,
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Access personal records not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get access personal records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// CreateAccessGroup godoc
// @Summary     Create a new Access Group record
// @Description Create a new access group record with the provided details.
// @Tags        Access
// @Accept      json
// @Produce     json
// @Param       access_group body     booking.CreateAccessGroupRequest true "Access Group details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.AccessGroup
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/access/group [post]
func (h *AccessHandler) CreateAccessGroup(c *gin.Context) {
	var req booking.CreateAccessGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the access group record
	grpcResponse, err := h.service.CreateAccessGroup(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access group record " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// ListAccessGroup godoc
// @Summary     List Access Group records
// @Description Get a list of access group records for a specific booking.
// @Tags        Access
// @Accept      json
// @Produce     json
// @Param       booking_group_id   path     string true "Booking Group ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListAccessGroupResponse
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/access/group/{booking_group_id} [get]
func (h *AccessHandler) ListAccessGroup(c *gin.Context) {
	bookingGroupID := c.Param("booking_group_id")

	// Use gRPC to get the access group records from the service
	grpcResponse, err := h.service.ListAccessGroup(context.Background(), &booking.ListAccessGroupRequest{
		BookingGroupId: bookingGroupID,
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Access group records not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get access group records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// CreateAccessCoach godoc
// @Summary     Create a new Access Coach record
// @Description Create a new access coach record with the provided details.
// @Tags        Access
// @Accept      json
// @Produce     json
// @Param       access_coach body     booking.CreateAccessCoachRequest true "Access Coach details"
// @Security    ApiKeyAuth
// @Success     201     {object} booking.AccessCoach
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/access/coach [post]
func (h *AccessHandler) CreateAccessCoach(c *gin.Context) {
	var req booking.CreateAccessCoachRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body " + err.Error()})
		return
	}

	// Use gRPC to create the access coach record
	grpcResponse, err := h.service.CreateAccessCoach(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access coach record " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grpcResponse)
}

// ListAccessCoach godoc
// @Summary     List Access Coach records
// @Description Get a list of access coach records for a specific booking.
// @Tags        Access
// @Accept      json
// @Produce     json
// @Param       booking_coach_id   path     string true "Booking Coach ID"
// @Security    ApiKeyAuth
// @Success     200     {object} booking.ListAccessCoachResponse
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/access/coach/{booking_coach_id} [get]
func (h *AccessHandler) ListAccessCoach(c *gin.Context) {
	bookingCoachID := c.Param("booking_coach_id")

	// Use gRPC to get the access coach records from the service
	grpcResponse, err := h.service.ListAccessCoach(context.Background(), &booking.ListAccessCoachRequest{
		BookingCoachId: bookingCoachID,
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Access coach records not found " + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get access coach records " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
