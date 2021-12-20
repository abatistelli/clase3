/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total
de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar
la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).

Sumar Servicios: recibe un array de servicio y
devuelve el precio total (precio * media hora trabajada,
si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).

Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final
(sumando el total de los 3).

*/

package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados float64
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(c chan float64, productos []Producto) {

	var precioFinal float64

	for _, p := range productos {
		precioFinal += p.precio * float64(p.cantidad)
	}
	c <- precioFinal
}

func sumarServicios(c chan float64, servicios []Servicio) {
	var precioFinal float64

	for _, s := range servicios {
		if s.minutosTrabajados < 30 {
			precioFinal += s.precio * (30)
		} else {
			precioFinal += s.precio * (s.minutosTrabajados * 0.5)
		}
	}

	c <- precioFinal
}

func sumarMantenimiento(c chan float64, mantenimientos []Mantenimiento) {
	var precioFinal float64

	for _, m := range mantenimientos {
		precioFinal += m.precio
	}

	c <- precioFinal
}

func sumaTotal(p []Producto, s []Servicio, m []Mantenimiento) float64 {
	c := make(chan float64)
	var sumaTotal float64

	go sumarProductos(c, p)
	sumaTotal += <-c

	go sumarServicios(c, s)
	sumaTotal += <-c

	go sumarMantenimiento(c, m)
	sumaTotal += <-c

	return sumaTotal
}

func main() {

	productosArray := []Producto{
		{"Fernet", 720, 2},
		{"Coca", 380, 6},
		{"Papas", 280, 3},
		{"Hielo", 150, 1},
		{"Helado", 1200, 1},
	}

	serviciosArray := []Servicio{
		{"Mecanico", 700, 3600},
		{"Peluqueria", 450, 90},
	}

	mantArray := []Mantenimiento{
		{"Mantenimiento1", 500},
		{"Mantenimiento2", 1500},
	}

	fmt.Printf("El precio final es: %.2f\n", sumaTotal(productosArray, serviciosArray, mantArray))

}
