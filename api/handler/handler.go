package handlers

import (
	"github.com/Athlevo/Api_booking_Athlevo/config"
	"google.golang.org/grpc"
)

// Handler struct holds all the individual entity handlers.
type Handler struct {
	BookingPersonalHandler      *BookingPersonalHandler
	BookingGroupHandler         *BookingGroupHandler
	BookingCoachHandler         *BookingCoachHandler
	SubscriptionPersonalHandler *SubscriptionPersonalHandler
	SubscriptionGroupHandler    *SubscriptionGroupHandler
	SubscriptionCoachHandler    *SubscriptionCoachHandler
	AccessHandler               *AccessHandler
	AccessBetaHandler           *AccessBetaHandler
}

// NewHandler creates a new Handler instance with initialized individual handlers.
func NewHandler(grpcConn *grpc.ClientConn, cfg *config.Config) *Handler {
	return &Handler{
		BookingPersonalHandler:      NewBookingPersonalHandler(grpcConn),
		BookingGroupHandler:         NewBookingGroupHandler(grpcConn),
		BookingCoachHandler:         NewBookingCoachHandler(grpcConn),
		SubscriptionPersonalHandler: NewSubscriptionPersonalHandler(grpcConn),
		SubscriptionGroupHandler:    NewSubscriptionGroupHandler(grpcConn),
		SubscriptionCoachHandler:    NewSubscriptionCoachHandler(grpcConn),
		AccessHandler:               NewAccessHandler(grpcConn),
		AccessBetaHandler:           NewAccessBetaHandler(grpcConn),
	}
}
