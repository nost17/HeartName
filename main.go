package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

type ColorDisponible struct {
	nombre string
	valor  ansi.Color
}

var ColoresDisponibles = [...]ColorDisponible{
	{nombre: "Verde", valor: lipgloss.Green},
	{nombre: "Rojo", valor: lipgloss.Red},
	{nombre: "Azul", valor: lipgloss.Blue},
	{nombre: "Amarillo", valor: lipgloss.Yellow},
	{nombre: "Cian", valor: lipgloss.Cyan},
	{nombre: "Magenta", valor: lipgloss.Magenta},
}

var AcentoSeleccionado ColorDisponible = ColoresDisponibles[1]

const Caracter string = "#"
const CentrarDibujo = true
const EspacioParaDibujo = 4
const ProporcionPorDefecto = 3
const ProporcionMaxima = 10
const ProporcionMinima = 1

var EspacioParaCentrar string
var Estilo lipgloss.Style
var EstiloNombre = lipgloss.NewStyle().Foreground(lipgloss.White).Bold(true)
var AnchoDeTerminal int = obtenerAnchoTerminal()

func generarEspacio(tamañoForma int) string {
	var tamañoEspacio int

	if CentrarDibujo {
		tamañoEspacio = (AnchoDeTerminal - tamañoForma) / 2
	} else {
		tamañoEspacio = EspacioParaDibujo
	}

	return strings.Repeat(" ", tamañoEspacio)
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
	salida.WriteString(EstiloNombre.Render(nombre))
	salida.WriteString(Estilo.Render(linea[:anchoLinea-anchoNombreDer]))
	salida.WriteString("\n")
}
func dibujarPicos(linea string, tamañoPico int, salida *strings.Builder) {
	tamañoMaximo := len(linea) / 2
	for i := tamañoPico; i <= tamañoMaximo; i += 4 {
		relleno := linea[:(i + 1)]
		espacio := strings.Repeat(" ", (tamañoMaximo-i)/2)
		salida.WriteString(EspacioParaCentrar)
		salida.WriteString(espacio)
		salida.WriteString(relleno)
		salida.WriteString(espacio)
		salida.WriteString(espacio)
		salida.WriteString(relleno)
		salida.WriteString("\n")
	}
}

func dibujarCuerpo(linea string, tamañoPico int, salida *strings.Builder) {
	tamañoMaximo := len(linea) / 2
	for i := (tamañoMaximo * 2) - 2; i > (tamañoPico / 2); i -= 4 {
		relleno := linea[:i]
		espacio := strings.Repeat(" ", ((tamañoMaximo*2)-i)/2)
		salida.WriteString(EspacioParaCentrar)
		salida.WriteString(espacio)
		salida.WriteString(relleno)
		salida.WriteString("\n")
	}
}

func imprimirTexto(texto string) {
	var saltarEspera bool = false
	for i := range texto {
		letra := string(texto[i])
		lipgloss.Print(letra)
		if letra == " " {
			continue
		} else if letra == "[" {
			saltarEspera = true
		} else if letra == "m" && saltarEspera {
			saltarEspera = false
		} else if !saltarEspera && letra != " " {
			esperar()
		}
		os.Stdout.Sync()
	}
}

func init() {
	lipgloss.EnableLegacyWindowsANSI(os.Stdout)
}

func main() {
	limpiarTerminal()
	tamaño := elegirProporcion()
	nombre := leerTexto("Ingresa un nombre")
	AcentoSeleccionado = elegirColor()
	Estilo = lipgloss.NewStyle().Foreground(AcentoSeleccionado.valor)

	ocultarCursor()
	nombre = strings.ToUpper(nombre)
	tamaño = tamaño + 1
	tamañoPico := tamaño
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
	limpiarTerminal()
	imprimirTexto(Estilo.Render(salida.String()))
	mostrarCursor()
	esperarTecla("Presione una tecla para continuar...")
}
