/*
DESAFIO 6: O For em Três Partes
Objetivo: Escrever um loop com inicialização, condição e pós-execução.

Requisitos:
1. Crie um loop 'for' que comece com uma variável 'i' valendo 1.
2. O loop deve continuar enquanto 'i' for menor ou igual a 10.
3. A cada iteração, incremente 'i' em 1 (i++).
4. Dentro do loop, imprima o valor de 'i'.
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 6: For Clássico...")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
