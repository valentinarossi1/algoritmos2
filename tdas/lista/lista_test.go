package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	//verifica que una lista vacia se comporte como tal
	lista := TDALista.CrearListaEnlazada[any]()

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.True(t, lista.EstaVacia())
}
func EnlistarListaGenerica[T any](dato1, dato2, dato3, dato4 T) TDALista.Lista[T] {
	//funcion auxiliar para enlistar datos de cualquier tipo a una lista
	lista := TDALista.CrearListaEnlazada[T]()

	lista.InsertarUltimo(dato1)
	lista.InsertarUltimo(dato2)
	lista.InsertarUltimo(dato3)
	lista.InsertarUltimo(dato4)

	return lista
}

func TestEnlistarYVaciaEnteros(t *testing.T) {
	//verifica que al enlist y borrar elemento, la lista quede vacia
	listaEnteros := EnlistarListaGenerica(2, 20, 200, 2000)

	require.Equal(t, 2, listaEnteros.BorrarPrimero())
	require.Equal(t, 20, listaEnteros.BorrarPrimero())
	require.Equal(t, 200, listaEnteros.BorrarPrimero())
	require.Equal(t, 2000, listaEnteros.BorrarPrimero())
	require.True(t, listaEnteros.EstaVacia())
}

func TestEnlistarYVaciaCadenas(t *testing.T) {
	//verifica que al enlistar y borrar cadenas, la lista quede vacia
	listaCadenas := EnlistarListaGenerica("primero", "segundo", "tercero", "cuarto")

	require.Equal(t, "primero", listaCadenas.BorrarPrimero())
	require.Equal(t, "segundo", listaCadenas.BorrarPrimero())
	require.Equal(t, "tercero", listaCadenas.BorrarPrimero())
	require.Equal(t, "cuarto", listaCadenas.BorrarPrimero())
	require.True(t, listaCadenas.EstaVacia())
}

// PRUEBAS ITERADOR EXTERNO
// Al insertar un elemento en la posición en la que se crea el iterador,
// efectivamente se inserta al principio.
func TestInsertarAlInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, iter.VerActual())
}

// Insertar un elemento cuando el iterador está al final efectivamente es
// equivalente a insertar al final.
func TestInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	iter.Insertar(2)
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	iter.Insertar(3)
	require.Equal(t, 3, lista.VerUltimo())
	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

// Insertar un elemento en el medio se hace en la posición correcta.
func TestInsertarEnElMedio(t *testing.T) {
	listaEnteros := EnlistarListaGenerica(2, 20, 200, 2000)
	iter := listaEnteros.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 200, iter.VerActual())
	iter.Insertar(1)
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 200, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 2000, iter.VerActual())
	require.Equal(t, 2000, listaEnteros.VerUltimo())
}

// Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
func TestEliminarPrimero(t *testing.T) {
	listaEnteros := EnlistarListaGenerica(2, 20, 200, 2000)
	iter := listaEnteros.Iterador()
	require.Equal(t, 2, iter.Borrar())
	require.Equal(t, 20, listaEnteros.VerPrimero())
}

// Remover el último elemento con el iterador cambia el último de la lista.
func TestEliminarUltimo(t *testing.T) {
	listaEnteros := EnlistarListaGenerica(2, 20, 200, 2000)
	iter := listaEnteros.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 2000, iter.VerActual())
	require.Equal(t, 2000, iter.Borrar())
	require.Equal(t, 200, listaEnteros.VerUltimo())
}

// Verificar que al remover un elemento del medio, este no está.
func TestEliminarMedio(t *testing.T) {
	listaEnteros := EnlistarListaGenerica(1, 2, 3, 4)
	iter := listaEnteros.Iterador()
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 2, iter.VerActual())
	require.Equal(t, 2, iter.Borrar())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 4, iter.VerActual())
}

// PRUEBAS ITERADOR INTERNO

func EnlistarListaGenericaLarga[T any](dato1, dato2, dato3, dato4, dato5, dato6, dato7 T) TDALista.Lista[T] {
	//funcion auxiliar para enlistar datos de cualquier tipo a una lista
	lista := TDALista.CrearListaEnlazada[T]()

	lista.InsertarUltimo(dato1)
	lista.InsertarUltimo(dato2)
	lista.InsertarUltimo(dato3)
	lista.InsertarUltimo(dato4)
	lista.InsertarUltimo(dato5)
	lista.InsertarUltimo(dato6)
	lista.InsertarUltimo(dato7)

	return lista
}
func TestSinCorte(t *testing.T) {
	listaEnteros := EnlistarListaGenericaLarga(1, 2, 3, 4, 5, 6, 7)
	suma := 0
	listaEnteros.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.Equal(t, 28, suma)
}

func TestConCorte(t *testing.T) {
	listaEnteros := EnlistarListaGenericaLarga(1, 2, 3, 4, 5, 6, 7)
	suma, contador := 0, 0
	listaEnteros.Iterar(func(v int) bool {
		if contador == 2 {
			return false
		}
		if v%2 == 0 {
			suma += v
			contador += 1
		}
		return true
	})
	require.Equal(t, 6, suma)
}
