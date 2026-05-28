package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

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
