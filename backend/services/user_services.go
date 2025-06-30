package services

import (
	"fmt"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/Utils"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/clients"
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
	token, err := Utils.GenerateJWT(userDAO.IDUsuario, userDAO.IsAdmin)
	if err != nil {
		return domain.Usuario{}, "", fmt.Errorf("error generating token: %w", err)
	}
	return usuario, token, nil
}

func GetActInscripto(iduser int) ([]domain.ActDeportiva, []domain.Horario, error) {
	actDAO, horario, err := clients.GetActInscripcion(iduser)

	if err != nil {
		return []domain.ActDeportiva{}, []domain.Horario{}, fmt.Errorf("error getting Act: %w", err)
	}

	acts := make([]domain.ActDeportiva, 0)
	for _, act := range actDAO {
		acts = append(acts, domain.ActDeportiva{
			IDActividad:    act.IDActividad,
			Nombre:         act.Nombre,
			NombreProfesor: act.NombreProfesor,
			Foto:           act.Foto,
			Descripcion:    act.Descripcion})

	}

	hs := make([]domain.Horario, 0)
	for _, horarioDAO := range horario {
		hs = append(hs, domain.Horario{
			IdHorario:     horarioDAO.IdHorario,
			IdActividad:   horarioDAO.IdActividad,
			Dia:           horarioDAO.Dia,
			HorarioInicio: horarioDAO.HorarioInicio,
			HorarioFin:    horarioDAO.HorarioFin,
			Cupos:         horarioDAO.Cupos,
		})
	}

	return acts, hs, nil
}
func InscripcionAct(iduser int, idact int, idhrario int) error {
	err := clients.GenerarInscripcion(iduser, idact, idhrario)
	if err != nil {
		return fmt.Errorf("error generando inscripcion: %w", err)
	}
	return nil
}
func Eliminarinscripcion(idinscrip int) error {
	err := clients.Eliminarinscripcion(idinscrip)
	if err != nil {
		return fmt.Errorf("error eliminando la actividad: %w", err)
	}
	return nil
}

func EliminarInscripcionPorUsuarioYHorario(idUsuario int, idHorario int) error {
	err := clients.EliminarInscripcionPorUsuarioYHorario(idUsuario, idHorario)
	if err != nil {
		return fmt.Errorf("error eliminando la inscripción: %w", err)
	}
	return nil
}
