package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
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

func listarColores() int {
	var (
		gris = lipgloss.Color("241")

		cabeceraEstilo  = lipgloss.NewStyle().Foreground(lipgloss.White).Bold(true).Align(lipgloss.Center)
		celdaEstilo     = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.White)
		colorPorDefecto = 0
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(gris)).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return cabeceraEstilo
			} else {
				return celdaEstilo
			}
		}).
		Headers("ID", "NOMBRE", "COLOR")

	for i, c := range ColoresDisponibles {
		nombre := c.nombre

		if c.nombre == AcentoSeleccionado.nombre {
			nombre = "* " + nombre
			colorPorDefecto = i + 1
		}
		t.Row(strconv.Itoa(i+1), nombre, lipgloss.NewStyle().Background(c.valor).Render("   "))
	}

	lipgloss.Println(t)

	return colorPorDefecto
}

func elegirColor() ColorDisponible {
	colorPorDefecto := listarColores()

	n := len(ColoresDisponibles)
	seleccion := leerNumeroTambienVacio("Elige un color [1 - "+strconv.Itoa(n)+"]", "Intenta otra vez", colorPorDefecto)
	seleccion = min(max(seleccion, 1), n)
	seleccion--

	return ColoresDisponibles[seleccion]
}

func elegirProporcion() int {
	seleccion := leerNumeroTambienVacio(fmt.Sprintf("%s [%d - %d]", "Elige una proporcion", ProporcionMinima, ProporcionMaxima), "Intenta otra vez", ProporcionPorDefecto)
	seleccion = min(max(seleccion, ProporcionMinima), ProporcionMaxima)
	return seleccion
}
