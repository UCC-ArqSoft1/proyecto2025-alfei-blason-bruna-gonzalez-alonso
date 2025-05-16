package domain

type ActDeportiva struct {
	IDActividad    int    `json:"id" gorm:"primary_key"`
	Nombre         string `json:"name"`
	NombreProfesor string `json:"professor"`
	Cupos          int    `json:"cupos"`
	IdCategoria    int    `json:"category"`
}

type Usuario struct {
	ID_usuario     int
	Nombre         string
	Apellido       string
	Nombre_usuario string
	DNI            int
	Mail           string
	Contrasenia    string
	Is_admin       bool
}

type Horario struct {
	Id_horario     int
	Dia            []string
	Horario_inicio []string
	Horario_fin    []string
}
type Token struct {
	Tiempo   int
	Id_token int
	Activo   bool
	Token    string
}
