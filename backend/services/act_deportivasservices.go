package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/domain"
)

func Login(username string, password string) (domain.Usuario, error) {

	// Hasheamos la contraseña que ingresó el usuario
	hash := sha256.Sum256([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])
	// Buscamos al usuario en la base
	result := clients.GetUserByUsername(username)
	if hashedPassword == result.Contraseniahash {

	}
	return user, nil
}
