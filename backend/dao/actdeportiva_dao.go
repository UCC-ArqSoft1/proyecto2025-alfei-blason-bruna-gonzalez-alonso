package dao

type ActDeportiva struct {
	IDActividad    int ` gorm:"primary_key",autoIncrement`
	Nombre         string
	NombreProfesor string
	Cupos          int
	IdCategoria    int
	Horarios       []Horario `gorm:"foreignkey:IdActividad"`
	foto           string
	descripcion    string
}
