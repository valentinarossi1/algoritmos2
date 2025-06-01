package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ORDENADO = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioVacioOrd(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestInicial(t *testing.T) {
	t.Log("Comprueba que el Diccionario puede guardar elementos, reemplazarlos, y eliminarlos")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("A", "hola")
	dic.Guardar("B", "hola")
	dic.Guardar("A", "chau")
	require.Equal(t, true, dic.Pertenece("A"))
	require.Equal(t, true, dic.Pertenece("B"))
	require.Equal(t, 2, dic.Cantidad())
	require.Equal(t, "chau", dic.Obtener("A"))
	dic.Borrar("A")
	dic.Borrar("B")
	require.False(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
}

func TestUnElementOrdenado(t *testing.T) {
	t.Log("Comprueba que un Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardarOrdenado(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoOrdenado(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func comparar(num1, num2 int) int {
	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	} else {
		return 0
	}
}

func TestBorrarRaiz(t *testing.T) {
	t.Log("Verifica que se puede borrar la raiz de un ABB correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	dic.Borrar(4)
	require.False(t, dic.Pertenece(4))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(4) })
	require.Equal(t, 2, dic.Cantidad())
}

func TestReemplazoDatoHopscotchOrdenado(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioBorrarOrdenado(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))

	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestConClavesNumericasOrdenado(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	clave := 10
	valor := 20

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestClaveVaciaOrdenado(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNuloOrdenado(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestGuardarYBorrarRepetidasVecesOrdenado(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces")

	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func TestIterarDiccionarioVacioOrdenado(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorConRangoVacio(t *testing.T) {
	t.Log("Valida que iterar sobre un diccionario vacio es simplemente tenerlo al final, aunque reciba un rango")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	desde := 10
	hasta := 100
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func BuscarOrd(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestDiccionarioIterarOrd(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, BuscarOrd(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, BuscarOrd(segundo, claves))
	require.EqualValues(t, valores[BuscarOrd(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, BuscarOrd(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorPorRangos(t *testing.T) {
	t.Log("Valida que el iterador recorra el ABB respetando limites inferior y superior")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 50; i < 200; i++ {
		dic.Guardar(i, i*2)
	}
	desde := 60
	hasta := 150
	iter := dic.IteradorRango(&desde, &hasta)
	for i := 60; i <= 150; i++ {
		clave, _ := iter.VerActual()

		require.EqualValues(t, true, iter.HaySiguiente())
		require.EqualValues(t, i, clave)
		iter.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinalOrd(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, BuscarOrd(primero, claves))
	require.NotEqualValues(t, -1, BuscarOrd(segundo, claves))
	require.NotEqualValues(t, -1, BuscarOrd(tercero, claves))
}

func TestIteradorPorRangosCompleto(t *testing.T) {
	t.Log("Valida que el iterador recorra todo el ABB cuando no se le establecen limites")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 0; i < 200; i++ {
		dic.Guardar(i, i)
	}
	iter := dic.IteradorRango(nil, nil)
	for i := 0; i < dic.Cantidad(); i++ {
		clave, _ := iter.VerActual()
		require.EqualValues(t, true, iter.HaySiguiente())
		require.EqualValues(t, i, clave)
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorPorRangosDesde(t *testing.T) {
	t.Log("Valida que el iterador recorra el ABB respetando el limite inferior")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 0; i < 15; i++ {
		dic.Guardar(i, i)
	}
	desde := 4
	iter := dic.IteradorRango(&desde, nil)
	for i := 4; i < dic.Cantidad(); i++ {
		clave, _ := iter.VerActual()
		require.EqualValues(t, true, iter.HaySiguiente())
		require.EqualValues(t, i, clave)
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorPorRangosHasta(t *testing.T) {
	t.Log("Valida que el iterador recorra el ABB respetando el limite superior")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 0; i < 15; i++ {
		dic.Guardar(i, i)
	}
	hasta := 9
	iter := dic.IteradorRango(nil, &hasta)
	for i := 0; i <= 9; i++ {
		clave, _ := iter.VerActual()
		require.EqualValues(t, true, iter.HaySiguiente())
		require.EqualValues(t, i, clave)
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorPorRangosDesdeInferior(t *testing.T) {
	t.Log("Valida que el iterador recorra el ABB respetando limites inferior y superior, aunque el inferior sea menor que la clave mas chica del diccionario")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 50; i < 100; i++ {
		dic.Guardar(i, i*2)
	}
	desde := 40
	hasta := 54
	iter := dic.IteradorRango(&desde, &hasta)
	clave1, _ := iter.VerActual()
	require.Equal(t, 50, clave1)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave2, _ := iter.VerActual()
	require.Equal(t, 51, clave2)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave3, _ := iter.VerActual()
	require.Equal(t, 52, clave3)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave4, _ := iter.VerActual()
	require.Equal(t, 53, clave4)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave5, _ := iter.VerActual()
	require.Equal(t, 54, clave5)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarFueraDeRango(t *testing.T) {
	t.Log("Iterar fuera del rango hace que el iterador esté al final")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 0; i < 20; i++ {
		dic.Guardar(i, i)
	}
	desde := 21
	hasta := 30
	iter := dic.IteradorRango(&desde, &hasta)

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarRangoCruzado(t *testing.T) {
	t.Log("Iterar con un rango desde > hasta hace que el iterador esté al final")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	for i := 0; i < 25; i++ {
		dic.Guardar(i, i)
	}
	desde := 7
	hasta := 3
	iter := dic.IteradorRango(&desde, &hasta)

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorInternoClavesOrd(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad
	desde := clave1
	hasta := clave3
	dic.IterarRango(&desde, &hasta, func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, BuscarOrd(cs[0], claves))
	require.NotEqualValues(t, -1, BuscarOrd(cs[1], claves))
	require.NotEqualValues(t, -1, BuscarOrd(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValoresOrd(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente (y una única vez) con el iterador interno")
	clave1 := 6
	clave2 := 4
	clave3 := 8
	clave4 := 20
	clave5 := 5

	dic := TDADiccionario.CrearABB[int, int](comparar)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	suma := 0
	sumaAux := &suma
	dic.IterarRango(nil, nil, func(clave int, dato int) bool {
		*sumaAux += dato
		return true
	})

	require.EqualValues(t, 20, suma)
}

func TestIteradorInternoValoresConBorradosOrd(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave1 := 6
	clave2 := 4
	clave3 := 8
	clave4 := 20
	clave5 := 5

	dic := TDADiccionario.CrearABB[int, int](comparar)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave1)

	factorial := 1
	ptrFactorial := &factorial
	dic.IterarRango(nil, nil, func(_ int, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 120, factorial)
}

func TestIteradorInternoConCorte(t *testing.T) {
	t.Log("Valida que el iterador interno recorra correctamente el ABB, respetando la condicion de corte")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	clave1 := 6
	clave2 := 4
	clave3 := 8
	clave4 := 2
	clave5 := 5

	dic.Guardar(clave1, 60)
	dic.Guardar(clave2, 22)
	dic.Guardar(clave3, 37)
	dic.Guardar(clave4, 49)
	dic.Guardar(clave5, 59)

	suma := 0
	sumaAux := &suma
	dic.IterarRango(nil, nil, func(_ int, dato int) bool {
		*sumaAux += dato
		return suma < 100
	})
	require.Equal(t, 130, suma)
}

func TestIteradorInternoRangoDesdeConCorte(t *testing.T) {
	t.Log("Valida que el iterador interno recorra correctamente el ABB, respetando el limite inferior y la condicion de corte")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	clave1 := 6
	clave2 := 4
	clave3 := 8
	clave4 := 2
	clave5 := 5
	clave6 := 19
	clave7 := 50

	dic.Guardar(clave1, 60)
	dic.Guardar(clave2, 22)
	dic.Guardar(clave3, 37)
	dic.Guardar(clave4, 49)
	dic.Guardar(clave5, 60)
	dic.Guardar(clave6, 43)
	dic.Guardar(clave7, 99)
	desde := clave1
	suma := 0
	sumaAux := &suma
	dic.IterarRango(&desde, nil, func(_ int, dato int) bool {
		*sumaAux += dato
		return suma < 200
	})
	require.Equal(t, 239, suma)
}

func TestIteradorInternoRangoHastaSinCorte(t *testing.T) {
	t.Log("Valida que el iterador interno recorra correctamente el ABB, respetando el limite superior")
	dic := TDADiccionario.CrearABB[int, int](comparar)
	clave1 := 6
	clave2 := 4
	clave3 := 8
	clave4 := 2
	clave5 := 5
	clave6 := 19
	clave7 := 50

	dic.Guardar(clave1, 60)
	dic.Guardar(clave2, 22)
	dic.Guardar(clave3, 37)
	dic.Guardar(clave4, 49)
	dic.Guardar(clave5, 60)
	dic.Guardar(clave6, 43)
	dic.Guardar(clave7, 99)
	hasta := clave3
	suma := 0
	sumaAux := &suma
	dic.IterarRango(nil, &hasta, func(_ int, dato int) bool {
		*sumaAux += dato
		return true
	})
	require.Equal(t, 228, suma)
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
	dic := TDADiccionario.CrearABB[int, int](comparar)

	claves := make([]int, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = i
	}

	claves = mezclarArreglo(claves)
	valores = mezclarArreglo(valores)

	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioOrd(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_ORDENADO {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenOrd(b, n)
			}
		})
	}
}
