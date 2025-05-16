package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/services"
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

	user, token, err := services.Login(req.Nombre_usuario, req.Contrasenia)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// En este punto, deberías generar un JWT (lo vemos en el siguiente paso)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"usuario": user,
		"token":   token, // en el paso 3 lo agregamos
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
