package dao

type Categoria struct {
	IDCategoria int `gorm:"primary_key",autoIncrement`
	Nombre      string
}
