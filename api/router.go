package api

import (
	"github.com/gin-gonic/gin"
	"github.com/victornascimento22/api-1.0/controllers"
)

func Router1() {
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.RegisterHandler)

	r.Run(":8080")

}
