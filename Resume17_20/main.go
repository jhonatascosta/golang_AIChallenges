/*
DESAFIO 5: O Combo de Memória (Boss Battle)
Objetivo: Combinar funções, ponteiros e slices para entender como os dados trafegam (Stack vs Heap).

Requisitos:
1. Crie uma função chamada 'dobrarValor' que recebe um PONTEIRO para um int (*int). Dentro dela, multiplique o valor por 2.
2. Crie uma função chamada 'alterarPrimeiroElemento' que recebe um SLICE de ints ([]int). Dentro dela, mude o valor do índice 0 para 999.
3. Na main, declare uma variável int com valor 10. Chame 'dobrarValor' passando o endereço dessa variável.
4. Na main, declare um slice de ints com os valores {1, 2, 3}. Chame 'alterarPrimeiroElemento' passando o slice.
5. Imprima a variável int e o slice após as funções serem chamadas. Ambos devem estar modificados!
*/
package main

import "fmt"

// Declare suas funções aqui
func dobrarValor(pointer *int) {
	*pointer = *pointer * 2
}

func alterarPrimeiroElemento(slice []int) {
	slice[0] = 999
}

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 5: Boss Battle de Memória...")
	valor := 10
	dobrarValor(&valor)
	slice1 := []int{1,2,3}
	alterarPrimeiroElemento(slice1)
	fmt.Println(valor, slice1)
}
