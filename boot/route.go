package boot

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/handler"
	"go-jahitin/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, toolkit *apipackages.Toolkit) {
	//Middlewares
	_ = middleware.SetupAuthMiddleware(toolkit)

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
			session.DELETE("/logout", sessionHandler.Logout)
		}

		userHandler := handler.NewUserHandler(toolkit)
		user := api.Group("/user")
		{
			user.POST("/register", userHandler.Register)
			user.GET("/:id", userHandler.GetOne)

		}

		tailorHandler := handler.NewTailorHandler(toolkit)
		tailorMaterialHandler := handler.NewTailorMaterialHandler(toolkit)
		tailorModelHandler := handler.NewTailorModelHandler(toolkit)
		tailor := api.Group("/tailor")
		{
			tailor.GET("/", tailorHandler.GetAll)
			tailor.POST("/model", tailorModelHandler.InsertBulk)
			tailor.POST("/material", tailorMaterialHandler.InsertBulk)
		}

		orderHandler := handler.NewOrderHandler(toolkit)
		order := api.Group("/order")
		{
			order.GET("/", orderHandler.GetAll)
			order.GET("/:id", orderHandler.GetOne)
			order.PUT("/:id", orderHandler.UpdateStatusOne)
			order.POST("/", orderHandler.InsertOne)
		}

		materialHandler := handler.NewMaterialHandler(toolkit)
		material := api.Group("/material")
		{
			material.GET("/", materialHandler.GetAll)
		}

		mdlHandler := handler.NewModelHandler(toolkit)
		mdl := api.Group("/model")
		{
			mdl.GET("/", mdlHandler.GetAll)
		}

	}
}
