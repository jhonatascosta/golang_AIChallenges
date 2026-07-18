/*
DESAFIO 4: O Básico do Defer
Objetivo: Entender como o 'defer' empurra a execução para o final da função.

Requisitos:
1. Dentro do main, digite: defer fmt.Println("1 - Execução finalizada (fechando recursos...)")
2. Na linha de baixo, digite: fmt.Println("2 - Processando dados importantes...")
3. Na linha de baixo, digite: fmt.Println("3 - Conectando ao banco de dados...")
4. Rode o código e observe a ordem em que as frases são impressas no terminal. A frase do defer deve aparecer por último!
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 4: O Básico do Defer...")
	defer fmt.Println("1 - Execução finalizada (fechando recursos...)")
	fmt.Println("2 - Processando dados importantes...")
	fmt.Println("3 - Conectando ao banco de dados...")
}
