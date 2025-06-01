package diccionario

import (
	"fmt"
	"hash/fnv"
)

const (
	_BORRADO           = "BORRADO"
	_VACIO             = ""
	_OCUPADO           = "OCUPADO"
	_CAPACIDAD_INICIAL = 97
	_FACTOR_AGRANDAR   = 2
	_FACTOR_ACHICAR    = 0.5
	_CRITERIO_AGRANDAR = 0.7
	_CRITERIO_ACHICAR  = 0.2
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado string
}

type hashCerrado[K comparable, V any] struct {
	tabla     []celdaHash[K, V]
	cantidad  int
	capacidad int
	borrados  int
}

type iterDiccionario[K comparable, V any] struct {
	posicion int
	hash     *hashCerrado[K, V]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.capacidad = _CAPACIDAD_INICIAL
	hash.tabla = crearTabla[K, V](hash.capacidad)

	return hash
}

func crearTabla[K comparable, V any](capacidad int) []celdaHash[K, V] {
	tabla := make([]celdaHash[K, V], capacidad)
	return tabla
}

/* func redimensionar[K comparable, V any](hash *hashCerrado[K, V], factor float32) *hashCerrado[K, V] {
	nuevoHash := new(hashCerrado[K, V])
	nuevaCapacidad := int(float32(hash.capacidad) * factor)
	nuevoHash.capacidad = nuevaCapacidad
	nuevoHash.tabla = crearTabla[K, V](nuevaCapacidad)
	nuevoHash.cantidad = hash.cantidad
	for i := 0; i < hash.capacidad; i++ {
		if hash.tabla[i].estado == _OCUPADO {
			clave := hash.tabla[i].clave
			dato := hash.tabla[i].dato
			nuevoHash.Guardar(clave, dato)
		}
	}
	nuevoHash.cantidad = hash.cantidad
	fmt.Printf("redimensione con cap %d: y la nueva capacidad es :%d", hash.capacidad, nuevoHash.capacidad)
	return nuevoHash
} */

// esta funcion crea un nuevo hash con una nueva capacidad. si la capacidad es menor a la inicial, deja la inicial. luego, pasa los elementos que estan ocupados al hash creado y lo devuelve
func redimensionar[K comparable, V any](hash *hashCerrado[K, V], factor float32) *hashCerrado[K, V] {
	nuevoHash := new(hashCerrado[K, V])
	nuevaCapacidad := int(float32(hash.capacidad) * factor)
	nuevoHash.capacidad = nuevaCapacidad
	nuevoHash.tabla = crearTabla[K, V](nuevaCapacidad)
	for i := 0; i < hash.capacidad; i++ {
		if hash.tabla[i].estado == _OCUPADO {
			clave := hash.tabla[i].clave
			dato := hash.tabla[i].dato
			pos := buscarPosicion(nuevoHash, clave, nuevaCapacidad)
			nuevoHash.tabla[pos] = celdaHash[K, V]{clave, dato, _OCUPADO}
		}
	}
	nuevoHash.cantidad = hash.cantidad
	return nuevoHash
}

// esta funcion verifica el valor del factor de carga y devuelve True si hay que redimensionar
func hayQueRedimensionar[K comparable, V any](hash *hashCerrado[K, V]) bool {
	factorDeCarga := float32(hash.cantidad+hash.borrados) / float32(hash.capacidad)
	return factorDeCarga > _CRITERIO_AGRANDAR || (factorDeCarga < _CRITERIO_ACHICAR && float32(hash.capacidad)*_FACTOR_ACHICAR > _CAPACIDAD_INICIAL)
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	posicion := buscarPosicion(hash, clave, hash.capacidad)
	if !hash.Pertenece(clave) { // aumento la cantidad solo si no estoy reemplazanado un dato
		hash.cantidad++
	}
	celda := celdaHash[K, V]{clave, dato, _OCUPADO}
	hash.tabla[posicion] = celda
	if hayQueRedimensionar(hash) { // duplicamos la capacidad
		*hash = *redimensionar(hash, _FACTOR_AGRANDAR)
	}
}

// esta funcion transforma un tipo de dato generico a un array de bytes
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// esta funcion devuelve la posicion que le corresponderia a una clave, segun la funcion de hashing
func calcularPosicion[K comparable](clave K, capacidad int) uint32 {
	claveByte := convertirABytes(clave)
	posicion := hashing(claveByte) % uint32(capacidad)
	return posicion
}

// funcion de hashing proporcionada por el paquete "hash/fnv" de Go
func hashing(s []byte) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	posicion := buscarPosicion(hash, clave, hash.capacidad)
	return hash.tabla[posicion].estado == _OCUPADO
}

// esta funcion busca la posicion verdadera de una clave, fijandose si la posicion brindada por el hashing esta ocupada por otra clave
func buscarPosicion[K comparable, V any](hash *hashCerrado[K, V], clave K, capacidad int) uint32 {
	posicion := calcularPosicion(clave, capacidad)
	for hash.tabla[posicion].estado != _VACIO && hash.tabla[posicion].clave != clave {
		posicion++
		posicion = posicion % uint32(capacidad)
	}
	return posicion
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := buscarPosicion(hash, clave, hash.capacidad)
	if hash.tabla[posicion].estado == _OCUPADO {
		return hash.tabla[posicion].dato
	}
	panic("La clave no pertenece al diccionario")
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	posicion := buscarPosicion(hash, clave, hash.capacidad)
	if hash.tabla[posicion].estado != _OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
	valorBorrado := hash.tabla[posicion].dato
	hash.tabla[posicion].estado = _BORRADO
	hash.cantidad--
	hash.borrados++
	if hayQueRedimensionar(hash) {
		*hash = *redimensionar(hash, _FACTOR_ACHICAR)
	}
	return valorBorrado
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

// esta funcion verifica si hay un siguiente elemento. si no lo hay, entra en panic
func (iterador *iterDiccionario[K, V]) verificarIteradorFinalizado() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iterador *iterDiccionario[K, V]) HaySiguiente() bool {
	return iterador.posicion < iterador.hash.capacidad
}

func (iterador *iterDiccionario[K, V]) VerActual() (K, V) {
	iterador.verificarIteradorFinalizado()
	return iterador.hash.tabla[iterador.posicion].clave, iterador.hash.tabla[iterador.posicion].dato
}

func (iterador *iterDiccionario[K, V]) Siguiente() {
	iterador.verificarIteradorFinalizado()
	iterador.posicion++
	for iterador.HaySiguiente() && iterador.hash.tabla[iterador.posicion].estado != _OCUPADO {
		iterador.posicion++
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iteradorDiccionario := new(iterDiccionario[K, V])
	iteradorDiccionario.hash = hash
	iteradorDiccionario.posicion = 0
	for iteradorDiccionario.posicion < iteradorDiccionario.hash.capacidad && iteradorDiccionario.hash.tabla[iteradorDiccionario.posicion].estado != _OCUPADO {
		iteradorDiccionario.posicion++
	}
	return iteradorDiccionario
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	visitado := true
	for i := 0; i < hash.capacidad && visitado; i++ {
		if hash.tabla[i].estado == _OCUPADO {
			visitado = visitar(hash.tabla[i].clave, hash.tabla[i].dato)
		}
	}
}
