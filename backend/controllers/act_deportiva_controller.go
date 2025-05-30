package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
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

type Crear struct {
	Nombre         string `json:"nombre"`
	NombreProfesor string `json:"nombreProfesor"`
	Cupos          int    `json:"cupos"`
	IdCategoria    int    `json:"idCategoria"`
	Dia            string `json:"Dia"`
	HorarioInicio  string `json:"horarioInicio"`
	HorarioFin     string `json:"horarioFin"`
	Foto           string `json:"foto"`
	Descripcion    string `json:"descripcion"`
}

func CrearAct(ctx *gin.Context) {
	var Act Crear
	if err := ctx.ShouldBindJSON(&Act); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error: ": "JSON invalido"})
		return
	}
	actividad := &dao.ActDeportiva{
		Nombre:         Act.Nombre,
		NombreProfesor: Act.NombreProfesor,
		Cupos:          Act.Cupos,
		IdCategoria:    Act.IdCategoria,
		Horarios: []dao.Horario{{
			Dia:           Act.Dia,
			HorarioInicio: Act.HorarioInicio,
			HorarioFin:    Act.HorarioFin,
		}},
		Foto:        Act.Foto,
		Descripcion: Act.Descripcion,
	}
	err := services.CrearActividad(actividad)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "La actividad se creo correctamente"})
}
func EliminarAct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	IDactividad, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	error := services.EliminarActividad(IDactividad)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "La actividad se elimino correctamente"})
}

func EditarAct(ctx *gin.Context) {
	var act Crear
	if err := ctx.ShouldBindJSON(&act); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error: ": "JSON invalido"})
		return
	}

	actividad := &dao.ActDeportiva{
		Nombre:         act.Nombre,
		NombreProfesor: act.NombreProfesor,
		Cupos:          act.Cupos,
		IdCategoria:    act.IdCategoria,
		Horarios: []dao.Horario{{
			Dia:           act.Dia,
			HorarioInicio: act.HorarioInicio,
			HorarioFin:    act.HorarioFin,
		}},
		Foto:        act.Foto,
		Descripcion: act.Descripcion,
	}
	err := services.CrearActividad(actividad)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "La actividad se edito correctamente"})
}
