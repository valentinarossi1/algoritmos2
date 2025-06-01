package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T] // el siguiente a desencolar
	ultimo  *nodoCola[T] // el que recien se encoló
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) verificarColaVacia() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	cola.verificarColaVacia()
	return cola.primero.dato
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	nodo.prox = nil
	return nodo
}

func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nuevoNodo := nodoCrear(elemento)

	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	cola.verificarColaVacia()
	elemento := cola.primero.dato    //guardo el elemento a desencolar en una variable
	cola.primero = cola.primero.prox //el primer nodo ahora pasa a ser el siguiente del primero

	if cola.primero == nil {
		cola.ultimo = nil //si al desencolar la cola queda vacia, se actualiza el ultimo a nil
	}
	return elemento
}
