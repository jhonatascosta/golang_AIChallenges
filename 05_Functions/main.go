/*
DESAFIO 3: Criando Funções
Objetivo: Isolar a lógica em blocos reutilizáveis.

Requisitos:
1. Crie uma função chamada 'somar' FORA da função main.
2. A função 'somar' deve receber dois parâmetros do tipo 'int' (ex: a, b) e retornar um 'int'.
3. Dentro da função 'somar', retorne a soma dos dois parâmetros.
4. Dentro da função main, chame a função 'somar' passando dois números, guarde o resultado em uma variável e imprima.
*/
package main

import "fmt"

// Declare sua função somar aqui
func somar(a, b int) int {
	return a + b
}

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 3: Funções Simples...")
	resultado := somar(10,20)
	fmt.Println("Resultado:", resultado)
}
