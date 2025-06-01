package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const _PRUEBA_VOLUMEN = 800

func TestColaVacia(t *testing.T) {
	//verifica que una cola vacia se comporte como tal
	cola := TDACola.CrearColaEnlazada[any]()

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func EncolarColaGenerica[T any](dato1, dato2, dato3, dato4 T) TDACola.Cola[T] {
	//funcion auxiliar para encolar datos de cualquier tipo a una cola
	cola := TDACola.CrearColaEnlazada[T]()

	cola.Encolar(dato1)
	cola.Encolar(dato2)
	cola.Encolar(dato3)
	cola.Encolar(dato4)

	return cola
}

func TestInvarianteEnteros(t *testing.T) {
	//verifica que se mantenga el invariante FIFO en una cola de enteros
	colaEnteros := EncolarColaGenerica(2, 20, 200, 2000)

	require.Equal(t, 2, colaEnteros.Desencolar())
	require.Equal(t, 20, colaEnteros.Desencolar())
	require.Equal(t, 200, colaEnteros.Desencolar())
	require.Equal(t, 2000, colaEnteros.Desencolar())
	require.True(t, colaEnteros.EstaVacia())
}

func TestInvarianteCadenas(t *testing.T) {
	//verifica que se mantenga el invariante FIFO en una cola de cadenas
	colaCadenas := EncolarColaGenerica("primero", "segundo", "tercero", "cuarto")

	require.Equal(t, "primero", colaCadenas.Desencolar())
	require.Equal(t, "segundo", colaCadenas.Desencolar())
	require.Equal(t, "tercero", colaCadenas.Desencolar())
	require.Equal(t, "cuarto", colaCadenas.Desencolar())
	require.True(t, colaCadenas.EstaVacia())
}

func TestInvarianteBooleanos(t *testing.T) {
	//verifica que se mantenga el invariante FIFO en una cola de booleanos
	colaBooleanos := EncolarColaGenerica(true, false, true, true)

	require.Equal(t, true, colaBooleanos.Desencolar())
	require.Equal(t, false, colaBooleanos.Desencolar())
	require.Equal(t, true, colaBooleanos.Desencolar())
	require.Equal(t, true, colaBooleanos.Desencolar())
	require.True(t, colaBooleanos.EstaVacia())
}

func TestVolumen(t *testing.T) {
	//prueba que se pueden encolar y desencolar muchos elementos cumpliendo el invariante, y valida que el primero sea el correcto
	cola := TDACola.CrearColaEnlazada[int]()

	for i := 0; i <= _PRUEBA_VOLUMEN; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}
	for j := 0; j <= _PRUEBA_VOLUMEN; j++ {
		require.Equal(t, j, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestEnteros(t *testing.T) {
	//comprueba que al desencolar una cola de enteros hasta que está vacía se comporta como recién creada
	colaEnteros := EncolarColaGenerica(1, 2, 3, 4)

	colaEnteros.Desencolar()
	colaEnteros.Desencolar()
	colaEnteros.Desencolar()
	colaEnteros.Desencolar()

	require.True(t, colaEnteros.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaEnteros.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaEnteros.VerPrimero() })

	colaEnteros.Encolar(50)
	require.False(t, colaEnteros.EstaVacia())
	require.Equal(t, 50, colaEnteros.VerPrimero())
}

func TestBooleanos(t *testing.T) {
	//comprueba que al desencolar una cola de booleanos hasta que está vacía se comporta como recién creada
	colaBooleanos := EncolarColaGenerica(true, false, true, false)

	colaBooleanos.Desencolar()
	colaBooleanos.Desencolar()
	colaBooleanos.Desencolar()
	colaBooleanos.Desencolar()

	require.True(t, colaBooleanos.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaBooleanos.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaBooleanos.VerPrimero() })

	colaBooleanos.Encolar(false)
	require.False(t, colaBooleanos.EstaVacia())
	require.Equal(t, false, colaBooleanos.VerPrimero())
}

func TestCadenas(t *testing.T) {
	//comprueba que al desencolar una cola de cadenas hasta que está vacía se comporta como recién creada
	colaCadenas := EncolarColaGenerica("hola", "hello", "oi", "bonjour")

	colaCadenas.Desencolar()
	colaCadenas.Desencolar()
	colaCadenas.Desencolar()
	colaCadenas.Desencolar()

	require.True(t, colaCadenas.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaCadenas.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaCadenas.VerPrimero() })

	colaCadenas.Encolar("hallo")
	require.False(t, colaCadenas.EstaVacia())
	require.Equal(t, "hallo", colaCadenas.VerPrimero())
}

func TestRecienCreada(t *testing.T) {
	//prueba que las acciones de Desencolar y VerPrimero en una cola recien creada son invalidas
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })

}

func TestRecienCreadaVacia(t *testing.T) {
	//prueba que la accion de EstaVacia en una cola recien creada es verdadera
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
}
