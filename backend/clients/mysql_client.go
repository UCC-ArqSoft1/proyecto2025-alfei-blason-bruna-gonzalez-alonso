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
	/*user := "root"
	password := "Agus2025uccBD-"
	host := "localhost"
	port := 3306
	database := "backend"*/
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
	DB.AutoMigrate(&dao.Horario{})
	DB.AutoMigrate(&dao.ActDeportiva{}) //crea tablas en la base de datos
	DB.AutoMigrate(&dao.Inscripcion{})
	DB.AutoMigrate(&dao.Categoria{})

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
		IdHorario:     1,
	})
	DB.Create(&dao.Inscripcion{
		IdInscripcion: 2,
		IdUsuario:     1,
		IdActividad:   2,
		IdHorario:     2,
	})
}

func GetUserByUsername(username string) (dao.Usuario, error) {
	var user dao.Usuario
	// SELECT * FROM users WHERE username = ? LIMIT 1
	txn := DB.First(&user, "nombre_usuario = ?", username)
	if txn.Error != nil {
		return dao.Usuario{}, fmt.Errorf("error getting user: %w", txn.Error)
	}
	return user, nil
}
func GetActbyId(ID int) (dao.ActDeportiva, error) {
	var Act dao.ActDeportiva

	txn := DB.First(&Act, "id_actividad = ?", ID)
	if txn.Error != nil {
		return dao.ActDeportiva{}, fmt.Errorf("error getting Activity: %w", txn.Error)
	}
	return Act, nil
}
func GetActs() ([]dao.ActDeportiva, error) {
	var Act []dao.ActDeportiva

	txn := DB.Find(&Act)
	if txn.Error != nil {
		return []dao.ActDeportiva{}, fmt.Errorf("error getting Activity: %w", txn.Error)
	}
	return Act, nil
}

func GetHorariosByActividad(idActividad int) ([]dao.Horario, error) {
	var horarios []dao.Horario
	err := DB.Where("id_actividad = ?", idActividad).Find(&horarios).Error
	return horarios, err
}

func GetActInscripcion(IDusuario int) ([]dao.ActDeportiva, error) {
	var Inscripcion []dao.Inscripcion
	txn := DB.Where("IDusuario = ?", IDusuario).Find(&Inscripcion).Error
	if txn.Error != nil {
		return []dao.ActDeportiva{}, fmt.Errorf("Error: el usuario no se encuentra inscripto %w", txn.Error)
	}

	var actividades []dao.ActDeportiva //lista que contendra las act a las que esta inscripto el usuario
	for _, insc := range Inscripcion { //recorre la lista de inscripciones obtenida anteriormente
		//insc es cada inscr individual
		var actividad dao.ActDeportiva //cada actividad correspondiente a una inscripcion
		if err := DB.First(&actividad, "actividad = ?", insc.IdActividad).Error; err == nil {
			actividades = append(actividades, actividad)
		}

	}
	return actividades, nil
}
func GenerarInscripcion(IDuser int, IDact int, IDhorario int) error {
	txn := DB.Create(&dao.Inscripcion{
		IdUsuario:   IDuser,
		IdActividad: IDact,
		IdHorario:   IDhorario,
	})
	if txn.Error != nil {
		return fmt.Errorf("Error: No se pudo realizar la inscripcion %w", txn.Error)
	}
	return nil
}
