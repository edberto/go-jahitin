package boot

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/handler"
	"go-jahitin/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, toolkit *apipackages.Toolkit) {
	//Middlewares
	authMiddleware := middleware.SetupAuthMiddleware(toolkit)

	//Routes
	api := r.Group("")
	{
		sessionHandler := handler.NewSessionHandler(toolkit)
		session := api.Group("/session")
		{
			session.POST("/login", sessionHandler.Login)
			session.POST("/refresh", sessionHandler.Refresh)
			session.Use(authMiddleware)
			session.DELETE("/logout", sessionHandler.Logout)
		}

		userHandler := handler.NewUserHandler(toolkit)
		user := api.Group("/user")
		{
			user.POST("/register", userHandler.Register)
			user.Use(authMiddleware)
			user.GET("/:id", userHandler.GetOne)

		}

		tailorHandler := handler.NewTailorHandler(toolkit)
		tailor := api.Group("/tailor")
		{
			tailor.GET("/", tailorHandler.GetAll)
		}
	}
}
