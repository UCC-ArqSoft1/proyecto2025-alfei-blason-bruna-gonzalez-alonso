package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/Utils"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/services"
	"strconv"
	"strings"
)

type ActIDActividad struct {
	IDactividad    int    `json:"id"`
	Nombre         string `json:"name"`
	NombreProfesor string `json:"professor"`
	Cupos          int    `json:"cupos"`
}

func ObtenerAct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	idactividad, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	actividad, horarios, error := services.GetAct(idactividad)
	if error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":         "Obtención exitosa",
		"ActId":           actividad.IDActividad,
		"NombreActividad": actividad.Nombre,
		"NombreProfesor":  actividad.NombreProfesor,
		"Foto":            actividad.Foto,
		"Descripcion":     actividad.Descripcion,
		"Horarios":        horarios,
	})
}

/*
func ObtenerTodasAct(ctx *gin.Context) {

		Actividades, err := services.GetTodasAct()
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, Actividades)
	}
*/
func ObtenerTodasAct(ctx *gin.Context) {
	filtro := ctx.Query("filtro") // Lee el parámetro ?nombre=...

	actividades, err := services.GetTodasAct(filtro)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, actividades)
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
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := Utils.VerifyToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	if !claims.IsAdmin {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No eres admin"})
		return
	}

	var act Crear
	if err := ctx.ShouldBindJSON(&act); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error: ": "JSON invalido"})
		return
	}
	actividad := &dao.ActDeportiva{
		Nombre:         act.Nombre,
		NombreProfesor: act.NombreProfesor,
		Horarios: []dao.Horario{{
			Dia:           act.Dia,
			HorarioInicio: act.HorarioInicio,
			HorarioFin:    act.HorarioFin,
			Cupos:         act.Cupos,
		}},
		Foto:        act.Foto,
		Descripcion: act.Descripcion,
	}
	err = services.CrearActividad(actividad)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "La actividad se creo correctamente"})
}
func EliminarAct(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := Utils.VerifyToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	if !claims.IsAdmin {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No eres admin"})
		return
	}

	idParam := ctx.Param("id")

	idactividad, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	error := services.EliminarActividad(idactividad)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "La actividad se elimino correctamente"})
}

func EditarAct(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := Utils.VerifyToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	if !claims.IsAdmin {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No eres admin"})
		return
	}

	idParam := ctx.Param("id")
	idactividad, err := strconv.Atoi(idParam)
	var act Crear
	if err := ctx.ShouldBindJSON(&act); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error: ": "JSON invalido"})
		return
	}

	actividad := &dao.ActDeportiva{
		IDActividad:    idactividad,
		Nombre:         act.Nombre,
		NombreProfesor: act.NombreProfesor,
		Horarios: []dao.Horario{{
			Dia:           act.Dia,
			HorarioInicio: act.HorarioInicio,
			HorarioFin:    act.HorarioFin,
			Cupos:         act.Cupos,
		}},
		Foto:        act.Foto,
		Descripcion: act.Descripcion,
	}
	err = services.EditarAct(actividad)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "La actividad se edito correctamente"})
}
