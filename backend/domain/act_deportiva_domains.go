package domain

type ActDeportiva struct {
	ID_actividad    int
	Nombre          string
	Nombre_profesor string
	Id_usuario      int
	Cupos           int
	Id_horario      int
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
