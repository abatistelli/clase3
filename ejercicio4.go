/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos
a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma
dirección de memoria en el main del programa como en las funciones.

Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/
package main

import "fmt"

type Usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func nuevoProducto(nombre string, precio float64) *Producto {

	var p = new(Producto)

	p.nombre = nombre
	p.precio = precio
	p.cantidad = 0

	return p
}

func agragarProducto(u *Usuario, p *Producto, cantidad int) {

	p.cantidad = cantidad
	u.productos = append(u.productos, *p)

}

func borrarProducto(u *Usuario) {
	u.productos = nil
}

func main() {

	//p1 := nuevoProducto("Arroz", 89.50)
	p2 := Producto{"Leche", 110.99, 10}

	u1 := Usuario{
		nombre:   "Usuario1",
		apellido: "Apellido",
		correo:   "mail123@gmail.com",
	}

	fmt.Printf("%+v\n", u1)

	agragarProducto(&u1, &p2, 2)

	fmt.Printf("%+v\n", u1)

	borrarProducto(&u1)

	fmt.Printf("%+v\n", u1)

}
