package main

import "fmt"

type Calzado struct {
	marca    string
	precio   int
	talla    int
	cantidad int // si hay existencias del mismo
}
type listaCalzados []Calzado

var lista_calzados_slice listaCalzados

func (lista *listaCalzados) agregarCalzado(marca string, precio int, talla int, cantidad int) {
	var i int
	var validar bool
	if talla < 34 || talla > 44 {
		fmt.Println("El tipo de talla no se puede agregar al inventario ")
	} else {
		for i = 0; i < len(*lista); i++ {
			if (*lista)[i].marca == marca && (*lista)[i].talla == talla { // en caso de existir ya la marca y la talla se aumenta el stock
				(*lista)[i].cantidad += cantidad
				validar = true
				break
			}
		}
		if validar != true { // se agrega
			*lista = append(*lista, Calzado{marca: marca, precio: precio, talla: talla, cantidad: cantidad})
		}
	}
}

func (lista *listaCalzados) buscarCalzado(marca string, talla int) int {
	var result = -1
	var i int
	for i = 0; i < len(*lista); i++ {
		if (*lista)[i].marca == marca && (*lista)[i].talla == talla {
			result = i
		}
	}
	return result
}

func (lista *listaCalzados) venderCalzado(marca string, talla int, cantidad int) { // cantidad a comprar
	var calzado = lista.buscarCalzado(marca, talla)
	var ListaAuxiliar listaCalzados
	if calzado != -1 && cantidad > 0 {
		fmt.Printf("Venta de %d calzados,talla %d, marca %s.\n", cantidad, talla, marca)
		if (*lista)[calzado].cantidad > cantidad { // si existe inventario
			(*lista)[calzado].cantidad = (*lista)[calzado].cantidad - cantidad
		} else { // se procede a vender y a eliminar del stock
			var i int
			for i = 0; i < len(*lista); i++ {
				if i == calzado { // se salta la posicion del calzado a eliminar
					continue
				} else { // se guardan temporal en otra lista
					ListaAuxiliar = append(ListaAuxiliar, Calzado{marca: (*lista)[i].marca, precio: (*lista)[i].precio, talla: (*lista)[i].talla, cantidad: (*lista)[i].cantidad})
				}
			}
			lista_calzados_slice = ListaAuxiliar
		}
	}
}

func main() {

	lista_calzados_slice.agregarCalzado("Nike", 20000, 36, 5)
	lista_calzados_slice.agregarCalzado("Adidas", 40000, 38, 4)
	lista_calzados_slice.agregarCalzado("Nike", 20000, 36, 2)
	lista_calzados_slice.agregarCalzado("Adidas", 35000, 39, 8)
	lista_calzados_slice.agregarCalzado("Jordan", 50000, 37, 3)
	lista_calzados_slice.agregarCalzado("Jordan", 50000, 40, 9)
	lista_calzados_slice.agregarCalzado("Adidas", 30000, 40, 8)
	lista_calzados_slice.agregarCalzado("Nike", 30000, 40, 10)
	lista_calzados_slice.agregarCalzado("Jordan", 55000, 42, 9)
	lista_calzados_slice.agregarCalzado("Adidas", 22000, 35, 6)

	fmt.Println("Lista principal")
	fmt.Println(lista_calzados_slice)
	fmt.Println("-------------------------")

	lista_calzados_slice.venderCalzado("Adidas", 40, 8)

	fmt.Println("Lista de zapatos Actulizada")
	fmt.Println(lista_calzados_slice)

} /*
	"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Walter\AppData\Local\Temp\GoLand\___go_build_Paradigma_Imperativo__1_.exe "C:\Users\Walter\Ejercicios\Paradigma Imperativo\TiendaZapato.go" #gosetup
	C:\Users\Walter\AppData\Local\Temp\GoLand\___go_build_Paradigma_Imperativo__1_.exe
	Lista principal
	[{Nike 20000 36 7} {Adidas 40000 38 4} {Adidas 35000 39 8} {Jordan 50000 37 3} {
	Jordan 50000 40 9} {Adidas 30000 40 8} {Nike 30000 40 10} {Jordan 55000 42 9} {A
	didas 22000 35 6}]
	-------------------------
	Venta de 8 calzados,talla 40, marca Adidas.
	Lista de zapatos Actulizada
	[{Nike 20000 36 7} {Adidas 40000 38 4} {Adidas 35000 39 8} {Jordan 50000 37 3} {
	Jordan 50000 40 9} {Nike 30000 40 10} {Jordan 55000 42 9} {Adidas 22000 35 6}]

	Process finished with the exit code 0
*/
