/*
DESAFIO 1: Variáveis e Tipos Básicos
Objetivo: Praticar as diferentes formas de declarar variáveis.

Requisitos:
1. Declare uma variável 'idade' do tipo int usando a palavra-chave 'var'.
2. Declare uma variável 'peso' do tipo float64 usando a palavra-chave 'var'.
3. Use o operador de declaração curta (:=) para criar uma variável 'nome' com uma string.
4. Use o operador (:=) para criar uma variável 'estudando' com o valor booleano true.
5. Imprima todas as variáveis usando fmt.Println.
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 1: Variáveis e Tipos...")
	var idade int = 27
	var peso float64 = 99.1
	nome := "Jhonata"
	estudando := true
	fmt.Println("Nome: ", nome, "Idade: ", idade, "Peso: ", peso, "Estudando: ", estudando)
}
