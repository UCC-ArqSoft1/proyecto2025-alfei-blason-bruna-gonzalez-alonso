package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/services"
	"strconv"
)

type LoginRequest struct {
	Nombre_usuario string `json:"usuario"`
	Contrasenia    string `json:"contrasenia"`
}

func Login(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "JSON inválido"})
		return
	}

	user, token, err := services.Login(req.Nombre_usuario, req.Contrasenia)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// En este punto, deberías generar un JWT (lo vemos en el siguiente paso)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"usuario": user.IDUsuario,
		"token":   token, // en el paso 3 lo agregamos
		"isAdmin": user.IsAdmin,
	})
}
func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if ctx.Request.Method == "OPTIONS" {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.Next()
}

func Eliminarinscripcion(ctx *gin.Context) {
	idParam := ctx.Param("id")

	idinscripcion, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID "})
		return
	}
	error := services.Eliminarinscripcion(idinscripcion)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "La inscripcion se elimino correctamente"})
}

type DesinscripcionReq struct {
	IdHorario int `json:"id_horario"`
}

func DesinscripcionActividad(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idUsuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id del usuario invalido"})
		return
	}

	var req DesinscripcionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err = services.EliminarInscripcionPorUsuarioYHorario(idUsuario, req.IdHorario)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "Desinscripción exitosa"})
}
func GetActInscripcion(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idusuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id del usuario invalido"})
		return
	}
	actividades, horario, err := services.GetActInscripto(idusuario)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var actividadesInscripto []gin.H

	for i, act := range actividades {
		actividad := gin.H{
			"NombreActividad": act.Nombre,
			"NombreProfesor":  act.NombreProfesor,
			"Cupos":           horario[i].Cupos,
			"Dia":             horario[i].Dia,
			"HoraInicio":      horario[i].HorarioInicio,
			"HoraFin":         horario[i].HorarioFin,
			"Foto":            act.Foto,
			"Descripcion":     act.Descripcion,
			"IDActividad":     act.IDActividad,
		}
		actividadesInscripto = append(actividadesInscripto, actividad)
	}

	ctx.JSON(http.StatusOK, actividadesInscripto)
}

type InscricionReq struct {
	IdActividad int `json:"id_actividad"`
	IdHorario   int `json:"id_horario"`
}

func InscripcionActividad(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idusuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id del usuario invalido"})
		return
	}

	var req InscricionReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err = services.InscripcionAct(idusuario, req.IdActividad, req.IdHorario)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "Incripcion exitosa"})
}
