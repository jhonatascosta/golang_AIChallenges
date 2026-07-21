/*
DESAFIO 5: A Pilha do Defer (LIFO)
Objetivo: Compreender a ordem de execução de múltiplos defers através de um loop.

Requisitos:
1. Crie um loop 'for' clássico onde o 'i' vai de 1 até 3 (inclusive).
2. Dentro do loop, coloque um defer imprimindo o valor de 'i'. Ex: defer fmt.Println("Defer número:", i)
3. Fora do loop, na última linha do main, imprima: "Fim da função main".
4. Responda em um comentário: Qual foi a ordem dos números impressos e por que isso acontece?
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 5: Empilhando Defers...")
	for i := 1; i <= 3; i++ {
		defer fmt.Println("Defer número:", i)
	}
	fmt.Println("Fim da função main")
	// A execução ficou primeiro o Print fora do loop e depois a sequencia em ordem decrescente do loop de 3,2,1. Saiu assim pois são executadas em pilha,
	// onde o ultimo a entrar, é o primeiro a sair, pois ele só e executado após o fim do escopo da função onde o defer está alinhado.
}
