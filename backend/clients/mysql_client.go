package clients

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/Utils"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
)

var (
	DB *gorm.DB
)

func init() {
	user := "root"
	password := "17122004Ff"
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
		IDUsuario:       1,
		NombreUsuario:   "mateo123",
		ContraseniaHash: Utils.HashSHA256("mateo"),
		Nombre:          "mateo",
		Apellido:        "Alfei",
		DNI:             43928426,
		Mail:            "mateo123@gmail.com",
		IsAdmin:         false,
	})

	DB.Create(&dao.ActDeportiva{
		IDActividad:    1,
		Nombre:         "Spinning",
		NombreProfesor: "Emiliano",
		Cupos:          10,
		IdCategoria:    1,
		Horarios:       []dao.Horario{dao.Horario{IdHorario: 1}, dao.Horario{IdHorario: 3}},
	})

	DB.Create(&dao.ActDeportiva{
		IDActividad:    2,
		Nombre:         "Yoga",
		NombreProfesor: "Juan",
		Cupos:          10,
		IdCategoria:    3,
		Horarios:       []dao.Horario{dao.Horario{IdHorario: 2}},
	})

	DB.Create(&dao.Horario{
		IdHorario:     1,
		Dia:           "Lunes",
		HorarioInicio: "10:00",
		HorarioFin:    "12:00",
	})

	DB.Create(&dao.Horario{
		IdHorario:     2,
		Dia:           "Martes",
		HorarioInicio: "10:00",
		HorarioFin:    "12:00",
	})
	DB.Create(&dao.Horario{
		IdHorario:     3,
		Dia:           "Martes",
		HorarioInicio: "10:00",
		HorarioFin:    "12:00",
	})

	DB.Create(&dao.Usuario{
		IDUsuario:       2,
		Nombre:          "Martina",
		Apellido:        "Valdo",
		NombreUsuario:   "Martina123456",
		DNI:             46032879,
		Mail:            "mmmm@gmail.com",
		ContraseniaHash: Utils.HashSHA256("12345"),
		IsAdmin:         true,
	})

	DB.Create(&dao.Categoria{
		IDCategoria: 1,
		Nombre:      "Musculacion",
	})

	DB.Create(&dao.Categoria{
		IDCategoria: 2,
		Nombre:      "Funcional",
	})

	DB.Create(&dao.Inscripcion{
		IdInscripcion: 1,
		IdUsuario:     1,
		IdActividad:   1,
		Dia:           "Lunes",
		HorarioInicio: "10:00",
		HorarioFin:    "12:00",
	})

}
func GetUserByUsername(username string) (dao.Usuario, error) {
	var user dao.Usuario
	// SELECT * FROM users WHERE username = ? LIMIT 1
	txn := DB.First(&user, "username = ?", username)
	if txn.Error != nil {
		return dao.Usuario{}, fmt.Errorf("error getting user: %w", txn.Error)
	}
	return user, nil
}
