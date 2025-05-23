package dao

type Inscripcion struct {
	IdInscripcion int `gorm:"primary_key",autoIncrement`
	IdUsuario     int
	IdActividad   int
	IdHorario     int
}
