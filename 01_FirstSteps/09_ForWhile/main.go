/*
DESAFIO 7: For com Apenas a Condição
Objetivo: Usar o 'for' como se fosse o 'while' de outras linguagens.

Requisitos:
1. Declare uma variável 'contador' com o valor 1 antes do loop.
2. Crie um loop 'for' que verifique apenas a condição: 'contador <= 5'.
3. Dentro do loop, imprima o valor do 'contador'.
4. Importante: Não esqueça de incrementar o 'contador' dentro do loop (contador++), senão você criará um loop infinito!
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 7: For como While...")
	var contador int = 1
	for contador <= 5 {
		fmt.Println(contador)
		contador++
	}
}
