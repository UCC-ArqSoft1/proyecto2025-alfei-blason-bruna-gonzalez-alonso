package dao

type ActDeportiva struct {
	IDActividad    int ` gorm:"primary_key",autoIncrement`
	Nombre         string
	NombreProfesor string
	IdCategoria    int
	Horarios       []Horario `gorm:"foreignkey:IdActividad"`
	Foto           string
	Descripcion    string
}
