package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterDiccionarioOrdenado[K comparable, V any] struct {
	desde *K
	hasta *K
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	abb   *abb[K, V]
}

type funcCmp[K comparable] func(K, K) int

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cantidad = 0
	abb.cmp = funcion_cmp
	return abb
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.izquierdo = nil
	nodo.derecho = nil
	nodo.clave = clave
	nodo.dato = dato
	return nodo
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodoActual, padre := buscarNodoYPadre(abb, clave, abb.raiz, nil)
	if nodoActual != nil {
		nodoActual.dato = dato
		return
	}
	nodoActual = crearNodo(clave, dato)
	abb.cantidad++
	if padre == nil {
		abb.raiz = nodoActual
	} else {
		if abb.cmp(clave, padre.clave) < 0 {
			padre.izquierdo = nodoActual
		} else {
			padre.derecho = nodoActual
		}
	}
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo, padre := buscarNodoYPadre(abb, clave, abb.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	borrado := nodo.dato
	if nodo.izquierdo != nil && nodo.derecho != nil {
		return borrarDosHijos(abb, nodo)
	}
	if nodo.izquierdo != nil || nodo.derecho != nil {
		return borrarUnHijo(abb, nodo, padre)
	}
	if padre == nil {
		abb.raiz = nil
	} else {
		res := abb.cmp(nodo.clave, padre.clave)
		if res < 0 {
			padre.izquierdo = nil
		} else {
			padre.derecho = nil
		}
	}
	abb.cantidad--
	return borrado
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := buscarNodoYPadre(abb, clave, abb.raiz, nil)
	return nodo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _ := buscarNodoYPadre(abb, clave, abb.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (iterador *iterDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	return !iterador.pila.EstaVacia()
}

func (iterador *iterDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() || iterador.pila.EstaVacia() {
		panic("El iterador termino de iterar")
	}
	actual := iterador.pila.VerTope()
	return actual.clave, actual.dato
}

func (iterador *iterDiccionarioOrdenado[K, V]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	desapilado := iterador.pila.Desapilar()
	if desapilado.derecho != nil {
		apilarIzquierdos(iterador, desapilado.derecho)
	}
}

func (abb *abb[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	iteradorDiccionario := new(iterDiccionarioOrdenado[K, V])
	iteradorDiccionario.abb = abb
	iteradorDiccionario.desde = desde
	iteradorDiccionario.hasta = hasta
	iteradorDiccionario.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarIzquierdos(iteradorDiccionario, abb.raiz)
	return iteradorDiccionario
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(abb, abb.raiz, desde, hasta, visitar)
}

func iterarRango[K comparable, V any](abb *abb[K, V], nodoActual *nodoAbb[K, V], desde, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodoActual == nil {
		return true
	}
	if desde == nil || abb.cmp(*desde, nodoActual.clave) < 0 { // si la clave del nodo es mayor que "desde", va hacia la izquierda
		if !iterarRango(abb, nodoActual.izquierdo, desde, hasta, visitar) {
			return false
		}
	}
	if (desde == nil || abb.cmp(nodoActual.clave, *desde) >= 0) && (hasta == nil || abb.cmp(nodoActual.clave, *hasta) <= 0) { // si el nodo esta dentro de los parametros recibidos, aplica la funcion visitar
		if !visitar(nodoActual.clave, nodoActual.dato) {
			return false
		}
	}
	if hasta == nil || abb.cmp(*hasta, nodoActual.clave) > 0 { // si la clave del nodo es menor que "hasta", va hacia la derecha
		if !iterarRango(abb, nodoActual.derecho, desde, hasta, visitar) {
			return false
		}
	}
	return true
}

// esta funcion recibe una clave y devuelve el nodo en el que se encuentra y su nodo padre
func buscarNodoYPadre[K comparable, V any](abb *abb[K, V], clave K, nodoActual, padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodoActual == nil {
		return nodoActual, padre
	}
	res := abb.cmp(clave, nodoActual.clave)
	if res < 0 {
		return buscarNodoYPadre(abb, clave, nodoActual.izquierdo, nodoActual)
	} else if res > 0 {
		return buscarNodoYPadre(abb, clave, nodoActual.derecho, nodoActual)
	} else {
		return nodoActual, padre
	}
}

// esta funcion borra un nodo que tiene un solo hijo, haciendo que el padre adopte a dicho hijo
func borrarUnHijo[K comparable, V any](abb *abb[K, V], nodo, padre *nodoAbb[K, V]) V {
	borrado := nodo.dato
	hijo := nodo.izquierdo
	if nodo.derecho != nil { // como sé que tiene un solo hijo, si el derecho no es nulo, este es el hijo
		hijo = nodo.derecho
	}
	if padre == nil {
		abb.raiz = hijo
	} else {
		res := abb.cmp(nodo.clave, padre.clave)
		if res < 0 {
			padre.izquierdo = hijo
		} else {
			padre.derecho = hijo
		}
	}
	abb.cantidad--
	return borrado
}

// esta funcion borra un nodo que tiene dos hijos, haciendo que el nodo pase a ser el menor elemento de su derecha
func borrarDosHijos[K comparable, V any](abb *abb[K, V], nodo *nodoAbb[K, V]) V {
	borrado := nodo.dato
	padreReemplazante := buscarPadreReemplazo(abb, nodo.derecho)
	var reemplazante *nodoAbb[K, V]
	if padreReemplazante == nil { // el reemplazante es el de la derecha
		reemplazante = nodo.derecho
	} else {
		reemplazante = padreReemplazante.izquierdo // si el reemplazante no es el de la derecha, es el de la izquierda
	}
	clave := reemplazante.clave
	dato := abb.Borrar(clave)
	nodo.clave, nodo.dato = clave, dato
	return borrado
}

// esta funcion busca el padre del nodo mas chico del subarbol derecho
func buscarPadreReemplazo[K comparable, V any](abb *abb[K, V], nodoActual *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodoActual == nil || nodoActual.izquierdo == nil {
		return nil
	}
	if nodoActual.izquierdo.izquierdo == nil {
		return nodoActual
	}
	return buscarPadreReemplazo(abb, nodoActual.izquierdo)
}

// esta funcion apila todos los elementos a la izquierda del nodo recibido, respetando el limite "hasta"
func apilarIzquierdos[K comparable, V any](iter *iterDiccionarioOrdenado[K, V], nodoActual *nodoAbb[K, V]) {
	if nodoActual == nil {
		return
	}
	if (iter.desde == nil || iter.abb.cmp(nodoActual.clave, *iter.desde) >= 0) && (iter.hasta == nil || iter.abb.cmp(nodoActual.clave, *iter.hasta) <= 0) {
		iter.pila.Apilar(nodoActual)
	} else if iter.desde == nil || iter.abb.cmp(nodoActual.clave, *iter.desde) < 0 {
		apilarIzquierdos(iter, nodoActual.derecho)
	}
	apilarIzquierdos(iter, nodoActual.izquierdo)
}
