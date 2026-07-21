/*
DESAFIO 5: Conhecendo os Arrays
Objetivo: Declarar e manipular uma lista de tamanho fixo.

Requisitos:
1. Declare um array de inteiros com tamanho 5 chamado 'notas'. (Ex: var notas [5]int).
2. Atribua valores para os índices 0, 1 e 4 do array. (Deixe o resto com o valor zero (zero value) padrão do Go).
3. Imprima o array completo.
4. Imprima apenas o valor do índice 1.
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 5: Arrays...")
	notas := [5]int{0: 10, 1: 20, 4:30}
	fmt.Println(notas)
	fmt.Println(notas[1])
}
