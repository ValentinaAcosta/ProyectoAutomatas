package main

import (
	"ProyectoAutomatasV/Controlador"
	"ProyectoAutomatasV/Modelo"
	"fmt"
	"strconv"
)

func main() {
	var cargaArchivo Controlador.CargaArchivo
	automata := new(Modelo.Automata)
	var nodosTot []*Modelo.Nodo
	var nodosTotales []Modelo.Nodo
	var aristasTotales []Modelo.Arista
	var cantidadNodos int
	fmt.Println("¿Cuantos nodos deseas ingresar? Mínimo 2, Máximo 4")
	fmt.Scanf("%d\n", &cantidadNodos)
	if cantidadNodos < 2 || cantidadNodos > 4 {
		fmt.Println("Error, esa cantidad de nodos no esta permitida")
		return
	}

	var aristasTot []*Modelo.Arista
	var cantidadAristas int
	fmt.Println("¿Cuantas Aristas desea ingresar?")
	fmt.Scanf("%d\n", &cantidadAristas)

	for i := 0; i < cantidadNodos; i++ {
		nodoAux := new(Modelo.Nodo)
		nodoAux = &Modelo.Nodo{Nombre: strconv.FormatInt(int64(i), 10)}
		nodosTot = append(nodosTot, nodoAux)
	}

	nodosTotales = make([]Modelo.Nodo, len(nodosTot))
	for i, n := range nodosTot {
		nodosTotales[i] = *n
	}

	for i := 0; i < cantidadAristas; i++ {
		aristAux := new(Modelo.Arista)

		fmt.Println("Para la arista número ", i+1, ": ")

		var cadena string
		fmt.Println("Ingrese la cadena")
		fmt.Scanf("%s\n", &cadena)

		var nodoIniNombre string
		fmt.Println("Ingrese el nombre del nodo inicial en números")
		fmt.Scanf("%s\n", &nodoIniNombre)
		nodoIni := new(Modelo.Nodo)
		nodoIni = &Modelo.Nodo{Nombre: nodoIniNombre}

		var nodoFinNombre string
		fmt.Println("Ingrese el nombre del nodo final en números")
		fmt.Scanf("%s\n", &nodoFinNombre)
		nodoFin := new(Modelo.Nodo)
		nodoFin = &Modelo.Nodo{Nombre: nodoFinNombre}

		aristAux = &Modelo.Arista{NodoInicial: *nodoIni, NodoFinal: *nodoFin, Cadena: cadena}
		aristasTot = append(aristasTot, aristAux)
	}

	aristasTotales = make([]Modelo.Arista, len(aristasTot))
	for i, n := range aristasTot {
		aristasTotales[i] = *n
	}

	automata = &Modelo.Automata{Nodos: nodosTotales, Aristas: aristasTotales, NodoIni: nodosTotales[0], NodoFin: nodosTotales[len(nodosTotales)-1]}

	//println("Esto es automata: ", automata.NodoIni.Nombre, ",\n", automata.NodoFin.Nombre, "\n", automata.Aristas[len(automata.Aristas)-1].Cadena, "\n", automata.Nodos[len(automata.Nodos)-1].Nombre)

	cargaArchivo.CargarArc(*automata)

}
