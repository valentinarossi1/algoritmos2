package auxiliares

import (
	"fmt"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	"time"
)

const (
	_DIF_TIEMPO_DOS         = 2
	_LIMITE_SOLICITUDES_DOS = 5
	_LARGO_IPS              = 4
)

type ParClaveDato struct {
	Clave string
	Dato  int
}

// esta funcion imprime las IPs que fueron detectadas como DoS
func ImprimirDoS(abbDoS TDADiccionario.Diccionario[string, int]) {
	iterador := abbDoS.Iterador()
	for iterador.HaySiguiente() {
		clave, _ := iterador.VerActual()
		fmt.Printf("DoS: %s\n", clave)
		iterador.Siguiente()
	}
}

// esta funcion guarda en un hash la cantidad de visitas que tiene cada recurso
func GuardarCantVisitas(hashRecursos TDADiccionario.Diccionario[string, int], recurso string) {
	if !hashRecursos.Pertenece(recurso) {
		hashRecursos.Guardar(recurso, 1)
	} else {
		cantidad := hashRecursos.Obtener(recurso)
		hashRecursos.Guardar(recurso, cantidad+1)
	}
}

// esta funcion verifica si una IP es detectada como DoS
func VerificarDoS(listaTiempo TDALista.Lista[time.Time]) bool {
	fechaYHoraNuevo := listaTiempo.VerUltimo()
	for !listaTiempo.EstaVacia() && (fechaYHoraNuevo.Sub(listaTiempo.VerPrimero()).Seconds() >= _DIF_TIEMPO_DOS) {
		listaTiempo.BorrarPrimero()
	}
	if listaTiempo.Largo() >= _LIMITE_SOLICITUDES_DOS {
		return true
	} else {
		return false
	}
}

// esta funcion compara dos IPs y devuelve -1 si la primera es menor, 1 si la primera es mayor y 0 si son iguales
func CompararIPs(ip1, ip2 string) int {
	ipUno := strings.Split(ip1, ".")
	ipDos := strings.Split(ip2, ".")

	for i := 0; i < _LARGO_IPS; i++ {
		bloque1, _ := strconv.Atoi(ipUno[i])
		bloque2, _ := strconv.Atoi(ipDos[i])

		if bloque1 < bloque2 {
			return -1
		} else if bloque2 < bloque1 {
			return 1
		}
	}
	return 0
}

// esta funcion compara dos datos y devuelve -1 si el primero es menor, 1 si el primero es mayor y 0 si son iguales
func Comparar(recurso1, recurso2 ParClaveDato) int {
	if recurso1.Dato < recurso2.Dato {
		return -1
	} else if recurso1.Dato > recurso2.Dato {
		return 1
	}
	return 0
}

// esta funcion agrega la fecha y hora de cada solicitud
func CargarTiempo(fechaYHora time.Time, ip string, lista TDALista.Lista[time.Time], hashIPs TDADiccionario.Diccionario[string, TDALista.Lista[time.Time]]) {
	lista.InsertarUltimo(fechaYHora)
	hashIPs.Guardar(ip, lista)
}

// esta funcion guarda un recurso y su cantidad de visitantes en un arreglo
func GuardarRecursos(hashRecursos TDADiccionario.Diccionario[string, int]) []ParClaveDato {
	iter := hashRecursos.Iterador()
	recursos := make([]ParClaveDato, hashRecursos.Cantidad())
	i := 0
	for iter.HaySiguiente() {
		recursos[i].Clave, recursos[i].Dato = iter.VerActual()
		i++
		iter.Siguiente()
	}
	return recursos
}

// esta funcion imprime los "n" recursos mas visitados
func ImprimirMasVisitados(heap TDAHeap.ColaPrioridad[ParClaveDato], n int) {
	var auxiliar []ParClaveDato
	min := min(heap.Cantidad(), n)
	for j := 0; j < min; j++ {
		desencolado := heap.Desencolar()
		auxiliar = append(auxiliar, desencolado)
		fmt.Printf("\t%s - %d\n", desencolado.Clave, desencolado.Dato)
	}
	for k := 0; k < len(auxiliar); k++ {
		heap.Encolar(auxiliar[k])
	}
}
