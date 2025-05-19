package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/Utils"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
)

func Login(username string, password string) (int, string, error, bool) {
	userDAO, err := clients.GetUserByUsername(username)
	if err != nil {
		return 0, "", fmt.Errorf("error getting user: %w", err), false
	}
	if Utils.HashSHA256(password) != userDAO.ContraseniaHash {
		return 0, "", fmt.Errorf("invalid password"), false
	}
	token, err := Utils.GenerateJWT(userDAO.IDUsuario)
	if err != nil {
		return 0, "", fmt.Errorf("error generating token: %w", err), false
	}
	return userDAO.IDUsuario, token, nil, userDAO.IsAdmin
}
