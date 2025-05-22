package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
)

func GetAct(IDact int) (int, string, error, string, int) {
	ActDAO, err := clients.GetActbyId(IDact)
	if err != nil {
		return 0, "", fmt.Errorf("error getting Activity: %w", err), " ", 0
	}
	return ActDAO.IDActividad, ActDAO.Nombre, nil, ActDAO.NombreProfesor, ActDAO.Cupos
}

func GetTodasAct() ([]dao.ActDeportiva, error) {
	ActDAO, err := clients.GetActs()
	if err != nil {
		return []dao.ActDeportiva{}, fmt.Errorf("error getting Activity: %w", err)
	}
	return ActDAO, nil
}
