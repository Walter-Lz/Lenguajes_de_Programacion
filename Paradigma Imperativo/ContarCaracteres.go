package main

import "fmt"

func ContarCaracteres(Cadena string) (string, int, string, int, string, int) {
	Palabras, Caracteres, lineas := 0, 0, 0
	for b := range Cadena {
		if string(Cadena[b]) == " " {
			Palabras += 1
		} else if string(Cadena[b]) == "." {
			Palabras += 1 // se contabiliza la ultima palabra
			lineas += 1
		}
	}
	Caracteres = len(Cadena) // se cuenta la cantidad de caracteres
	return "El numero de caracteres es: ", Caracteres, " \nCantidad de palabras: ", Palabras, "\nCantidad de Lineas: ", lineas

}

func main() {
	fmt.Println(ContarCaracteres("Hoy es viernes."))

	fmt.Println(ContarCaracteres("Hoy juega la seleccion de Costa Rica."))

}
