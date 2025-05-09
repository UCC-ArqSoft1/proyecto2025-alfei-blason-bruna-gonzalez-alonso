package dao

type Usuario struct {
	ID_usuario      int    `gorm:"primary_key"`
	Nombre_usuario  string `gorm:unique`
	Contraseniahash string `gorm:not_null`
	Nombre          string
	Apellido        string
	DNI             int
	Mail            string
	Is_admin        bool
}
