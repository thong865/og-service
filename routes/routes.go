package routes

import (
	cf "MS/config"
	apiControllerV1 "MS/controller/v1"
	"MS/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// r.Static("/storage", "storage")
	// r.Static("/templates", "templates")
	// r.LoadHTMLGlob("templates/*")

	/*
	* Core Configuration
	 */
	r.Use(cf.NewCore)

	v1 := r.Group("/api/v1")
	v1.Use(middlewares.UserMiddlewares())
	{
		v1.POST("user-list", apiControllerV1.UserList)
		v1.GET("/test", apiControllerV1.GetList)
		v1.POST("/login", apiControllerV1.SVLOTLOGIN)
		v1.POST("/register", apiControllerV1.Register)

	}

	return r
}
