package api

import (
	"github.com/Athlevo/Api_booking_Athlevo/api/auth"
	_ "github.com/Athlevo/Api_booking_Athlevo/api/docs"
	handlers "github.com/Athlevo/Api_booking_Athlevo/api/handler"

	"github.com/Athlevo/Api_booking_Athlevo/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
																											
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath  /v1
func NewRouter(grpcConn *grpc.ClientConn, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler := handlers.NewHandler(grpcConn, cfg)

	// API versioning
	v1 := router.Group("/v1")

	ca := auth.CasbinEnforcer()

	v1.Use(auth.CasbinMiddleware(ca))
	{
		// Booking Personal routes
		bookingPersonal := v1.Group("/booking-personal")
		{
			bookingPersonal.POST("", handler.BookingPersonalHandler.CreateBookingPersonal)
			bookingPersonal.GET(":id", handler.BookingPersonalHandler.GetBookingPersonal)
			bookingPersonal.PUT(":id", handler.BookingPersonalHandler.UpdateBookingPersonal)
			bookingPersonal.DELETE(":id", handler.BookingPersonalHandler.DeleteBookingPersonal)
			bookingPersonal.GET("", handler.BookingPersonalHandler.ListBookingPersonal)
		}

		// Booking Group routes
		bookingGroup := v1.Group("/booking-group")
		{
			bookingGroup.POST("", handler.BookingGroupHandler.CreateBookingGroup)
			bookingGroup.GET(":id", handler.BookingGroupHandler.GetBookingGroup)
			bookingGroup.PUT(":id", handler.BookingGroupHandler.UpdateBookingGroup)
			bookingGroup.DELETE(":id", handler.BookingGroupHandler.DeleteBookingGroup)
			bookingGroup.GET("", handler.BookingGroupHandler.ListBookingGroup)
		}

		// Booking Coach routes
		bookingCoach := v1.Group("/booking-coach")
		{
			bookingCoach.POST("", handler.BookingCoachHandler.CreateBookingCoach)
			bookingCoach.GET(":id", handler.BookingCoachHandler.GetBookingCoach)
			bookingCoach.PUT(":id", handler.BookingCoachHandler.UpdateBookingCoach)
			bookingCoach.DELETE(":id", handler.BookingCoachHandler.DeleteBookingCoach)
			bookingCoach.GET("", handler.BookingCoachHandler.ListBookingCoach)
		}

		// Subscription Personal routes
		subscriptionPersonal := v1.Group("/subscription-personal")
		{
			subscriptionPersonal.POST("", handler.SubscriptionPersonalHandler.CreateSubscriptionPersonal)
			subscriptionPersonal.GET(":id", handler.SubscriptionPersonalHandler.GetSubscriptionPersonal)
			subscriptionPersonal.PUT(":id", handler.SubscriptionPersonalHandler.UpdateSubscriptionPersonal)
			subscriptionPersonal.DELETE(":id", handler.SubscriptionPersonalHandler.DeleteSubscriptionPersonal)
			subscriptionPersonal.GET("", handler.SubscriptionPersonalHandler.ListSubscriptionPersonal)
		}

		// Subscription Group routes
		subscriptionGroup := v1.Group("/subscription-group")
		{
			subscriptionGroup.POST("", handler.SubscriptionGroupHandler.CreateSubscriptionGroup)
			subscriptionGroup.GET(":id", handler.SubscriptionGroupHandler.GetSubscriptionGroup)
			subscriptionGroup.PUT(":id", handler.SubscriptionGroupHandler.UpdateSubscriptionGroup)
			subscriptionGroup.DELETE(":id", handler.SubscriptionGroupHandler.DeleteSubscriptionGroup)
			subscriptionGroup.GET("", handler.SubscriptionGroupHandler.ListSubscriptionGroup)
		}

		// Subscription Coach routes
		subscriptionCoach := v1.Group("/subscription-coach")
		{
			subscriptionCoach.POST("", handler.SubscriptionCoachHandler.CreateSubscriptionCoach)
			subscriptionCoach.GET(":id", handler.SubscriptionCoachHandler.GetSubscriptionCoach)
			subscriptionCoach.PUT(":id", handler.SubscriptionCoachHandler.UpdateSubscriptionCoach)
			subscriptionCoach.DELETE(":id", handler.SubscriptionCoachHandler.DeleteSubscriptionCoach)
			subscriptionCoach.GET("", handler.SubscriptionCoachHandler.ListSubscriptionCoach)
		}

		// Access routes
		access := v1.Group("/access")
		{
			access.POST("/personal", handler.AccessHandler.CreateAccessPersonal)
			access.GET("/personal/:booking_personal_id", handler.AccessHandler.ListAccessPersonal)
			access.POST("/group", handler.AccessHandler.CreateAccessGroup)
			access.GET("/group/:booking_group_id", handler.AccessHandler.ListAccessGroup)
			access.POST("/coach", handler.AccessHandler.CreateAccessCoach)
			access.GET("/coach/:booking_coach_id", handler.AccessHandler.ListAccessCoach)
		}
		accessBeta := v1.Group("/access-beta")
		{
			accessBeta.GET("/personal/:sport_hall_id", handler.AccessBetaHandler.CheckUserAccessBetaPersonal)
		}
	}

	return router
}
