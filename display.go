package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func esperarTecla(msg string) {
	fmt.Println()
	fmt.Print(msg + " ")
	fd := int(os.Stdin.Fd())

	estadoOriginal, err := term.MakeRaw(fd)
	if err != nil {
		return
	}

	defer term.Restore(fd, estadoOriginal)

	tecla := make([]byte, 1)
	os.Stdin.Read(tecla)
}

func ocultarCursor() {
	fmt.Print("\x1b[?25l")
}
func mostrarCursor() {
	fmt.Print("\033[?25h")
}
func limpiarTerminal() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[H\033[3J")
}
