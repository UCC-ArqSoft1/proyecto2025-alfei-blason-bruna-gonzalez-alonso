package domain

type ActDeportiva struct {
	IDActividad    int       `json: "IDActividad"`
	Nombre         string    `json: "Nombre"`
	NombreProfesor string    `json: "NombreProfesor"`
	Horarios       []Horario `json: "Horarios"`
	Foto           string    `json: "Foto"`
	Descripcion    string    `json: "Descripcion"`
}

type Usuario struct {
	IDUsuario       int    `json:"id"`
	NombreUsuario   string `json:"username"`
	ContraseniaHash string `json:"password"`
	Nombre          string `json:"name"`
	Apellido        string `json:"lastname"`
	DNI             int    `json:"dni"`
	Mail            string `json:"mail"`
	IsAdmin         bool   `json:"admin: "`
	Foto            string `json:"fotoURL"`
}

type Horario struct {
	IdHorario     int    `json:"IdHorario"`
	IdActividad   int    `json:"IdActividad"`
	Dia           string `json:"Dia"`
	HorarioInicio string `json:"HorarioInicio"`
	HorarioFin    string `json:"HorarioFin"`
	Cupos         int    `json:"Cupos"`
}

type ActConHorarios struct {
	Actividad ActDeportiva `json:"actividad"`
	Horarios  []Horario    `json:"horarios"`
}

type Inscripcion struct {
	IdInscripcion int `json:"id"`
	IdUsuario     int `json:"user"`
	IdActividad   int `json:"activity"`
	IdHorario     int `json:"Horario"`
}
