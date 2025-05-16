package dao

type Categoria struct {
	IDCategoria int `gorm:"primary_key"`
	Nombre      string
}
