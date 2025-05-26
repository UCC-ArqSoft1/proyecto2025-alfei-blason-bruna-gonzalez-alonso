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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	user, token, err, IsAdmin := services.Login(req.Nombre_usuario, req.Contrasenia)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// En este punto, deberías generar un JWT (lo vemos en el siguiente paso)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"usuario": user,
		"token":   token, // en el paso 3 lo agregamos
		"isAdmin": IsAdmin,
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

func GetActInscripcion(ctx *gin.Context) {
	idParam := ctx.Param("id")
	IDusuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id del usuario invalido"})
		return
	}
	actividades, Horario, err := services.GetActInscripto(IDusuario)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	for i, act := range actividades {
		ctx.JSON(http.StatusOK, gin.H{
			"NombreActividad": act.Nombre,
			"NombreProfesor":  act.NombreProfesor,
			"Cupos":           act.Cupos,
			"Dia":             Horario[i].Dia,
			"Hora Inicio":     Horario[i].HorarioInicio,
			"Hora Fin":        Horario[i].HorarioFin,
		})
	}
}

type InscricionReq struct {
	IdActividad int `json:"id_actividad"`
	IdHorario   int `json:"id_horario"`
}

func InscripcionActividad(ctx *gin.Context) {
	idParam := ctx.Param("id")
	IDusuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id del usuario invalido"})
		return
	}

	var req InscricionReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err = services.InscripcionAct(IDusuario, req.IdActividad, req.IdHorario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Mensaje": "Incripcion exitosa"})
}
