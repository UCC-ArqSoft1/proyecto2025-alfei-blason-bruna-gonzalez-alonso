package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/domain"
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
func GetAct(IDact int) (domain.ActDeportiva, []domain.Horario, error) {
	ActDAO, err := clients.GetActbyId(IDact)
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
	act.IDActividad = ActDAO.IDActividad
	act.Nombre = ActDAO.Nombre
	act.Descripcion = ActDAO.Descripcion
	act.NombreProfesor = ActDAO.NombreProfesor
	act.Foto = ActDAO.Foto
	act.IdCategoria = ActDAO.IdCategoria
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
			IdCategoria:    act.IdCategoria,
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

func EliminarActividad(Idact int) error {
	err := clients.EliminarAct(Idact)
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

// BuscarActividades filtra por nombre, categoría o HORA (HH:mm)
// Si q == "", devuelve todas
/*func FiltrarActividades(q string) ([]domain.ActDeportiva, error) {
	var acts []dao.ActDeportiva

	db := clients.DB.Model(&dao.Actividad{}).
		Preload("Horarios")

	if q != "" {
		like := "%" + q + "%"
		db = db.Joins("LEFT JOIN horarios h ON h.id_actividad = actividads.id").
			Where(`
                actividads.nombre      LIKE ?
			   OR actividads.profesor  LIKE ?
             OR DATE_FORMAT(h.hora_inicio, '%H:%i') LIKE ?`,
				like, like, like).
			Group("actividads.id") // evita duplicados
	}

	if err := db.Find(&acts).Error; err != nil {
		return nil, err
	}

	// Map a domain + horarios
	results := make([]domain.ActividadesDeportivas, 0, len(acts))
	for _, a := range acts {
		// armamos slice de horarios para la respuesta
		hs := make([]domain.Horario, 0, len(a.Horarios))
		for _, h := range a.Horarios {
			hs = append(hs, domain.Horario{
				Id:         h.Id,
				Dia:        h.Dia,
				HoraInicio: h.HoraInicio,
				HoraFin:    h.HoraFin,
			})
		}

		results = append(results, domain.ActividadesDeportivas{
			Id:          a.Id,
			Nombre:      a.Nombre,
			Descripcion: a.Descripcion,
			Categoria:   a.Categoria,
			CupoTotal:   a.CupoTotal,
			Profesor:    a.Profesor,
			Imagen:      a.Imagen,
			Horarios:    hs,
	}

	return results, nil
}*/
