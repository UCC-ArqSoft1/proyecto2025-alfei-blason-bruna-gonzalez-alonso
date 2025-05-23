package domain

type ActDeportiva struct {
	IDActividad    int    `json:"id" gorm:"primary_key"`
	Nombre         string `json:"name"`
	NombreProfesor string `json:"professor"`
	Cupos          int    `json:"cupos"`
	IdCategoria    int    `json:"category"`
}
type Usuario struct {
	IDUsuario     int    `json:"id" gorm:"primary_key"`
	Nombre        string `json:"name"`
	Apellido      string `json:"lastname"`
	NombreUsuario string `json:"username"`
	DNI           int    `json:"dni"`
	Mail          string `json:"mail"`
	Contrasenia   string `json:"password"`
	IsAdmin       bool   `json:"admin: "`
}

type Horario struct {
	IdHorario     int      `json:"id" gorm:"primary_key"`
	Dia           []string `json:"days"`
	HorarioInicio []string `json:"hourstart"`
	HorarioFin    []string `json:"hourfinish"`
	IdActividad   int      `json:"activity"`
}
type Inscripcion struct {
	IdInscripcion int `json:"id" gorm:"primary_key"`
	IdUsuario     int `json:"user"`
	IdActividad   int `json:"activity"`
	IdHorario     int `json:"Horario"`
}

type Categoria struct {
	IdCategoria int    `json:"id" gorm:"primary_key"`
	Nombre      string `json:"name"`
}
