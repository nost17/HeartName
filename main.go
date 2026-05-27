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
	tamaño = int(float32(tamaño) * 1.5)
	n := tamaño - 1
	maximo := (n * 2) + 3 + (tamaño / 2)
	// var j int = 0
	for i := range tamaño {
		relleno := strings.Repeat(string(Caracter), (n+i)-(n-i)+(tamaño/2)+3)
		espacio := strings.Repeat(" ", (maximo-len(relleno))/2)
		fmt.Print(espacio, relleno, espacio, espacio, relleno)
		// for j := range (tamaño * 2) - 1 {
		// relleno := strings.Repeat(string(Caracter), (n+i)-(n-i)+(tamaño/2)+3)
		// if j >= (n - i) {
		// 	if j == n {
		// 		for range (tamaño / 2) + 2 {
		// 			fmt.Printf("%c", Caracter)
		// 		}
		// 	} else {
		// 		fmt.Printf("%s", relleno)
		// 		continue
		// 	}
		// } else {
		// 	fmt.Printf("%c", ' ')
		// }
		fmt.Printf("\n")
		time.Sleep((1 * time.Second) / 3)
		// }
	}
}
