package cola_prioridad

const _CAPACIDAD_INICIAL = 20
const _CONDICION_ACHICAR = 4
const _FACTOR_REDIMENSION = 2

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(colaConPrioridad[T])
	heap.datos = make([]T, _CAPACIDAD_INICIAL)
	heap.cmp = funcion_cmp

	return heap
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(colaConPrioridad[T])
	datos := make([]T, len(arreglo))
	copy(datos, arreglo)
	if len(datos) == 0 {
		heap.datos = make([]T, _CAPACIDAD_INICIAL) // si el arreglo estaba vacio, lo creamos con una capacidad inicial
	} else {
		heap.datos = heapify(datos, funcion_cmp)
		heap.cant = len(datos)
	}
	heap.cmp = funcion_cmp
	return heap
}

func (heap *colaConPrioridad[T]) Encolar(elemento T) {
	if heap.cant == cap(heap.datos) {
		nuevaCapacidad := cap(heap.datos) * _FACTOR_REDIMENSION
		heap.redimensionarHeap(nuevaCapacidad)
	}
	heap.datos[heap.cant] = elemento
	upHeap(heap.datos, heap.cant, heap.cmp)
	heap.cant++
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	heap.verificarColaVacia()
	if (heap.cant*_CONDICION_ACHICAR) <= cap(heap.datos) && (cap(heap.datos)/_FACTOR_REDIMENSION) > _CAPACIDAD_INICIAL {
		nuevaCapacidad := cap(heap.datos) / _FACTOR_REDIMENSION
		heap.redimensionarHeap(nuevaCapacidad)
	}
	borrado := heap.datos[0]
	heap.datos[heap.cant-1], heap.datos[0] = heap.datos[0], heap.datos[heap.cant-1]
	heap.cant--
	downHeap(heap.datos, 0, heap.cant, heap.cmp)

	return borrado
}

func (heap *colaConPrioridad[T]) VerMax() T {
	heap.verificarColaVacia()
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func HeapSort[T any](datos []T, funcion_cmp func(T, T) int) []T {
	datos = heapify(datos, funcion_cmp)
	for i := len(datos) - 1; i >= 0; i-- {
		datos[i], datos[0] = datos[0], datos[i]
		downHeap(datos, 0, i, funcion_cmp)
	}
	return datos
}

func (heap *colaConPrioridad[T]) redimensionarHeap(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, heap.datos)
	heap.datos = nuevosDatos
}

func (heap *colaConPrioridad[T]) verificarColaVacia() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func calcularPosPadre(posHijo int) int {
	padre := (posHijo - 1) / 2
	return padre
}

func calcularPosHijos(posPadre int) (int, int) {
	posHijoIzq := 2*posPadre + 1
	posHijoDer := 2*posPadre + 2
	return posHijoIzq, posHijoDer
}

func heapify[T any](datos []T, funcion_cmp func(T, T) int) []T {
	if len(datos) == 0 {

		return datos
	}
	for i := len(datos) - 1; i >= 0; i-- {
		downHeap(datos, i, len(datos), funcion_cmp)
	}
	return datos
}

func upHeap[T any](datos []T, posHijo int, funcion_cmp func(T, T) int) {
	if posHijo == 0 {
		return
	}
	posPadre := calcularPosPadre(posHijo)
	res := funcion_cmp(datos[posPadre], datos[posHijo])
	if res < 0 {
		datos[posPadre], datos[posHijo] = datos[posHijo], datos[posPadre]
		upHeap(datos, posPadre, funcion_cmp)
	}
}

func downHeap[T any](datos []T, posPadre, cantidad int, funcion_cmp func(T, T) int) {
	posHijoIzq, posHijoDer := calcularPosHijos(posPadre)
	if posHijoIzq >= cantidad { // verifico si tiene hijo izquierdo
		return
	}
	max := posHijoIzq
	if posHijoDer < cantidad {
		res := funcion_cmp(datos[posHijoIzq], datos[posHijoDer])
		if res < 0 {
			max = posHijoDer
		}
	}
	res1 := funcion_cmp(datos[posPadre], datos[max])
	if res1 < 0 {
		datos[posPadre], datos[max] = datos[max], datos[posPadre]
		downHeap(datos, max, cantidad, funcion_cmp)
	}
}
