/*
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto,
con la información de productos comprados, separados por punto y coma (csv).

Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

package main

import (
	"fmt"
	"log"
	"os"
)

type producto struct {
	Id       string
	Precio   float64
	Cantidad int
}

func encadenarString(values []producto) string {
	cadena := "Id;Precio;Cantidad\n"

	for _, value := range values {
		cadena += fmt.Sprintf("%s;%.2f;%d\n", value.Id, value.Precio, value.Cantidad)
	}

	return cadena
}

func main() {

	productos := []producto{
		{"1", 1000, 10},
		{"2", 800, 8},
		{"3", 20, 2},
		{"4", 450.32, 1}}

	cadena := encadenarString(productos)

	fmt.Println(cadena)

	d1 := []byte(cadena)

	err := os.WriteFile("./productos.csv", d1, 0644)

	if err != nil {
		log.Fatal(err)
	}

}
