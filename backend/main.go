package main

import (
	"github.com/gin-gonic/gin"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/controllers"
)

func main() {
	router := gin.New()
	router.Use(controllers.CORS)
	router.POST("/users/login", controllers.Login)
	router.GET("/act_deportiva/:id", controllers.ObtenerAct)
	router.GET("/act_deportiva", controllers.ObtenerTodasAct)
	router.GET("/users/:id/inscripciones", controllers.GetActInscripcion)
	router.POST("/users/:id/inscripciones", controllers.InscripcionActividad)
	router.Run(":8080")
}
