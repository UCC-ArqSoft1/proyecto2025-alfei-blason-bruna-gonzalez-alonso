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

	user, err := services.Login(req.Nombre_usuario, req.Contrasenia)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// En este punto, deberías generar un JWT (lo vemos en el siguiente paso)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"usuario": user,
		// "token": token, // en el paso 3 lo agregamos
	})
}

/*var usuarios []domain.Usuario // Simula una "base de datos" temporal

func RegisterUser(u domain.Usuario) domain.Usuario {
	// Hashear contraseña
	hash := sha256.New()
	hash.Write([]byte(u.Contraseña))
	u.Contraseña = hex.EncodeToString(hash.Sum(nil))

	// Asignar ID simulado
	u.ID_usuario = len(usuarios) + 1

	// Guardar en lista
	usuarios = append(usuarios, u)

	return u
}
*/
