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
	router.DELETE("/users/:id/inscripciones", controllers.DesinscripcionActividad)
	router.POST("/act_deportiva", controllers.CrearAct)
	router.PUT("/act_deportiva/:id", controllers.EditarAct)
	router.DELETE("/act_deportiva/:id", controllers.EliminarAct)
	router.DELETE("/users/inscripciones/:id", controllers.Eliminarinscripcion)
	router.Run(":8080")
}
