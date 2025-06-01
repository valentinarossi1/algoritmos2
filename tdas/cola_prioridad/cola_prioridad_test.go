package cola_prioridad_test

import (
	"fmt"
	"math/rand"
	"strings"

	TDACOLAPRIORIDAD "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const _PRUEBA_HEAPSORT = 10000

var TAMS_VOLUMEN_ORDENADO = []int{12500, 25000, 50000, 100000, 200000, 400000}

func comparar(num1, num2 int) int {
	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	} else {
		return 0
	}
}

func TestColaVacia(t *testing.T) {
	t.Log("Comprueba que heap vacio se comporte como tal")
	heap := TDACOLAPRIORIDAD.CrearHeap[int](comparar)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	heap.Encolar(1)
	require.False(t, heap.EstaVacia())
}

func TestArregloVacio(t *testing.T) {
	t.Log("Comprueba que un heap creado desde un arreglo vacio se comporte correctamente")
	var arreglo []int
	heap := TDACOLAPRIORIDAD.CrearHeapArr[int](arreglo, comparar)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	heap.Encolar(1)
	require.False(t, heap.EstaVacia())
}

func TestInicial(t *testing.T) {
	t.Log("Comprueba que el heap puede guardar elementos y eliminarlos")
	heap := TDACOLAPRIORIDAD.CrearHeap[int](comparar)
	heap.Encolar(1)
	heap.Encolar(2)
	heap.Encolar(3)
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())
	require.False(t, heap.EstaVacia())
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())
	require.False(t, heap.EstaVacia())
	heap.Desencolar()
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
}

func TestInicialArreglo(t *testing.T) {
	t.Log("Comprueba que el heap creado con un arreglo puede guardar elementos y eliminarlos")
	arreglo := []int{4, 6, 9, -4, 0, 7, 3}
	heap := TDACOLAPRIORIDAD.CrearHeapArr[int](arreglo, comparar)
	require.Equal(t, 7, heap.Cantidad())
	require.Equal(t, 9, heap.VerMax())
	require.False(t, heap.EstaVacia())
	require.Equal(t, 9, heap.Desencolar())
	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())
	require.False(t, heap.EstaVacia())
	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()
	require.Equal(t, -4, heap.VerMax())
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que un heap con un elemento se comporte como tal")
	heap := TDACOLAPRIORIDAD.CrearHeap[string](strings.Compare)
	heap.Encolar("A")
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, "A", heap.VerMax())
	require.False(t, heap.EstaVacia())
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
}

func TestHeapConDuplicados(t *testing.T) {
	t.Log("Comprueba que un heap con elementos duplicados funciona correctamente")
	heap := TDACOLAPRIORIDAD.CrearHeap[int](comparar)
	heap.Encolar(2)
	heap.Encolar(2)
	heap.Encolar(2)
	heap.Encolar(1)
	heap.Encolar(1)
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 2, heap.VerMax())
	require.EqualValues(t, 2, heap.Desencolar())
	require.EqualValues(t, 2, heap.VerMax())
	require.EqualValues(t, 2, heap.Desencolar())
	require.EqualValues(t, 2, heap.VerMax())
	require.EqualValues(t, 2, heap.Desencolar())
	require.EqualValues(t, 1, heap.VerMax())
}

func TestVolumenEncolarDesencolar(t *testing.T) {
	t.Log("Comprueba que un heap pueda encolar y desencolar muchos elementos ordenados")
	heap := TDACOLAPRIORIDAD.CrearHeap(strings.Compare)
	for i := 0; i <= 1000; i++ {
		heap.Encolar("i")
		require.EqualValues(t, i+1, heap.Cantidad())
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, "i", heap.VerMax())
	}
	for i := 1000; i >= 0; i-- {
		require.EqualValues(t, "i", heap.VerMax())
		require.EqualValues(t, "i", heap.Desencolar())
		require.EqualValues(t, i, heap.Cantidad())
	}
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

// esta funcion mezcla un arreglo intercambiando aleatoriamente sus elementos
func mezclarArreglo(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func ejecutarPruebaVolumenOrd(b *testing.B, n int) {
	heap := TDACOLAPRIORIDAD.CrearHeap[int](comparar)
	valores := make([]int, n)
	for i := 0; i < n; i++ {
		valores[i] = i
	}
	valores = mezclarArreglo(valores)

	for i := range valores {
		heap.Encolar(valores[i])
	}

	require.EqualValues(b, n, heap.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = (heap.Desencolar() == len(valores)-1-i)
		if !ok {
			break
		}
	}

	require.True(b, ok, "Desecolar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, heap.Cantidad())
}

func BenchmarkHeapDesordenado(b *testing.B) {
	b.Log("Encola distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que VerMax funcione para cada elemento " +
		"y que luego podemos desencolar sin problemas")
	for _, n := range TAMS_VOLUMEN_ORDENADO {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenOrd(b, n)
			}
		})
	}
}

func TestHeapSortDesordenado(t *testing.T) {
	t.Log("Comprueba que HeapSort ordenea bien un arreglo desordenado")
	arreglo := []int{4, 6, 9, 7, 3, -16, 0}
	arregloOrdenado := []int{-16, 0, 3, 4, 6, 7, 9}
	arreglo = TDACOLAPRIORIDAD.HeapSort(arreglo, comparar)
	require.EqualValues(t, arregloOrdenado, arreglo)
}

func TestHeapSortOrdenado(t *testing.T) {
	t.Log("Comprueba que si HeapSort recibe un arreglo ordenado, no lo modifica ")
	arreglo := []int{1, 2, 3, 4, 5}
	arregloOrdenado := []int{1, 2, 3, 4, 5}
	arreglo = TDACOLAPRIORIDAD.HeapSort(arreglo, comparar)
	require.EqualValues(t, arregloOrdenado, arreglo)
}

func TestHeapSortVolumen(t *testing.T) {
	t.Log("Comprueba que HeapSort funciona correctamente para un arreglo de muchos elementos desordenados")
	arregloOrdenado := make([]int, _PRUEBA_HEAPSORT)
	for i := 0; i < _PRUEBA_HEAPSORT; i++ {
		arregloOrdenado[i] = i
	}
	arreglo := mezclarArreglo(arregloOrdenado)
	arreglo = TDACOLAPRIORIDAD.HeapSort(arreglo, comparar)
	require.EqualValues(t, arreglo, arregloOrdenado)
}
