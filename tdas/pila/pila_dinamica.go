package pila

const _CAPACIDAD_INICIAL = 13
const _CONDICION_ACHICAR = 4
const _FACTOR_REDIMENSION = 2

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL)

	return pila
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		nuevaCapacidad := cap(pila.datos) * _FACTOR_REDIMENSION
		pila.redimensionarPila(nuevaCapacidad)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) verificarPilaVacia() {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
}

func (pila *pilaDinamica[T]) Desapilar() T {
	pila.verificarPilaVacia()
	if (pila.cantidad*_CONDICION_ACHICAR) <= cap(pila.datos) && (cap(pila.datos)/_FACTOR_REDIMENSION) > _CAPACIDAD_INICIAL {
		nuevaCapacidad := cap(pila.datos) / _FACTOR_REDIMENSION
		pila.redimensionarPila(nuevaCapacidad)
	}
	elemento := pila.datos[pila.cantidad-1]
	pila.cantidad--

	return elemento
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	pila.verificarPilaVacia()
	elemento := pila.datos[pila.cantidad-1]

	return elemento
}

func (pila *pilaDinamica[T]) redimensionarPila(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, pila.datos)
	pila.datos = nuevosDatos
}
