package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/Utils"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
)

func Login(username string, password string) (int, string, error) {
	userDAO, err := clients.GetUserByUsername(username)
	if err != nil {
		return 0, "", fmt.Errorf("error getting user: %w", err)
	}
	if Utils.HashSHA256(password) != userDAO.Contraseniahash {
		return 0, "", fmt.Errorf("invalid password")
	}
	token, err := Utils.GenerateJWT(userDAO.ID_usuario)
	if err != nil {
		return 0, "", fmt.Errorf("error generating token: %w", err)
	}
	return userDAO.ID_usuario, token, nil
}
