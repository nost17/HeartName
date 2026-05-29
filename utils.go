package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/term"
)

var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)

func esperar() {
	time.Sleep(8 * time.Millisecond)
}

func obtenerAnchoTerminal() int {
	anchoTerminal, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		fmt.Println("Error al obtener el tamaño de la terminal:", err)
		panic(1)
	}
	return anchoTerminal
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

func leerNumeroTambienVacio(msg string, error_msg string, defecto int) int {
	for {
		var string_input string = leerTexto(msg)
		if string_input == "" {
			return defecto
		} else if output, err := strconv.Atoi(string_input); err != nil {
			fmt.Println(error_msg)
		} else {
			return output
		}
	}
}

func imprimirColores() int {
	var colorPorDefecto int = 0
	fmt.Println()
	for i, c := range ColoresDisponibles {
		nombre := c.nombre

		if c.nombre == AcentoSeleccionado.nombre {
			nombre += " [Por defecto]"
			colorPorDefecto = i + 1
		}

		lipgloss.Printf("%d. %s %s    ", i+1, nombre, lipgloss.NewStyle().Background(c.valor).Render("   "))
	}
	fmt.Print("\n\n")
	return colorPorDefecto
}

func elegirColor() ColorDisponible {
	colorPorDefecto := imprimirColores()

	n := len(ColoresDisponibles)
	seleccion := leerNumeroTambienVacio("Elija un color [1 - "+strconv.Itoa(n)+"]", "Intenta otra vez", colorPorDefecto)
	seleccion = min(max(seleccion, 1), n)
	seleccion--

	return ColoresDisponibles[seleccion]
}
