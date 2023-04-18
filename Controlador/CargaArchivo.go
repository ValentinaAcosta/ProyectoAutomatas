package Controlador

import (
	"ProyectoAutomatasV/Modelo"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type CargaArchivo struct {
}

func (c *CargaArchivo) CargarArc(aut Modelo.Automata) {

	file, err := os.Open("archivo/datos.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	var cadena Modelo.Cadena
	err = json.NewDecoder(file).Decode(&cadena)
	if err != nil {
		fmt.Println("Error al decodificar el archivo:", err)
		return
	}

	//nombres := strings.Split(cadena.Cad, ",")
	//fmt.Println(nombres) // Imprime: [Juan Pedro Maria Laura]

	//fmt.Println("Nombreeeee:", cadena.Cad[0])

	var res []bool
	res = c.EvaluarCadenas(cadena, aut)

	for i := 0; i < len(res); i++ {
		if res[i] == true {
			println("La cadena SI pertenece al lenguaje")
		} else {
			println("La cadena NO pertenece al lenguaje")
		}
	}

}

func (c *CargaArchivo) EvaluarCadenas(cade Modelo.Cadena, aut Modelo.Automata) []bool {
	var cont int
	var resp []bool
	cont = 0
	for i := 0; i < len(cade.Cad); i++ {
		cont = 0
		var c []string
		c = strings.Split(cade.Cad[i], " ")
		for j := 0; j < len(c); j++ {
			println("arista: ", aut.Aristas[j].Cadena, " c: ", c[j])
			if strings.Contains(aut.Aristas[j].Cadena, c[j]) {
				println("Entro al if")
				cont++
			}
		}
		//println("Contador ", cont, " len ", len(c))
		if cont == len(c) {

			//println("Agregue true")
			resp = append(resp, true)
		} else {
			//println("Agregue false")
			resp = append(resp, false)
		}

	}

	return resp

}
