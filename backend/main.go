package main

import (
	"github.com/gin-gonic/gin"
	"proyecto2025-alfei-blason-bruna-gonzalez-alonso/controllers"
)

func main() {
	router := gin.New()
	router.Use(controllers.CORS)
	router.POST("/users/login", controllers.Login)
	router.GET("/act_deportiva/:id", controllers.ObtenerAct)
	router.Run(":8080")
}

/*func mostrar(user domain.Usuario) {
	println(user.ID_usuario, "\n",
		user.Mail,
		user.Contraseña)
	return
}
func main() {
	var usuario1 = domain.Usuario{
		ID_usuario: 1,
		Nombre:     "juan",
		Apellido:   "perez",
		DNI:        46032879,
		Mail:       "agus123@gmail.com",
		Contraseña: "12345678",
		Is_admin:   false,
	}
	var user = controllers.RegisterUser(usuario1)
	mostrar(user)
}*/
