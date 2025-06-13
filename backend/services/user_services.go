package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/Utils"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/domain"
)

func Login(username string, password string) (domain.Usuario, string, error) {
	userDAO, err := clients.GetUserByUsername(username)
	var usuario domain.Usuario
	usuario.IDUsuario = userDAO.IDUsuario
	usuario.IsAdmin = userDAO.IsAdmin

	if err != nil {
		return domain.Usuario{}, "", fmt.Errorf("error getting user: %w", err)
	}
	if Utils.HashSHA256(password) != userDAO.ContraseniaHash {
		return domain.Usuario{}, "", fmt.Errorf("invalid password")
	}
	token, err := Utils.GenerateJWT(userDAO.IDUsuario)
	if err != nil {
		return domain.Usuario{}, "", fmt.Errorf("error generating token: %w", err)
	}
	return usuario, token, nil
}

func GetActInscripto(IDuser int) ([]dao.ActDeportiva, []dao.Horario, error) {
	ActDAO, Horario, err := clients.GetActInscripcion(IDuser)
	if err != nil {
		return []dao.ActDeportiva{}, Horario, fmt.Errorf("error getting Act: %w", err)
	}
	return ActDAO, Horario, nil
}
func InscripcionAct(IDuser int, IDact int, IDhrario int) error {
	err := clients.GenerarInscripcion(IDuser, IDact, IDhrario)
	if err != nil {
		return fmt.Errorf("error generando inscripcion: %w", err)
	}
	return nil
}
