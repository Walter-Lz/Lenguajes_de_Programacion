package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

// -- funciones que se implementa para la funcion Sort --
func (l *listaProductos) Len() int {
	return len(*l)
	//TODO implement me
}

func (l *listaProductos) Less(i, j int) bool {
	return (*l)[i].precio < (*l)[j].precio

}
func (l *listaProductos) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
} // --           --         --

var lProductos listaProductos
var path = "Productos.txt"

func crearArchivo() {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("Archivo creado exitosamente", path)
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var i int
	var validar bool
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre { // aumentar el stock en caso de existencia
			(*l)[i].cantidad += cantidad
			(*l)[i].precio = precio
			validar = true
			break
		}
	}
	if validar != true {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}
func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}
func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	var lProductosAuxiliar listaProductos
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad > cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
		} else {
			var i int
			for i = 0; i < len(*l); i++ {
				if i == prod-1 { // se salta la posicion del producto con 0 existencia
					continue
				} else {
					lProductosAuxiliar = append(lProductosAuxiliar, producto{nombre: (*l)[i].nombre, cantidad: (*l)[i].cantidad, precio: (*l)[i].precio})
				}
			}
			lProductos = lProductosAuxiliar
			// fmt.Println("No se puede vender mayor cantidad de productos que los que hay en existencia")
		}
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}
}
func (l *listaProductos) listarProductosMínimos() listaProductos {
	var lProductosAuxiliarMinima listaProductos
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			lProductosAuxiliarMinima = append(lProductosAuxiliarMinima, producto{nombre: (*l)[i].nombre, cantidad: (*l)[i].cantidad, precio: (*l)[i].precio})
		}
	}
	// debe retornar una nueva lista con productos con existencia mínima
	return lProductosAuxiliarMinima
}
func (l *listaProductos) AumentarInventarioDeminimos(listaMinimos listaProductos) {
	var i int
	producto := 0
	for i = 0; i < len(listaMinimos); i++ {
		producto = l.buscarProducto((listaMinimos)[i].nombre)
		(*l)[producto].cantidad = existenciaMinima
	}
}
func llenarDatos() { // caso de que el archivo no contenga inf
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("Manteca", 5, 1000)
	lProductos.agregarProducto("cacao", 10, 1700)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("azucar", 12, 4500)
	ActualizarArchivo()

}
func llenarDatosArchivos() { // cargar los datos de archivo
	listaArchivo := CargarDatosArchivo()
	var i int
	var lista []string
	for i = 0; i < len(listaArchivo); i++ {
		productos := listaArchivo[i]
		NuevosProductos := strings.Split(productos, ",")
		for a, b := range NuevosProductos {
			lista = append(lista, b)
			a = a
		}
		valor1 := lista[1]
		valor2 := lista[2]
		cantidad, a3 := strconv.Atoi(valor1)
		if a3 != nil {
			// ... handle error
			panic(a3)
		}
		costo, a4 := strconv.Atoi(valor2)
		if a4 != nil {
			// ... handle error
			panic(a4)
		}
		lProductos = append(lProductos, producto{nombre: lista[0], cantidad: cantidad, precio: costo}) // se agrega a la lista
		lista = nil                                                                                    // se limpia la lista para volver a utilizar las mismas posiciones
	}

}
func ActualizarArchivo() { // escribir en el archivo
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()
	var a int
	for a = 0; a < len(lProductos); a++ {
		_, err = file.WriteString(lProductos[a].nombre)
		_, err = file.WriteString(",")
		_, err = file.WriteString(strconv.Itoa(lProductos[a].cantidad))
		_, err = file.WriteString(",")
		_, err = file.WriteString(strconv.Itoa(lProductos[a].precio))
		_, err = file.WriteString("\n")
	}
	err = file.Sync()
	if existeError(err) {
		return
	}
}
func ImprimirArchivo() {
	bytesLeidos, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}
	contenido := string(bytesLeidos)
	fmt.Printf("El contenido del archivo es:\n%s", contenido)
}
func CargarDatosArchivo() []string { // leer archivo
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
func Ordenamiento(listaOrdenar listaProductos) listaProductos {
	sort.Sort(&listaOrdenar)
	return listaOrdenar
}

func main() {
	crearArchivo()
	//llenarDatos()
	llenarDatosArchivos() // info desde archv
	fmt.Println("----------------------Datos Cargados de Archivo--------------")
	fmt.Println(lProductos)
	lProductos.venderProducto("arroz", 4)
	fmt.Println(lProductos)
	lProductos.agregarProducto("azucar", 20, 1500)
	fmt.Println(lProductos)
	lProductos.venderProducto("frijoles", 4)
	fmt.Println(lProductos)
	lProductos.venderProducto("leche", 10)

	fmt.Println("\nSe modifica la lista con los productos disponibles")
	fmt.Println(lProductos)

	fmt.Println("\nSe Agrega una funcion con Existencias minimas")
	fmt.Println(lProductos.listarProductosMínimos())

	fmt.Println("\nSe Agrega una funcion para aumentar las Existencias minimas")
	lista := lProductos.listarProductosMínimos()  // se busca la lista de existencias minimas
	lProductos.AumentarInventarioDeminimos(lista) // se pasa como parametro
	fmt.Println(lProductos)

	fmt.Println("--------------------Archivo---------------------------")
	ActualizarArchivo() // se actualizan los cambios
	fmt.Println("Leer Archivo..............")
	ImprimirArchivo()

	fmt.Println("Lista de Productos ordenada por precio.............")
	fmt.Println(Ordenamiento(lProductos))

}
