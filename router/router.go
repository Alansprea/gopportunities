package router

import "github.com/gin-gonic/gin"

func Inicialize() {
	//Initialize Router
	router := gin.Default()

	//Iniltialize Routes
	inicializeRoutes(router)

	//Run server
	router.Run(":8080")

}
