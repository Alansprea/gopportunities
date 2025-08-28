package router

import "github.com/gin-gonic/gin"

func Inicialize() {
	router := gin.Default()
	//Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//estamos rodando nossa API
	router.Run(":8080")

}
