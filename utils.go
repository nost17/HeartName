package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"charm.land/huh/v2"
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

func ingresarParametros(acento *ColorDisponible) (string, int) {
	var seleccionProporcion int
	var seleccionNombre = ""
	opcionesDeColores := make([]huh.Option[ColorDisponible], len(ColoresDisponibles))
	for i, v := range ColoresDisponibles {
		opcionesDeColores[i] = huh.NewOption(v.nombre, v).Selected(v.nombre == acento.nombre)
	}

	huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title(fmt.Sprintf("Elige la proporcion [%d - %d]", ProporcionMinima, ProporcionMaxima)).
			Validate(func(s string) error {
				salida, err := strconv.Atoi(s)
				if err != nil {
					return fmt.Errorf("Intenta otra vez")
				} else if salida < ProporcionMinima || salida > ProporcionMaxima {
					return fmt.Errorf("El numero debe estar entre %d y %d", ProporcionMinima, ProporcionMaxima)
				}
				seleccionProporcion = salida
				return nil
			}),
		huh.NewInput().
			Title("Ingresa un nombre").
			CharLimit(40).
			Placeholder("Puede estar en blanco").Validate(func(s string) error { return nil }),
		huh.NewSelect[ColorDisponible]().Value(acento).
			Title("Selecciona un color").Options(opcionesDeColores...),
	)).Run()

	return seleccionNombre, seleccionProporcion
}
