package domain

type ActDeportiva struct {
	IDActividad    int       `json: "id"`
	Nombre         string    `json: "name"`
	NombreProfesor string    `json: "profesor"`
	IdCategoria    int       `json: "category"`
	Horarios       []Horario `json: "horario"`
	Foto           string    `json: "fotoURL"`
	Descripcion    string    `json: "descripcion"`
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
	IdHorario     int    `json:"id"`
	IdActividad   int    `json:"activity"`
	Dia           string `json:"days"`
	HorarioInicio string `json:"hourstart"`
	HorarioFin    string `json:"hourfinish"`
	Cupos         int    `json:"cupos"`
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

type Categoria struct {
	IdCategoria int    `json:"id"`
	Nombre      string `json:"name"`
}
