/*
DESAFIO 1: If com Statement Curto (Short Statement)
Objetivo: Usar a sintaxe idiomática do Go para inicializar e testar uma variável na mesma linha do 'if'.

Requisitos:
1. Crie um bloco 'if' que declare uma variável 'numero' com o valor 42 usando ':=', e logo após o ponto e vírgula (;), verifique se 'numero' é par (numero % 2 == 0).
2. Se for par, imprima: "[numero] é PAR".
3. Se não for (else), imprima: "[numero] é ÍMPAR".
4. Tente acessar a variável 'numero' fora do bloco if/else usando um fmt.Println. O que acontece? (Responda em formato de comentário no final do código).
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 1: If com short statement...")
	if numero := 42; numero % 2 == 0 {
		fmt.Printf("%d é PAR\n", numero)
	} else {
		fmt.Printf("%d é ÍMPAR\n", numero)
	}
	//fmt.Println(numero) Não é possivel acessar o valor de numero fora do escopo do IF.
}
