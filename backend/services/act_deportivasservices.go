package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/domain"
)

func GetAct(IDact int) (domain.ActDeportiva, []domain.Horario, error) {
	actDAO, err := clients.GetActbyId(IDact)
	if err != nil {
		return domain.ActDeportiva{}, []domain.Horario{}, fmt.Errorf("error getting horarios: %w", err)
	}
	horarios, err := clients.GetHorariosByActividad(IDact)
	if err != nil {
		return domain.ActDeportiva{}, []domain.Horario{}, fmt.Errorf("error getting horarios: %w", err)
	}
	hs := make([]domain.Horario, 0)
	for _, horarioDAO := range horarios {
		hs = append(hs, domain.Horario{
			IdHorario:     horarioDAO.IdHorario,
			IdActividad:   horarioDAO.IdActividad,
			Dia:           horarioDAO.Dia,
			HorarioInicio: horarioDAO.HorarioInicio,
			HorarioFin:    horarioDAO.HorarioFin,
			Cupos:         horarioDAO.Cupos,
		})
	}

	var act domain.ActDeportiva
	act.IDActividad = actDAO.IDActividad
	act.Nombre = actDAO.Nombre
	act.Descripcion = actDAO.Descripcion
	act.NombreProfesor = actDAO.NombreProfesor
	act.Foto = actDAO.Foto
	return act, hs, nil

}

/*
	func GetTodasAct() ([]domain.ActConHorarios, error) {
		acts, err := clients.GetActs()
		if err != nil {
			return nil, fmt.Errorf("error getting Activities: %w", err)
		}

		var actividadesConHorarios []domain.ActConHorarios

		for _, act := range acts {
			horarios, err := clients.GetHorariosByActividad(act.IDActividad)
			if err != nil {
				// Si falla obtener horarios de una actividad, podés loguear y seguir o retornar el error
				return nil, fmt.Errorf("error getting horarios for actividad %d: %w", act.IDActividad, err)
			}
			hs := make([]domain.Horario, 0)
			for _, horarioDAO := range horarios {
				hs = append(hs, domain.Horario{
					IdHorario:     horarioDAO.IdHorario,
					IdActividad:   horarioDAO.IdActividad,
					Dia:           horarioDAO.Dia,
					HorarioInicio: horarioDAO.HorarioInicio,
					HorarioFin:    horarioDAO.HorarioFin,
					Cupos:         horarioDAO.Cupos,
				})
			}
			var actss domain.ActDeportiva
			actss.IDActividad = act.IDActividad
			actss.Nombre = act.Nombre
			actss.Descripcion = act.Descripcion
			actss.NombreProfesor = act.NombreProfesor
			actss.Foto = act.Foto
			actss.IdCategoria = act.IdCategoria
			actConHorarios := domain.ActConHorarios{
				Actividad: actss,
				Horarios:  hs,
			}

			actividadesConHorarios = append(actividadesConHorarios, actConHorarios)

		}

		return actividadesConHorarios, nil
	}
*/
func GetTodasAct(filtro string) ([]domain.ActConHorarios, error) {
	acts, err := clients.GetActs(filtro)
	if err != nil {
		return nil, fmt.Errorf("error getting Activities: %w", err)
	}

	var actividadesConHorarios []domain.ActConHorarios

	for _, act := range acts {
		horarios, err := clients.GetHorariosByActividad(act.IDActividad)
		if err != nil {
			return nil, fmt.Errorf("error getting horarios for actividad %d: %w", act.IDActividad, err)
		}
		var hs []domain.Horario
		for _, horarioDAO := range horarios {
			hs = append(hs, domain.Horario{
				IdHorario:     horarioDAO.IdHorario,
				IdActividad:   horarioDAO.IdActividad,
				Dia:           horarioDAO.Dia,
				HorarioInicio: horarioDAO.HorarioInicio,
				HorarioFin:    horarioDAO.HorarioFin,
				Cupos:         horarioDAO.Cupos,
			})
		}
		actss := domain.ActDeportiva{
			IDActividad:    act.IDActividad,
			Nombre:         act.Nombre,
			Descripcion:    act.Descripcion,
			NombreProfesor: act.NombreProfesor,
			Foto:           act.Foto,
		}
		actividadesConHorarios = append(actividadesConHorarios, domain.ActConHorarios{
			Actividad: actss,
			Horarios:  hs,
		})
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

func EliminarActividad(idact int) error {
	err := clients.EliminarAct(idact)
	if err != nil {
		return fmt.Errorf("error eliminando la actividad: %w", err)
	}
	return nil
}

func EditarAct(act *dao.ActDeportiva) error {
	err := clients.EditarAct(act)
	if err != nil {
		return fmt.Errorf("error generando Actividad: %w", err)
	}
	return nil
}
