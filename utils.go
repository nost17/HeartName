package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/term"
)

var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)
var TemaCompleto huh.Theme = huh.ThemeFunc(func(isDark bool) *huh.Styles {
	lightDark := lipgloss.LightDark(isDark)
	var ColorTitulos = lightDark(lipgloss.Color("114"), lipgloss.Color("77"))
	var t *huh.Styles = huh.ThemeCharm(isDark)

	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(lipgloss.White).Bold(true)
	t.Focused.Title = t.Focused.Title.Foreground(ColorTitulos).Bold(true)
	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(lipgloss.White)
	// t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(ColorTitulos).Bold(true)
	t.Blurred.Title = t.Focused.Title

	return t
})

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

func ingresarParametros(acento *ColorDisponible) (string, int, error) {
	var seleccionProporcion int
	var seleccionNombre = ""
	textoPrompt := lipgloss.NewStyle().Foreground(lipgloss.Black).Bold(true).Render("? ")
	opcionesDeColores := make([]huh.Option[ColorDisponible], len(ColoresDisponibles))
	for i, v := range ColoresDisponibles {
		opcionesDeColores[i] = huh.NewOption(v.nombre, v).Selected(v.nombre == acento.nombre)
	}

	err := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title(fmt.Sprintf("Elige la proporcion [%d - %d]", ProporcionMinima, ProporcionMaxima)).
			Prompt(textoPrompt).
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
			Value(&seleccionNombre).
			CharLimit(40).
			Prompt(textoPrompt).
			Placeholder("Puede estar en blanco").Validate(func(s string) error { return nil }),
		huh.NewSelect[ColorDisponible]().Value(acento).
			Title("Selecciona un color").Options(opcionesDeColores...),
	)).WithTheme(TemaCompleto).Run()

	return seleccionNombre, seleccionProporcion, err
}
