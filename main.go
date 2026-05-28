package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)

const Caracter string = "#"
const AcentoSeleccionado color.Attribute = color.FgGreen

func esperar() {
	time.Sleep((1 * time.Second) / 3)
}

func leerTexto(msg string) string {
	for {
		fmt.Print(msg + ": ")
		if reader.Scan() {
			return reader.Text()
		}
		if err := reader.Err(); err != nil {
			fmt.Println(err)
		}
	}

}

func leerNumero(msg string, error_msg string) int {
	for {
		fmt.Printf("%s: ", msg)
		var string_input string
		if reader.Scan() {
			string_input = reader.Text()
		}
		if err := reader.Err(); err != nil {
			fmt.Println(err)
		}
		if output, err := strconv.Atoi(string_input); err != nil {
			fmt.Println(error_msg)
		} else {
			return output
		}
	}
}

func imprimirNombre(nombre string, linea string, salida *color.Color) {
	n := len(nombre)
	anchoLinea := len(linea)
	anchoNombre := n

	if anchoNombre >= anchoLinea-1 || anchoNombre == 0 {
		fmt.Println(linea)
		return
	}

	anchoLinea = anchoLinea / 2
	anchoNombre = anchoNombre / 2

	anchoNombreIzq := len(nombre[:anchoNombre]) + 1
	anchoNombreDer := len(nombre[anchoNombre:]) + 1
	nombre = fmt.Sprintf(" %s ", nombre)
	nombreColoreado := color.New(color.FgWhite, color.Bold).Sprint(nombre)
	nombreIndice := strings.Index(nombreColoreado, nombre)
	salida.Print(linea[:anchoLinea-anchoNombreIzq])
	for i := range nombreColoreado {
		fmt.Print(string(nombreColoreado[i]))
		if i >= nombreIndice && i <= nombreIndice+n {
			esperar()
		}
	}
	salida.Println(linea[:anchoLinea-anchoNombreDer])
}
func imprimirPicos(linea string, tamañoPico, tamañoMaximo int) {
	for i := tamañoPico; i <= tamañoMaximo; i += 2 {
		relleno := linea[:(i + 1)]
		espacio := strings.Repeat(" ", (tamañoMaximo-i)/2)
		fmt.Print(espacio, relleno, espacio, espacio, relleno, "\n")
		// fmt.Print(len(relleno), len(espacio), i, maximo, tamañoExtra, "\n")
	}
}

func imprimirCuerpo(linea string, tamañoPico, tamañoMaximo int) {
	for i := (tamañoMaximo * 2) - 2; i > tamañoPico+(tamañoPico/2); i -= 2 {
		relleno := linea[:i]
		espacio := strings.Repeat(" ", ((tamañoMaximo*2)-i)/2)
		fmt.Print(espacio, relleno, "\n")
	}
}

func main() {
	tamaño := leerNumero("Ingresa proporcion", "Error, vuelve a intentarlo")
	nombre := leerTexto("Ingresa un nombre")
	nombre = strings.ToUpper(nombre)
	tamaño = tamaño + 1
	tamañoPico := (tamaño / 2) + 2
	/*
		El ancho máximo de un triangulo será igual a
		`(altura * 2) + 1` y tenemos que sumar el ancho
		extra de las puntas (este ancho es para cada una)
	*/
	maximo := (tamaño * 2) + tamañoPico - 1
	linea := strings.Repeat(Caracter, maximo*2)

	salidaColor := color.New(AcentoSeleccionado)

	imprimirPicos(linea, tamañoPico, maximo)
	imprimirNombre(nombre, linea, salidaColor)
	imprimirCuerpo(linea, tamañoPico, maximo)

}
