package handler

import (
	"github.com/gin-gonic/gin"
	"testbackendGraudate/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		cameras := api.Group("/camera")
		{
			cameras.POST("/", h.createCamera)
			cameras.GET("/", h.listCameras)
			cameras.GET("/:id", h.getCameraById)
			cameras.PUT("/:id", h.updateCamera)
			cameras.DELETE("/:id", h.deleteCamera)
		}

		situations := router.Group("/situations")
		{
			situations.POST("/", h.createSituation)
			situations.GET("/", h.listSituations)
			situations.GET("/:id", h.getSituationById)
			situations.PUT("/:id", h.updateSituation)
			situations.DELETE("/:id", h.deleteSituation)
		}

		incidentTypes := router.Group("/incidentTypes")
		{
			incidentTypes.POST("/", h.createIncidentType)
			incidentTypes.GET("/", h.listIncidentTypes)
			incidentTypes.GET("/:id", h.getIncidentTypeById)
			incidentTypes.PUT("/:id", h.updateIncidentType)
			incidentTypes.DELETE("/:id", h.deleteIncidentType)
		}

		users := router.Group("/users")
		{
			users.GET("/", h.listUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}

		notifications := router.Group("/notifications")
		{
			notifications.POST("/", h.createNotification)
			notifications.GET("/", h.listNotifications)
			notifications.GET("/:id", h.getNotificationById)
			notifications.PUT("/:id", h.updateNotification)
			notifications.DELETE("/:id", h.deleteNotification)
		}

		usersSituations := router.Group("/users_situation")
		{
			usersSituations.POST("/", h.createUserSituation)
			usersSituations.GET("/", h.listUserSituations)
			usersSituations.GET("/:id", h.getUserSituationById)
			usersSituations.PUT("/:id", h.updateUserSituation)
			usersSituations.DELETE("/:id", h.deleteUserSituation)
		}
	}

	return router
}
