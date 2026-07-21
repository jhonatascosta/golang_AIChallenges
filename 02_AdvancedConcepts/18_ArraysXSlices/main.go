/*
DESAFIO 2: Comportamento de Cópia
Objetivo: Provar que Arrays copiam valores (passagem por valor), enquanto Slices apontam para a mesma memória (referência ao array subjacente).

Requisitos:
1. Crie um ARRAY de 3 inteiros chamado 'meuArray' e atribua ele a uma nova variável 'copiaArray'.
2. Mude o valor do índice 0 de 'copiaArray' para 99.
3. Crie um SLICE de inteiros chamado 'meuSlice' (sem tamanho fixo) e atribua ele a uma nova variável 'copiaSlice'.
4. Mude o valor do índice 0 de 'copiaSlice' para 99.
5. Imprima os 4 (meuArray, copiaArray, meuSlice, copiaSlice).
6. Responda num comentário: Por que os arrays ficaram diferentes, mas os slices ficaram iguais?
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 2: Arrays vs Slices...")
	meuArray := [3]int{18, 20, 33}
	copiaArray := meuArray
	copiaArray[0] = 99
	meuSlice := []int{10, 17, 29, 103}
	copiaSlice := meuSlice
	copiaSlice[0] = 99
	fmt.Println(meuArray, copiaArray, meuSlice, copiaSlice)
	/*
		No caso do Array, ele aponta os valores em si, por isso ele primeiro mostra os valores da variavel meuArray e depois mostra os valores de 
		copiaArray, pois são arrays com valores diferentes. Já os slices eles apontam os ponteiros, assim quando o valor do indice 0 e alterado, 
		o valor na memoria também é alterado, assim quando é impresso ele pega o valor que está na memoria.
	*/
}
