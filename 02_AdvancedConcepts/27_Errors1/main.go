/*
DESAFIO 1: Tratamento de Erros Básico
Objetivo: Criar uma função que retorna um valor e um erro, e tratá-lo no main.

Instruções:
1. Crie uma função chamada 'Dividir(a, b float64) (float64, error)'.
2. Se 'b' for igual a zero, retorne 0 e um erro usando 'errors.New("não é possível dividir por zero")'.
3. Caso contrário, retorne o resultado da divisão e 'nil' para o erro.
4. No main, chame a função e use o clássico 'if err != nil' para exibir o erro ou o resultado.
*/
package main

import (
	"errors"
	"fmt"
	"math"
	// Lembre-se de importar o pacote "errors" se necessário!
)
func dividir(a,b float64) (float64, error) {
		if b == 0 {
			return math.NaN(),errors.New("não é possível dividir por zero")
		}
		return a / b, nil
}
func main() {
	fmt.Println("Iniciando Desafio 1: Tratamento de Erros...")
	// Seu código de teste aqui
	
	resultado, err := dividir(7, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("O resultado da divisão:", resultado)
}
