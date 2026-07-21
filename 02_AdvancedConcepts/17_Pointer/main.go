/*
DESAFIO 1: Apontando para a Memória
Objetivo: Entender como criar e desreferenciar ponteiros.

Requisitos:
1. Declare uma variável comum chamada 'valor' do tipo int com o valor 100.
2. Declare um ponteiro chamado 'ponteiroValor' que aponte para a variável 'valor' (use o operador &).
3. Imprima o valor da variável 'valor' e o endereço de memória armazenado em 'ponteiroValor'.
4. Usando apenas o 'ponteiroValor' (com o operador de desreferência *), mude o valor para 200.
5. Imprima a variável 'valor' novamente para provar que ela foi alterada pelo ponteiro!
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 1: Ponteiros...")
	valor := 100
	ponteiroValor := &valor
	fmt.Println(valor, ponteiroValor)
	*ponteiroValor = 200
	fmt.Println(valor)
}
