package clients

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/Utils"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/dao"
	"strconv"
)

var (
	DB *gorm.DB
)

func init() {
	/*user := "root"
	password := "Agus2025uccBD-"
	host := "localhost"
	port := 3306
	database := "backend"

	user := "root"
	password := "17122004Ff"

	host := "localhost"
	port := 3306*/

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	//database := os.Getenv("BD_NAME")
	database := "backend"

	if port == "" {
		port = "3306" // default
	}
	portstr, err := strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("puerto inválido: %v", err))
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, password, host, portstr, database)

	/*dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, password, host, port, database)

	var err error*/
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error conectando a DB: %v", err))
	}

	/*DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error connecting to DB: %v", err))
	}*/

	DB.AutoMigrate(&dao.Usuario{})
	DB.AutoMigrate(&dao.Horario{})
	DB.AutoMigrate(&dao.ActDeportiva{}) //crea tablas en la base de datos
	DB.AutoMigrate(&dao.Inscripcion{})

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
		Descripcion: "Actividad cardiovascular que se realiza en una bicicleta fija al ritmo de la música, guiada por un instructor." +
			" Mejora la resistencia, quema calorías y fortalece piernas y glúteos.",
		Foto: "https://as01.epimg.net/deporteyvida/imagenes/2019/09/03/portada/1567536855_286772_1567537023_noticia_normal.jpg ",
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
		Descripcion: "Disciplina física y mental originaria de la India que combina posturas, ejercicios de respiración" +
			"y meditación para mejorar la flexibilidad, la fuerza, el equilibrio y el bienestar general.",
		Foto: "https://phantom-elmundo.unidadeditorial.es/95aebb12721c45a14b949cca2d81c06d/crop/0x0/2475x1666/resize/1200/f/jpg/assets/multimedia/imagenes/2021/08/27/16300683348682.jpg ",
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

	DB.Create(&dao.ActDeportiva{
		Nombre:         "Musculacion",
		NombreProfesor: "José",
		Descripcion: "Es un tipo de entrenamiento físico que busca desarrollar y fortalecer los músculos mediante ejercicios con pesas y resistencia." +
			" Mejora la fuerza, la salud y la forma del cuerpo.",
		Foto: "https://www.rocfit.com/wp-content/uploads/Equipos-de-musculacion-para-entrenamiento-de-tren-inferior.jpg  ",
		Horarios: []dao.Horario{
			{
				Dia:           "Jueves",
				HorarioInicio: "08:00",
				HorarioFin:    "10:00",
				Cupos:         20,
			},
			{
				Dia:           "Martes",
				HorarioInicio: "19:00",
				HorarioFin:    "21:00",
				Cupos:         10,
			},
			{
				Dia:           "Lunes",
				HorarioInicio: "15:00",
				HorarioFin:    "16:00",
				Cupos:         15,
			},
		},
	})

	DB.Create(&dao.ActDeportiva{
		Nombre:         "Funcional",
		NombreProfesor: "Paula",
		Descripcion: "Es una forma de entrenamiento que mejora la fuerza, el equilibrio y la movilidad con ejercicios que imitan movimientos de la vida diaria." +
			" Ideal para ganar rendimiento y prevenir lesiones.",
		Foto: " https://img.freepik.com/foto-gratis/gente-trabajando-interior-junto-pesas_23-2149175410.jpg?ga=GA1.1.414139249.1716316757&semt=ais_hybrid&w=740",
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
			{
				Dia:           "Lunes",
				HorarioInicio: "07:00",
				HorarioFin:    "09:00",
				Cupos:         10,
			},
		},
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
func GetActbyId(id int) (dao.ActDeportiva, error) {
	var act dao.ActDeportiva

	txn := DB.First(&act, "id_actividad = ?", id)
	if txn.Error != nil {
		return dao.ActDeportiva{}, fmt.Errorf("error getting Activity: %w", txn.Error)
	}
	return act, nil
}

/*
	func GetActs() ([]dao.ActDeportiva, error) {
		var Act []dao.ActDeportiva

		txn := DB.Find(&Act)
		if txn.Error != nil {
			return []dao.ActDeportiva{}, fmt.Errorf("error getting Activity: %w", txn.Error)
		}
		return Act, nil
	}
*/

func GetActs(filtro string) ([]dao.ActDeportiva, error) {
	var acts []dao.ActDeportiva
	query := DB.Model(&dao.ActDeportiva{})

	if filtro != "" {
		like := "%" + filtro + "%"
		query = query.Where(`
				act_deportivas.nombre LIKE ? OR
				act_deportivas.nombre_profesor LIKE ?`,
			like, like).
			Group("act_deportivas.id_actividad")
	}
	if err := query.Find(&acts).Error; err != nil {
		return nil, fmt.Errorf("error getting Activity: %w", err)
	}
	return acts, nil
}

func GetHorariosByActividad(idActividad int) ([]dao.Horario, error) {
	var horarios []dao.Horario
	err := DB.Where("id_actividad = ?", idActividad).Find(&horarios).Error
	return horarios, err
}
func GetActInscripcion(IDusuario int) ([]dao.ActDeportiva, []dao.Horario, error) {
	var Inscripcion []dao.Inscripcion
	var horarios []dao.Horario
	err := DB.Where("Id_usuario = ?", IDusuario).Find(&Inscripcion).Error
	if err != nil {
		return []dao.ActDeportiva{}, horarios, fmt.Errorf("Error: el usuario no se encuentra inscripto %w", err)
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
			horarios = append(horarios, horario)
		}
	}
	return actividades, horarios, nil
}
func GenerarInscripcion(iduser int, idact int, idhorario int) error {
	var actividad dao.ActDeportiva
	var horario dao.Horario

	if err := DB.First(&actividad, idact).Error; err != nil {
		return fmt.Errorf("Error: No se encontró la actividad: %w", err)
	}
	if err := DB.First(&horario, idhorario).Error; err != nil {
		return fmt.Errorf("Error: No se encontró el horario: %w", err)
	}
	if horario.Cupos <= 0 {
		return fmt.Errorf("Error: No hay cupos disponibles para esta actividad")
	}

	err := DB.Transaction(func(tx *gorm.DB) error {
		// Validar si ya existe la inscripción
		var count int64
		if err := tx.Model(&dao.Inscripcion{}).
			Where("id_usuario = ? AND id_actividad = ? AND id_horario = ?", iduser, idact, idhorario).
			Count(&count).Error; err != nil {
			return fmt.Errorf("Error al verificar inscripción existente: %w", err)
		}
		if count > 0 {
			return fmt.Errorf("Error: El usuario ya está inscripto en esta actividad y horario")
		}

		// Actualizar cupos si todavía hay disponibles
		if err := tx.Model(&horario).Where("cupos > 0").Update("cupos", gorm.Expr("cupos - ?", 1)).Error; err != nil {
			return fmt.Errorf("Error al actualizar cupos: %w", err)
		}

		// Crear inscripción
		if err := tx.Create(&dao.Inscripcion{
			IdUsuario:   iduser,
			IdActividad: idact,
			IdHorario:   idhorario,
		}).Error; err != nil {
			return fmt.Errorf("Error: No se pudo realizar la inscripción: %w", err)
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

func EliminarAct(idact int) error {
	txn := DB.Delete(&dao.ActDeportiva{}, idact)
	if txn.Error != nil {
		return fmt.Errorf("error eliminando actividad con ID %d: %w", idact, txn.Error)
	}
	return nil
}
func Eliminarinscripcion(idiscrip int) error {
	var inscripcion dao.Inscripcion
	var horario dao.Horario
	DB.First(&inscripcion, "id_inscripcion = ?", idiscrip)
	txn := DB.First(&horario, "id_horario = ?", inscripcion.IdHorario)
	if err := txn.Model(&horario).Update("cupos", gorm.Expr("cupos + ?", 1)).Error; err != nil {
		return fmt.Errorf("Error al actualizar cupos: %w", err)
	}
	txn2 := DB.Delete(&dao.Inscripcion{}, idiscrip)
	if txn.Error != nil {
		return fmt.Errorf("error eliminando actividad con ID %d: %w", idiscrip, txn2.Error)
	}
	return nil
}

func EliminarInscripcionPorUsuarioYHorario(idUsuario int, idHorario int) error {
	var inscripcion dao.Inscripcion
	if err := DB.Where("id_usuario = ? AND id_horario = ?", idUsuario, idHorario).First(&inscripcion).Error; err != nil {
		return fmt.Errorf("Error: No se encontró la inscripción: %w", err)
	}
	return Eliminarinscripcion(inscripcion.IdInscripcion)
}

func EditarAct(act *dao.ActDeportiva) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&dao.ActDeportiva{}).
			Where("id_actividad = ?", act.IDActividad).
			Updates(map[string]interface{}{
				"nombre":          act.Nombre,
				"nombre_profesor": act.NombreProfesor,
				"foto":            act.Foto,
				"descripcion":     act.Descripcion,
			}).Error; err != nil {
			return fmt.Errorf("error actualizando actividad: %w", err)
		}
		if err := tx.Where("id_actividad = ?", act.IDActividad).
			Delete(&dao.Horario{}).Error; err != nil {
			return fmt.Errorf("error eliminando horarios existentes: %w", err)
		}
		for _, horario := range act.Horarios {
			horario.IdActividad = act.IDActividad
			if err := tx.Create(&horario).Error; err != nil {
				return fmt.Errorf("error creando nuevo horario: %w", err)
			}
		}

		return nil
	})
}
