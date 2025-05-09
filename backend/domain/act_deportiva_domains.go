package domain

type ActDeportiva struct {
	ID_actividad int
	Nombre       string
	Id_profesor  int
	Id_usuario   int
	Cupos        int
	Id_horario   int
}

type Usuario struct {
	ID_usuario int
	ID_token   int
	Nombre     string
	Apellido   string
	DNI        int
	Mail       string
	Contraseña string
	Is_admin   bool
}

type Horario struct {
	Id_horario     int
	Id_actividad   int
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

type Profesor struct {
	ID_profesor int
	ID_usuario  int
	Nombre      string
	Apellido    string
}
