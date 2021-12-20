/*
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que
vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la
estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:

La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.
*/

package main

import "fmt"

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func cambiarNombre(u *usuario, nombre, apellido string) {
	u.nombre = nombre
	u.apellido = apellido
}

func cambiarEdad(u *usuario, edad int) {
	u.edad = edad
}

func cambiarCorreo(u *usuario, correo string) {
	u.correo = correo
}

func cambiarContrseña(u *usuario, contraseña string) {
	u.contraseña = contraseña
}

func main() {
	u1 := usuario{
		nombre:     "TestNombre",
		apellido:   "Testapellido",
		edad:       23,
		correo:     "Testcorreo@gmail.com",
		contraseña: "123",
	}

	fmt.Printf("%+v\n", u1)

	cambiarNombre(&u1, "Cambie", "NyA")
	cambiarEdad(&u1, 40)
	cambiarCorreo(&u1, "nuevomail@gmail.com")
	cambiarContrseña(&u1, "321")

	fmt.Printf("%+v\n", u1)
}
