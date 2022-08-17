package main

import "fmt"

func CambiarPosiciones(listaNumeral []int, direccion string, posiciones int) []int {
	ListaOrdenada := make([]int, len(listaNumeral))
	for i := 0; i < len(listaNumeral); i++ {
		if direccion == "derecha" {
			if i+posiciones > len(listaNumeral)-1 { // caso de que se pase de la cantidad de items
				ListaOrdenada[i+posiciones-len(listaNumeral)] = listaNumeral[i] // se le resta la cantidad presente del arreglo
			} else {
				ListaOrdenada[i+posiciones] = listaNumeral[i]
			}
		} else if direccion == "izquierda" {
			if i-posiciones < 0 { // caso de que se reste por debajo de 0
				ListaOrdenada[i-posiciones+len(listaNumeral)] = listaNumeral[i] // se procede a sumarle la cantidad correspondiente al arreglo
			} else {
				ListaOrdenada[i-posiciones] = listaNumeral[i]
			}
		}
	}
	return ListaOrdenada
}

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Lista Original:", lista)
	fmt.Println("-----Rotando hacia la izquierda, 2 posiciones-----")
	fmt.Println(CambiarPosiciones(lista, "izquierda", 2))

	lista2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nLista Original:", lista2)
	fmt.Println("-----Rotando hacia la derecha, 3 posiciones-----")
	fmt.Println(CambiarPosiciones(lista2, "derecha", 3))

}

/*
	"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Walter\AppData\Local\Temp\GoLand\___go_build_CambiarPosiciones_go.exe "C:\Users\Walter\Ejercicios\Paradigma Imperativo\CambiarPosiciones.go" #gosetup
	C:\Users\Walter\AppData\Local\Temp\GoLand\___go_build_CambiarPosiciones_go.exe
	Lista Original: [1 2 3 4 5 6 7 8 9 10]
	-----Rotando hacia la izquierda, 2 posiciones-----
	[3 4 5 6 7 8 9 10 1 2]

	Lista Original: [1 2 3 4 5 6 7 8 9 10]
	-----Rotando hacia la derecha, 3 posiciones-----
	[8 9 10 1 2 3 4 5 6 7]

	Process finished with the exit code 0

*/
