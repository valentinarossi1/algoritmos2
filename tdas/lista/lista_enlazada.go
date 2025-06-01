package lista

type nodoLista[T any] struct {
	valor     T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func nodoCrear[T any](valor T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.valor = valor
	nodo.siguiente = nil
	return nodo
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevoNodo := nodoCrear(valor)

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
		lista.primero = nuevoNodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevoNodo := nodoCrear(valor)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) verificarListaVacia() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.verificarListaVacia()
	elemento := lista.primero.valor
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return elemento
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.verificarListaVacia()
	return lista.primero.valor
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.verificarListaVacia()
	return lista.ultimo.valor
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo

}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	elemento := lista.primero
	visitado := true
	for i := 0; i < lista.largo && visitado; i++ {
		visitado = visitar(elemento.valor)
		elemento = elemento.siguiente
	}
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	posicion *nodoLista[T]
	anterior *nodoLista[T]
}

func (iterador *iterListaEnlazada[T]) verificarIteradorFinalizado() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	iterador.verificarIteradorFinalizado()
	return iterador.posicion.valor

}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.posicion != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	iterador.verificarIteradorFinalizado()
	iterador.anterior = iterador.posicion
	iterador.posicion = iterador.posicion.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := nodoCrear(elemento)
	B := iterador.posicion
	A := iterador.anterior

	if A == nil && B == nil { // lista vacia
		iterador.lista.primero = nuevoNodo
		iterador.lista.ultimo = nuevoNodo
		iterador.posicion = nuevoNodo

	} else if A == nil { // inserto en primera posicion
		nuevoNodo.siguiente = B
		iterador.lista.primero = nuevoNodo
		iterador.posicion = nuevoNodo
	} else if B == nil { // inserto el ultimo
		A.siguiente = nuevoNodo
		iterador.posicion = nuevoNodo
		iterador.lista.ultimo = nuevoNodo
	} else {
		A.siguiente = nuevoNodo
		nuevoNodo.siguiente = B
		iterador.posicion = nuevoNodo
	}
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	iterador.verificarIteradorFinalizado()
	borrado := iterador.posicion.valor
	B := iterador.posicion
	A := iterador.anterior
	C := B.siguiente
	if A == nil { // borro en primera posicion
		iterador.posicion = C
		iterador.anterior = nil
		iterador.lista.primero = C
	} else if B.siguiente == nil { // borro el ultimo
		A.siguiente = nil
		iterador.posicion = nil
		iterador.lista.ultimo = A
	} else {
		A.siguiente = C
		iterador.posicion = C
		iterador.anterior = A
	}
	iterador.lista.largo--
	return borrado
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iteradorLista := new(iterListaEnlazada[T])
	iteradorLista.lista = lista
	iteradorLista.posicion = lista.primero
	return iteradorLista
}
