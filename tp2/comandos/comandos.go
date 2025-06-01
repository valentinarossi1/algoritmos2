package comandos

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	"time"
	Aux "tp2/auxiliares"
)

const (
	_POS_IP           = 0
	_POS_FECHA_Y_HORA = 1
	_FORMATO_TIME     = "2006-01-02T15:04:05-07:00"
	_POS_RECURSO      = 3
)

func Agregar_archivo(archivo string, abb TDADiccionario.DiccionarioOrdenado[string, int], hashRecursos TDADiccionario.Diccionario[string, int]) {
	logs, err := os.Open(archivo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando agregar_archivo\n")
		return
	}
	defer logs.Close()
	hashIPs := TDADiccionario.CrearHash[string, TDALista.Lista[time.Time]]()
	abbDoS := TDADiccionario.CrearABB[string, int](Aux.CompararIPs)
	scanner := bufio.NewScanner(logs)
	for scanner.Scan() {
		linea := scanner.Text()
		log := strings.Split(linea, "\t")
		ip := log[_POS_IP]
		fechaYHoraString := log[_POS_FECHA_Y_HORA]
		fechaYHora, _ := time.Parse(_FORMATO_TIME, fechaYHoraString)
		if !hashIPs.Pertenece(ip) {
			listaTiempos := TDALista.CrearListaEnlazada[time.Time]()
			Aux.CargarTiempo(fechaYHora, ip, listaTiempos, hashIPs)
		} else {
			listaTiempos := hashIPs.Obtener(ip)
			Aux.CargarTiempo(fechaYHora, ip, listaTiempos, hashIPs)
		}
		listaAnalizar := hashIPs.Obtener(ip)
		if Aux.VerificarDoS(listaAnalizar) {
			abbDoS.Guardar(ip, 1)
		}
		recurso := log[_POS_RECURSO]
		abb.Guardar(ip, 1)
		Aux.GuardarCantVisitas(hashRecursos, recurso)
	}
	Aux.ImprimirDoS(abbDoS)
	fmt.Println("OK")
}

func Ver_visitantes(desde, hasta string, abb TDADiccionario.DiccionarioOrdenado[string, int]) {
	iterador := abb.IteradorRango(&desde, &hasta)
	fmt.Println("Visitantes:")
	for iterador.HaySiguiente() {
		clave, _ := iterador.VerActual()
		fmt.Printf("\t%s\n", clave)
		iterador.Siguiente()
	}
	fmt.Println("OK")
}

func Ver_mas_visitados(n int, hashRecursos TDADiccionario.Diccionario[string, int]) {
	recursos := Aux.GuardarRecursos(hashRecursos)
	heap := TDAHeap.CrearHeapArr[Aux.ParClaveDato](recursos, Aux.Comparar)
	fmt.Println("Sitios más visitados:")
	Aux.ImprimirMasVisitados(heap, n)
	fmt.Println("OK")
}
