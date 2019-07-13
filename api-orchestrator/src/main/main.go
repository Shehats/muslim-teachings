package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/muslim-teachings/api-orchestrator/src/main/controllers"
	"github.com/muslim-teachings/api-orchestrator/src/main/middleware"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	router.Use(middleware.AuthenicationMiddleware())

	router.NoMethod(func(context *gin.Context) {
		fmt.Println("okay")
	})

	router.NoRoute(func(context *gin.Context) {
		fmt.Println("No methods")
	})

	router.HandleMethodNotAllowed = true

	api := router.Group("/api")
	{
		// Array handling apis
		api.GET("/teachings/:type", controllers.GetTeachings)
		api.POST("/teachings/:type", controllers.AddTeachings)
		api.PUT("/teachings/:type", controllers.UpdateTeachings)
		api.DELETE("/teachings/:type", controllers.DeleteTeachings)

		// Single handling apis
		api.GET("/teachings/:type/:slug", controllers.GetTeaching)
		api.POST("/teachings/:type/:slug", controllers.AddTeaching)
		api.PUT("/teachings/:type/:slug", controllers.UpdateTeaching)
		api.DELETE("/teachings/:type/:slug", controllers.DeleteTeaching)
	}

	{
		// Array handling apis
		api.GET("/categories/:type", controllers.GetCategories)
		api.POST("/categories/:type", controllers.AddCategories)
		api.PUT("/categories/:type", controllers.UpdateCategories)
		api.DELETE("/categories/:type", controllers.DeleteCategories)

		// Single handling apis
		api.GET("/categories/:type/:slug", controllers.GetCategory)
		api.POST("/categories/:type/:slug", controllers.AddCategory)
		api.PUT("/categories/:type/:slug", controllers.UpdateCategory)
		api.DELETE("/categories/:type/:slug", controllers.DeleteCategory)
	}

	{
		api.GET("/quran", controllers.GetQuran)
		api.GET("/quran/:surah", controllers.GetSurah)
		api.GET("/ayat/surah/:surah/from/:from/to/:to", controllers.GetAyat)
	}

	{
		api.GET("/hadiths", controllers.GetHadiths)
		api.GET("/hadiths/:topic", controllers.GetHadithsByTopic)
		api.GET("/hadith/:slug", controllers.GetHadith)
	}

	// Start and run the server
	router.Run(":3000")
}
