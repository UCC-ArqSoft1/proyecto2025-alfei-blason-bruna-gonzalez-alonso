package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/Utils"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
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

func GetActInscripto(IDuser int) ([]dao.ActDeportiva, error) {
	ActDAO, err := clients.GetActInscripcion(IDuser)
	if err != nil {
		return []dao.ActDeportiva{}, fmt.Errorf("error getting Act: %w", err)
	}
	return ActDAO, nil
}
func InscripcionAct(IDuser int, IDact int, IDhrario int) error {
	err := clients.GenerarInscripcion(IDuser, IDact, IDhrario)
	if err != nil {
		return fmt.Errorf("error generando inscripcion: %w", err)
	}
	return nil
}
