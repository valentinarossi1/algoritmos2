package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const _PRUEBA_VOLUMEN = 800

func TestPilaVacia(t *testing.T) {
	//verifica que una pila vacia se comporte como tal
	pila := TDAPila.CrearPilaDinamica[any]()

	pila.Apilar(1)
	pila.Desapilar()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func ApilarPilaGenerica[T any](dato1, dato2, dato3, dato4 T) TDAPila.Pila[T] {
	//funcion auxiliar para apilar datos de cualquier tipo a una pila
	pila := TDAPila.CrearPilaDinamica[T]()

	pila.Apilar(dato1)
	pila.Apilar(dato2)
	pila.Apilar(dato3)
	pila.Apilar(dato4)

	return pila
}

func TestInvarianteEnteros(t *testing.T) {
	//verifica que se mantenga el invariante LIFO en una pila de enteros
	pilaEnteros := ApilarPilaGenerica(2, 20, 200, 2000)

	require.Equal(t, 2000, pilaEnteros.Desapilar())
	require.Equal(t, 200, pilaEnteros.Desapilar())
	require.Equal(t, 20, pilaEnteros.Desapilar())
	require.Equal(t, 2, pilaEnteros.Desapilar())
	require.True(t, pilaEnteros.EstaVacia())
}

func TestInvarianteCadenas(t *testing.T) {
	//verifica que se mantenga el invariante LIFO en una pila de cadenas
	pilaCadenas := ApilarPilaGenerica("cuarto", "tercero", "segundo", "primero")

	require.Equal(t, "primero", pilaCadenas.VerTope())
	require.Equal(t, "primero", pilaCadenas.Desapilar())
	require.Equal(t, "segundo", pilaCadenas.Desapilar())
	require.Equal(t, "tercero", pilaCadenas.Desapilar())
	require.Equal(t, "cuarto", pilaCadenas.Desapilar())
	require.True(t, pilaCadenas.EstaVacia())
}

func TestInvarianteBooleanos(t *testing.T) {
	//verifica que se mantenga el invariante LIFO en una pila de booleanos
	pilaBooleanos := ApilarPilaGenerica(true, false, true, true)

	require.Equal(t, true, pilaBooleanos.Desapilar())
	require.Equal(t, true, pilaBooleanos.Desapilar())
	require.Equal(t, false, pilaBooleanos.Desapilar())
	require.Equal(t, true, pilaBooleanos.Desapilar())
	require.True(t, pilaBooleanos.EstaVacia())
}

func TestVolumen(t *testing.T) {
	//prueba que se pueden apilar y desapilar muchos elementos cumpliendo el invariante, y valida que el tope siempre sea el correcto
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i <= _PRUEBA_VOLUMEN; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for j := _PRUEBA_VOLUMEN; j >= 0; j-- {
		require.Equal(t, j, pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestEnteros(t *testing.T) {
	//comprueba que al desapilar una pila de enteros hasta que está vacía se comporta como recién creada
	pilaEnteros := ApilarPilaGenerica(1, 2, 3, 4)

	pilaEnteros.Desapilar()
	pilaEnteros.Desapilar()
	pilaEnteros.Desapilar()
	pilaEnteros.Desapilar()

	require.True(t, pilaEnteros.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaEnteros.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaEnteros.VerTope() })

	pilaEnteros.Apilar(50)
	require.False(t, pilaEnteros.EstaVacia())
	require.Equal(t, 50, pilaEnteros.VerTope())
}

func TestBooleanos(t *testing.T) {
	//comprueba que al desapilar una pila de booleanos hasta que está vacía se comporta como recién creada.
	pilaBooleanos := ApilarPilaGenerica(true, false, true, false)

	pilaBooleanos.Desapilar()
	pilaBooleanos.Desapilar()
	pilaBooleanos.Desapilar()
	pilaBooleanos.Desapilar()

	require.True(t, pilaBooleanos.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaBooleanos.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaBooleanos.VerTope() })

	pilaBooleanos.Apilar(false)
	require.False(t, pilaBooleanos.EstaVacia())
	require.Equal(t, false, pilaBooleanos.VerTope())
}

func TestCadenas(t *testing.T) {
	//comprueba que al desapilar una pila de cadenas hasta que está vacía se comporta como recién creada.
	pilaCadenas := ApilarPilaGenerica("hola", "hello", "oi", "bonjour")

	pilaCadenas.Desapilar()
	pilaCadenas.Desapilar()
	pilaCadenas.Desapilar()
	pilaCadenas.Desapilar()

	require.True(t, pilaCadenas.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaCadenas.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaCadenas.VerTope() })

	pilaCadenas.Apilar("hallo")
	require.False(t, pilaCadenas.EstaVacia())
	require.Equal(t, "hallo", pilaCadenas.VerTope())
}

func TestRecienCreada(t *testing.T) {
	//prueba que las acciones de Desapilar y VerTope en una pila recien creada son invalidas
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

}

func TestRecienCreadaVacia(t *testing.T) {
	//prueba que la accion de EstaVacia en una pila recien creada es invalida
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
}
