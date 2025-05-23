package dao

type Usuario struct {
	IDUsuario       int    ` gorm:"primary_key",autoIncrement`
	NombreUsuario   string `gorm:"unique"`
	ContraseniaHash string `gorm:"not_null"`
	Nombre          string
	Apellido        string
	DNI             int
	Mail            string `gorm:"unique"`
	IsAdmin         bool
}
