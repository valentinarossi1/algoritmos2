package main

import (
	"bufio"
	op "dc/operaciones"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

func leerYCalcular(scanner *bufio.Scanner, pila TDAPila.Pila[int64]) {
	for scanner.Scan() {
		linea := scanner.Text()
		tokens := strings.Fields(linea)
		procesarTokens(tokens, pila)
	}
}

func procesarTokens(tokens []string, pila TDAPila.Pila[int64]) {
	exitoCalculo := true
	for i := 0; i < len(tokens) && exitoCalculo; i++ {
		token := tokens[i]
		num, err := strconv.ParseInt(token, 10, 64)
		if err == nil {
			pila.Apilar(num)
		} else {
			exitoCalculo = op.Calcular(token, pila)
		}
	}
	imprimirResultado(exitoCalculo, pila)
}

func imprimirResultado(exitoCalculo bool, pila TDAPila.Pila[int64]) {
	if pila.EstaVacia() || !exitoCalculo {
		fmt.Printf("ERROR\n")
	} else {
		res := pila.Desapilar()
		if !pila.EstaVacia() {
			fmt.Printf("ERROR\n")
		} else {
			fmt.Println(res)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pila := TDAPila.CrearPilaDinamica[int64]()
	leerYCalcular(scanner, pila)
}
