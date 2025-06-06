package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
)

/*func GetAct(IDact int) (int, string, error, string, int) {
	ActDAO, err := clients.GetActbyId(IDact)
	if err != nil {
		return 0, "", fmt.Errorf("error getting Activity: %w", err), " ", 0
	}
	return ActDAO.IDActividad, ActDAO.Nombre, nil, ActDAO.NombreProfesor, ActDAO.Cupos
}*/
/*
func GetTodasAct() ([]dao.ActDeportiva, error) {
	ActDAO, err := clients.GetActs()
	if err != nil {
		return []dao.ActDeportiva{}, fmt.Errorf("error getting Activity: %w", err)
	}
	return ActDAO, nil
}*/
func GetAct(IDact int) (int, string, error, string, []dao.Horario) {
	ActDAO, err := clients.GetActbyId(IDact)
	if err != nil {
		return 0, "", fmt.Errorf("error getting Activity: %w", err), " ", nil
	}
	horarios, err := clients.GetHorariosByActividad(IDact)
	if err != nil {
		return 0, "", fmt.Errorf("error getting horarios: %w", err), " ", nil
	}
	return ActDAO.IDActividad, ActDAO.Nombre, nil, ActDAO.NombreProfesor, horarios

}
func GetTodasAct() ([]dao.ActConHorarios, error) {
	acts, err := clients.GetActs()
	if err != nil {
		return nil, fmt.Errorf("error getting Activities: %w", err)
	}

	var actividadesConHorarios []dao.ActConHorarios

	for _, act := range acts {
		horarios, err := clients.GetHorariosByActividad(act.IDActividad)
		if err != nil {
			// Si falla obtener horarios de una actividad, podés loguear y seguir o retornar el error
			return nil, fmt.Errorf("error getting horarios for actividad %d: %w", act.IDActividad, err)
		}

		actConHorarios := dao.ActConHorarios{
			Actividad: act,
			Horarios:  horarios,
		}

		actividadesConHorarios = append(actividadesConHorarios, actConHorarios)
	}

	return actividadesConHorarios, nil
}
func CrearActividad(actividad *dao.ActDeportiva) error {
	err := clients.CrearAct(actividad)
	if err != nil {
		return fmt.Errorf("error generando Actividad: %w", err)
	}
	return nil
}

func EliminarActividad(Idact int) error {
	err := clients.EliminarAct(Idact)
	if err != nil {
		return fmt.Errorf("error eliminando la actividad: %w", err)
	}
	return nil
}

func EditarAct(act dao.ActDeportiva) error {
	err := clients.EditarAct(act)
	if err != nil {
		return fmt.Errorf("error generando Actividad: %w", err)
	}
	return nil
}
