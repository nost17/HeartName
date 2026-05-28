package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func leerTexto(msg string) string {
	for {
		fmt.Printf("%s: ", msg)
		text, err := reader.ReadString('\n')
		var new_string string = strings.TrimSpace(string(text))
		if err != nil || len(new_string) == 0 {
			continue
		}
		return new_string
	}
}

func leerNumero(msg string, error_msg string) int {
	for {
		fmt.Printf("%s: ", msg)
		var raw_input string
		_, err := fmt.Scan(&raw_input)
		input, err2 := strconv.Atoi(raw_input)
		if err != nil || err2 != nil {
			fmt.Println(error_msg)
			continue
		}
		return input
	}
}

func main() {
	const Caracter string = "#"
	tamaño := leerNumero("Ingresa proporcion", "Error, vuelve a intentarlo")
	tamaño = tamaño + 1
	tamañoExtra := (tamaño / 2) + 2
	/*
		El ancho máximo de un triangulo será igual a
		`(altura * 2) + 1` y tenemos que sumar el ancho
		extra de las puntas (este ancho es para cada una)
	*/
	maximo := (tamaño * 2) + tamañoExtra - 1
	for i := tamañoExtra; i <= maximo; i += 2 {
		relleno := strings.Repeat(Caracter, i+1)
		espacio := strings.Repeat(" ", (maximo-i)/2)
		fmt.Print(espacio, relleno, espacio, espacio, relleno, "\n")
		// fmt.Print(len(relleno), len(espacio), i, maximo, tamañoExtra, "\n")
	}
	for i := maximo * 2; i > maximo-tamañoExtra; i -= 2 {
		relleno := strings.Repeat(Caracter, i)
		espacio := strings.Repeat(" ", ((maximo*2)-i)/2)
		fmt.Print(espacio, relleno, "\n")
	}
	time.Sleep((1 * time.Second) / 3)
}
