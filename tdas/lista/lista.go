package lista

type Lista[T any] interface {
	// EstaVacia devuelve verdadero si la lista no tiene elementos , false en caso contrario.
	EstaVacia() bool
	InsertarPrimero(T)
	InsertarUltimo(T)
	BorrarPrimero() T
	VerPrimero() T
	VerUltimo() T
	Largo() int
	Iterar(visitar func(T) bool)
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	VerActual() T
	HaySiguiente() bool
	Siguiente()
	Insertar(T)
	Borrar() T
}
