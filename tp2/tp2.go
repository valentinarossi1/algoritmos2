package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	Aux "tp2/auxiliares"
	Comandos "tp2/comandos"
)

const (
	_LARGO_AGREGAR           = 2
	_LARGO_VER_VISITANTES    = 3
	_LARGO_VER_MAS_VISITADOS = 2
	_POS_COMANDO             = 0
	_POS_DESDE               = 1
	_POS_HASTA               = 2
	_POS_ARCHIVO             = 1
	_POS_N_MAS_VISITADOS     = 1
)

func leerEntrada(scanner *bufio.Scanner) {
	abb := TDADiccionario.CrearABB[string, int](Aux.CompararIPs)
	hashRecursos := TDADiccionario.CrearHash[string, int]()
	for scanner.Scan() {
		linea := scanner.Text()
		identificarComando(linea, abb, hashRecursos)
	}
}

// esta funcion identifica y ejecuta el comando
func identificarComando(linea string, abb TDADiccionario.DiccionarioOrdenado[string, int], hashRecursos TDADiccionario.Diccionario[string, int]) {
	entrada := strings.Fields(linea)
	comando := entrada[_POS_COMANDO]
	switch comando {
	case "agregar_archivo":
		if len(entrada) == _LARGO_AGREGAR {
			archivo := entrada[_POS_ARCHIVO]
			rutaArchivo := archivo
			Comandos.Agregar_archivo(rutaArchivo, abb, hashRecursos)
		} else {
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", linea)
		}
	case "ver_visitantes":
		if len(entrada) == _LARGO_VER_VISITANTES {
			desde := entrada[_POS_DESDE]
			hasta := entrada[_POS_HASTA]
			Comandos.Ver_visitantes(desde, hasta, abb)
		} else {
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", linea)
		}
	case "ver_mas_visitados":
		if len(entrada) == _LARGO_VER_MAS_VISITADOS {
			n, _ := strconv.Atoi(entrada[_POS_N_MAS_VISITADOS])
			Comandos.Ver_mas_visitados(n, hashRecursos)
		} else {
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", linea)
		}
	default:
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", linea)
		return
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	leerEntrada(scanner)
}
