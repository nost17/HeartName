package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"golang.org/x/term"
)

var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)

var estilo = lipgloss.NewStyle().Foreground(lipgloss.Green)
var estiloNombre = lipgloss.NewStyle().Foreground(lipgloss.White).Bold(true)
var AnchoDeTerminal int = obtenerAnchoTerminal()
var EspacioParaCentrar string

const Caracter string = "#"
const CentrarDibujo = true
const EspacioParaDibujo = 4

// const AcentoSeleccionado color.Attribute = color.FgGreen

func esperar() {
	time.Sleep((1 * time.Second) / 25)
}

func obtenerAnchoTerminal() int {
	anchoTerminal, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error al obtener el tamaño de la terminal:", err)
		panic(1)
	}
	return anchoTerminal
}

func generarEspacio(tamañoForma int) string {
	var tamañoEspacio int

	if CentrarDibujo {
		tamañoEspacio = (AnchoDeTerminal - tamañoForma) / 2
	} else {
		tamañoEspacio = EspacioParaDibujo
	}

	return strings.Repeat(" ", tamañoEspacio)
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

func dibujarNombre(nombre string, linea string, salida *strings.Builder) {
	n := len(nombre)
	anchoLinea := len(linea)
	anchoNombre := n

	salida.WriteString(EspacioParaCentrar)
	if anchoNombre >= anchoLinea-1 || anchoNombre == 0 {
		salida.WriteString(linea)
		salida.WriteString("\n")
		return
	}

	anchoLinea = anchoLinea / 2
	anchoNombre = anchoNombre / 2

	anchoNombreIzq := len(nombre[:anchoNombre]) + 1
	anchoNombreDer := len(nombre[anchoNombre:]) + 1
	nombre = fmt.Sprintf(" %s ", nombre)
	salida.WriteString(linea[:anchoLinea-anchoNombreIzq])
	salida.WriteString(estiloNombre.Render(nombre))
	salida.WriteString(estilo.Render(linea[:anchoLinea-anchoNombreDer]))
	salida.WriteString("\n")
}
func dibujarPicos(linea string, tamañoPico int, salida *strings.Builder) {
	tamañoMaximo := len(linea) / 2
	for i := tamañoPico; i <= tamañoMaximo; i += 2 {
		relleno := linea[:(i + 1)]
		espacio := strings.Repeat(" ", (tamañoMaximo-i)/2)
		salida.WriteString(EspacioParaCentrar)
		salida.WriteString(espacio)
		salida.WriteString(relleno)
		salida.WriteString(espacio)
		salida.WriteString(espacio)
		salida.WriteString(relleno)
		salida.WriteString("\n")
		// fmt.Print(espacio, relleno, espacio, espacio, relleno, "\n")
		// fmt.Print(len(relleno), len(espacio), i, maximo, tamañoExtra, "\n")
	}
}

func dibujarCuerpo(linea string, tamañoPico int, salida *strings.Builder) {
	tamañoMaximo := len(linea) / 2
	for i := (tamañoMaximo * 2) - 2; i > tamañoPico+(tamañoPico/2); i -= 2 {
		relleno := linea[:i]
		espacio := strings.Repeat(" ", ((tamañoMaximo*2)-i)/2)
		salida.WriteString(EspacioParaCentrar)
		salida.WriteString(espacio)
		salida.WriteString(relleno)
		salida.WriteString("\n")
		// fmt.Print(len(relleno), "-")
	}
}

func imprimirTexto(texto string) {
	var saltarEspera bool = false
	for i := range texto {
		letra := string(texto[i])
		if letra == "[" {
			saltarEspera = true
		} else if letra == "m" && saltarEspera {
			saltarEspera = false
		} else if !saltarEspera && letra != " " {
			esperar()
		}
		lipgloss.Print(letra)
		os.Stdout.Sync()
	}
	// lipgloss.Println(texto)
}

func ocultarCursor() {
	fmt.Print("\x1b[?25l")
}
func mostrarCursor() {
	fmt.Print("\033[?25h")
}
func limpiarTerminal() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	limpiarTerminal()
	tamaño := leerNumero("Ingresa proporcion", "Error, vuelve a intentarlo")
	nombre := leerTexto("Ingresa un nombre")
	ocultarCursor()
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
	EspacioParaCentrar = generarEspacio(maximo * 2)
	salida := strings.Builder{}
	salida.WriteString("\n")
	dibujarPicos(linea, tamañoPico, &salida)
	dibujarNombre(nombre, linea, &salida)
	dibujarCuerpo(linea, tamañoPico, &salida)
	imprimirTexto(estilo.Render(salida.String()))
	mostrarCursor()
}
