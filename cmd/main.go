package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victornascimento22/api-1.0/controllers"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.RegisterHandler)

	r.Run(":8080")

}
