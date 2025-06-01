package operaciones

import (
	"fmt"
	"math"
	TDAPila "tdas/pila"
)

type Operacion struct {
	operador string
	aridad   int                                    // aridad es la cantidad de operandos que necesita la operacion
	operar   func(operandos []int64) (int64, error) // operandos debe ser un slice de "aridad" operandos en el orden correcto
}

func suma(operandos []int64) (int64, error) {
	return operandos[0] + operandos[1], nil
}

func resta(operandos []int64) (int64, error) {
	return operandos[1] - operandos[0], nil
}

func multiplicacion(operandos []int64) (int64, error) {
	return operandos[0] * operandos[1], nil
}

func division(operandos []int64) (int64, error) {
	if operandos[0] != 0 {
		return operandos[1] / operandos[0], nil
	}
	return 0, fmt.Errorf("ERROR")
}

func potencia(operandos []int64) (int64, error) {
	if operandos[0] >= 0 {
		res := int64(math.Pow(float64(operandos[1]), float64(operandos[0])))
		return res, nil
	}
	return 0, fmt.Errorf("ERROR")
}

func raiz(operandos []int64) (int64, error) {
	if operandos[0] >= 0 {
		res := int64(math.Sqrt(float64(operandos[0])))
		return res, nil
	}
	return 0, fmt.Errorf("ERROR")
}

func logaritmo(operandos []int64) (int64, error) {
	if operandos[0] > 1 && operandos[1] > 0 {
		res := int64(math.Log(float64(operandos[1])) / math.Log(float64(operandos[0])))
		return res, nil
	}
	return 0, fmt.Errorf("ERROR")
}

func opTernario(operandos []int64) (int64, error) {
	if operandos[2] == 0 {
		res := operandos[0]
		return res, nil
	} else if operandos[2] != 0 {
		res := operandos[1]
		return res, nil
	}
	return 0, fmt.Errorf("ERROR")
}

func desapilarOperandos(pila TDAPila.Pila[int64], aridad int) ([]int64, bool) {
	operandos := make([]int64, aridad)
	for i := 0; i < aridad; i++ {
		if pila.EstaVacia() {
			return operandos, false
		}
		operando := pila.Desapilar()
		operandos[i] = operando
	}
	return operandos, true
}

func identificarOperador(token string) (Operacion, bool) {
	switch token {
	case "+":
		return Operacion{operador: "+", aridad: 2, operar: suma}, true
	case "-":
		return Operacion{operador: "-", aridad: 2, operar: resta}, true
	case "*":
		return Operacion{operador: "*", aridad: 2, operar: multiplicacion}, true
	case "/":
		return Operacion{operador: "/", aridad: 2, operar: division}, true
	case "sqrt":
		return Operacion{operador: "sqrt", aridad: 1, operar: raiz}, true
	case "^":
		return Operacion{operador: "^", aridad: 2, operar: potencia}, true
	case "log":
		return Operacion{operador: "log", aridad: 2, operar: logaritmo}, true
	case "?":
		return Operacion{operador: "?", aridad: 3, operar: opTernario}, true
	default:
		return Operacion{}, false
	}
}

func Calcular(token string, pila TDAPila.Pila[int64]) bool {
	operacion, operadorValido := identificarOperador(token)
	if operadorValido {
		operandos, exito := desapilarOperandos(pila, operacion.aridad)
		if exito {
			resultado, err := operacion.operar(operandos)
			if err == nil {
				pila.Apilar(resultado)
				return true
			}
		}
	}
	return false
}
