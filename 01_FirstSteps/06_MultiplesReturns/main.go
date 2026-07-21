/*
DESAFIO 4: Múltiplos Retornos
Objetivo: Utilizar uma das minhas features mais queridas: retornar mais de um valor!

Requisitos:
1. Crie uma função chamada 'dividir' que recebe dois 'float64' (numerador e denominador).
2. A função deve retornar dois valores: o resultado da divisão (float64) e um booleano (bool) indicando se a divisão foi válida.
3. Regra: Se o denominador for 0, retorne 0.0 e false (pois não podemos dividir por zero). Caso contrário, retorne o resultado e true.
4. Na main, chame a função, receba os dois retornos em variáveis e imprima-os.
*/
package main

import "fmt"

// Declare sua função dividir aqui
func dividir(numerador, denominador float64) (float64,bool) {
	divisao := numerador/denominador
	if denominador == 0 {
		return divisao, false
	} else {
		return divisao, true
	}
}

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 4: Múltiplos Retornos...")
	resultado, valido := dividir(10.0, 0.0)
	fmt.Printf("A divisão é válida? %t \nO resultado da divisão é: %.5f\n", valido, resultado)
}
