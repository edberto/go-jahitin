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
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "go-jahitin",
			})
		})

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
			tailor.Use(authMiddleware)
			tailor.GET("/", tailorHandler.GetAll)
		}

		orderHandler := handler.NewOrderHandler(toolkit)
		order := api.Group("/order")
		{
			tailor.Use(authMiddleware)
			order.GET("/", orderHandler.GetAll)
			order.GET("/:id", orderHandler.GetOne)
			order.PUT("/:id", orderHandler.UpdateStatusOne)
			order.POST("/", orderHandler.InsertOne)
		}

		materialHandler := handler.NewMaterialHandler(toolkit)
		material := api.Group("/material")
		{
			material.Use(authMiddleware)
			material.GET("/", materialHandler.GetAll)
		}

		mdlHandler := handler.NewModelHandler(toolkit)
		mdl := api.Group("/model")
		{
			mdl.Use(authMiddleware)
			mdl.GET("/", mdlHandler.GetAll)
		}

	}
}
