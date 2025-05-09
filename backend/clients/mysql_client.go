package clients

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
)

var (
	DB *gorm.DB
)

func init() {
	user := "root"
	password := "root"
	host := "localhost"
	port := 3306
	database := "backend"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, password, host, port, database)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error connecting to DB: %v", err))
	}

	DB.AutoMigrate(&dao.Usuario{})
	DB.Create(&dao.Usuario{
		ID_usuario:      1,
		Mail:            "emiliano",
		Contraseniahash: "mateo",
	})
}
func GetUserByUsername(username string) dao.Usuario {
	var user dao.Usuario
	// SELECT * FROM users WHERE username = ? LIMIT 1
	DB.First(&user, "username = ?", username)
	return user
}
