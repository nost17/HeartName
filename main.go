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
	const Caracter rune = '#'
	tamaño := leerNumero("Ingresa proporcion", "Error, vuelve a intentarlo")
	tamaño = tamaño + 1
	tamañoExtra := (tamaño / 2) + 2
	n := tamaño - 1
	/*
		El ancho máximo de un triangulo será igual a
		`(altura * 2) + 1` y tenemos que sumar el ancho
		extra de las puntas (este ancho es para cada una)
	*/
	maximo := (n * 2) + tamañoExtra + 1
	// var j int = 0
	for i := range tamaño {
		/*
			Para obtener la cantidad de numeros en un intervalo [a,b]
			es `b - a + 1` a -> (n-i) y b -> (n+i)
		*/
		relleno := strings.Repeat(string(Caracter), (n+i)-(n-i)+tamañoExtra+1)
		espacio := strings.Repeat(" ", (maximo-len(relleno))/2)
		fmt.Print(espacio, relleno, espacio, espacio, relleno)
		fmt.Printf("\n")
		time.Sleep((1 * time.Second) / 3)
	}
}
