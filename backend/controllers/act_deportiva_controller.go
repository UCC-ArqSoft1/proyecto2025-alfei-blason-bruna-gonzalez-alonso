package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/services"
	"strconv"
)

type ActIDActividad struct {
	IDactividad    int    `json:"id"`
	Nombre         string `json:"name"`
	NombreProfesor string `json:"professor"`
	Cupos          int    `json:"cupos"`
}

func ObtenerAct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	IDactividad, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	IDactivity, Nombreact, err, Nombreprofesor, cupos, horarios := services.GetAct(IDactividad)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":         "Obtención exitosa",
		"ActId":           IDactivity,
		"NombreActividad": Nombreact,
		"NombreProfesor":  Nombreprofesor,
		"Cupos":           cupos,
		"Horarios":        horarios,
	})
}
func ObtenerTodasAct(ctx *gin.Context) {

	Actividades, err := services.GetTodasAct()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Actividades)
}
