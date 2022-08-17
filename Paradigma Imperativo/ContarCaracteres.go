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

} /*
	"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Walter\AppData\Local\Temp\GoLand\___go_build_Paradigma_Imperativo__2_.exe "C:\Users\Walter\Ejercicios\Paradigma Imperativo\ContarCaracteres.go" #gosetup
	C:\Users\Walter\AppData\Local\Temp\GoLand\___go_build_Paradigma_Imperativo__2_.exe
	El numero de caracteres es:  15
	Cantidad de palabras:  3
	Cantidad de Lineas:  1
	El numero de caracteres es:  37
	Cantidad de palabras:  7
	Cantidad de Lineas:  1

	Process finished with the exit code 0

*/
