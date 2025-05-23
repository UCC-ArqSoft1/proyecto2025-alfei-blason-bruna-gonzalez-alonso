package dao

type Inscripcion struct {
	IdInscripcion int `gorm:"primary_key"`
	IdUsuario     int
	IdActividad   int
	IdHorario     int
}
