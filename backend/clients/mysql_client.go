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
		NombreUsuario:   "mateo123",
		ContraseniaHash: Utils.HashSHA256("mateo"),
		Nombre:          "mateo",
		Apellido:        "Alfei",
		DNI:             43928426,
		Mail:            "mateo123@gmail.com",
		IsAdmin:         false,
	})

	DB.Create(&dao.ActDeportiva{
		Nombre:         "Spinning",
		NombreProfesor: "Emiliano",
		IdCategoria:    1,
		Horarios: []dao.Horario{
			{
				Dia:           "Martes",
				HorarioInicio: "18:00",
				HorarioFin:    "20:00",
				Cupos:         10,
			},
			{
				Dia:           "Viernes",
				HorarioInicio: "14:00",
				HorarioFin:    "15:00",
				Cupos:         10,
			},
		},
	})

	DB.Create(&dao.ActDeportiva{
		Nombre:         "Yoga",
		NombreProfesor: "Juan",
		IdCategoria:    3,
		Horarios: []dao.Horario{
			{
				Dia:           "Lunes",
				HorarioInicio: "10:00",
				HorarioFin:    "12:00",
				Cupos:         10,
			},
			{
				Dia:           "Martes",
				HorarioInicio: "10:00",
				HorarioFin:    "12:00",
				Cupos:         10,
			},
		},
	})

	DB.Create(&dao.Usuario{
		Nombre:          "Martina",
		Apellido:        "Valdo",
		NombreUsuario:   "Martina123456",
		DNI:             46032879,
		Mail:            "mmmm@gmail.com",
		ContraseniaHash: Utils.HashSHA256("12345"),
		IsAdmin:         true,
	})

	DB.Create(&dao.Categoria{
		Nombre: "Musculacion",
	})

	DB.Create(&dao.Categoria{
		Nombre: "Funcional",
	})

	DB.Create(&dao.Inscripcion{
		IdUsuario:   1,
		IdActividad: 1,
		IdHorario:   1,
	})
	DB.Create(&dao.Inscripcion{
		IdUsuario:   1,
		IdActividad: 2,
		IdHorario:   2,
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
func GetActInscripcion(IDusuario int) ([]dao.ActDeportiva, []dao.Horario, error) {
	var Inscripcion []dao.Inscripcion
	var Horarios []dao.Horario
	err := DB.Where("Id_usuario = ?", IDusuario).Find(&Inscripcion).Error
	if err != nil {
		return []dao.ActDeportiva{}, Horarios, fmt.Errorf("Error: el usuario no se encuentra inscripto %w", err)
	}
	var actividades []dao.ActDeportiva //lista que contendra las act a las que esta inscripto el usuario
	for _, insc := range Inscripcion { //recorre la lista de inscripciones obtenida anteriormente
		//insc es cada inscr individual
		var actividad dao.ActDeportiva //cada actividad correspondiente a una inscripcion
		if err := DB.First(&actividad, "id_actividad = ?", insc.IdActividad).Error; err == nil {
			actividades = append(actividades, actividad)
		}
		var horario dao.Horario
		if err := DB.First(&horario, "id_horario = ?", insc.IdHorario).Error; err == nil {
			Horarios = append(Horarios, horario)
		}
	}
	return actividades, Horarios, nil
}
func GenerarInscripcion(IDuser int, IDact int, IDhorario int) error {
	var actividad dao.ActDeportiva
	var horario dao.Horario
	if err := DB.First(&actividad, IDact).Error; err != nil {
		return fmt.Errorf("Error: No se encontró la actividad: %w", err)
	}
	if err := DB.First(&horario, IDact).Error; err != nil {
		return fmt.Errorf("Error: No se encontró el horario: %w", err)
	}
	if horario.Cupos <= 0 {
		return fmt.Errorf("Error: No hay cupos disponibles para esta actividad")
	}
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&horario).Where("cupos > 0").Update("cupos", gorm.Expr("cupos - ?", 1)).Error; err != nil {
			return fmt.Errorf("Error al actualizar cupos: %w", err)
		}
		txn := DB.Create(&dao.Inscripcion{
			IdUsuario:   IDuser,
			IdActividad: IDact,
			IdHorario:   IDhorario,
		})
		if txn.Error != nil {
			return fmt.Errorf("Error: No se pudo realizar la inscripcion %w", txn.Error)
		}
		return nil
	})
	return err
}

func CrearAct(actividad *dao.ActDeportiva) error {
	txn := DB.Create(actividad)
	if txn.Error != nil {
		return fmt.Errorf("Error: No se pudo crear la actividad %w", txn.Error)
	}
	return nil
}

func EliminarAct(IDact int) error {
	txn := DB.Delete(&dao.ActDeportiva{}, IDact)
	if txn.Error != nil {
		return fmt.Errorf("error eliminando actividad con ID %d: %w", IDact, txn.Error)
	}
	return nil
}

func EditarAct(act dao.ActDeportiva) error {
	tnx := DB.Save(act)
	if tnx.Error != nil {
		return fmt.Errorf("error editando actividad %w", act, tnx.Error)
	}
	return nil
}
