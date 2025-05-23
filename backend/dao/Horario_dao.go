package dao

type Horario struct {
	IdHorario     int `gorm:"primary_key",autoIncrement`
	IdActividad   int
	Dia           string
	HorarioInicio string
	HorarioFin    string
}

type ActConHorarios struct {
	Actividad ActDeportiva `json:"actividad"`
	Horarios  []Horario    `json:"horarios"`
}
