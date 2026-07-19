/*
DESAFIO 3: Capacidade Dinâmica
Objetivo: Usar a função make para alocar fatias (slices) na memória Heap e observar como o Go aumenta a capacidade magicamente.

Requisitos:
1. Use a função 'make' para criar um slice de strings chamado 'frutas' com tamanho (len) 2 e capacidade (cap) 4.
2. Imprima o slice, o 'len' e o 'cap' atuais (use as funções nativas len() e cap()).
3. Atribua "Maçã" e "Banana" aos índices 0 e 1.
4. Agora, use a função 'append' para adicionar "Laranja", "Uva" e "Manga" ao slice.
5. Imprima novamente o slice, o 'len' e o 'cap'. 
6. Responda num comentário: O que aconteceu com a capacidade (cap) quando você ultrapassou o limite inicial de 4?
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 3: Make, Len e Cap...")
	frutas := make([]string, 2, 4)
	fmt.Println(frutas, len(frutas), cap(frutas))
	frutas[0] = "Maçã"
	frutas[1]	=	"Banana"
	frutas = append(frutas, "Laranja", "Uva", "Manga")
	fmt.Println(frutas, len(frutas), cap(frutas))
	// A capacidade tem por padrão a funcionalidade de dobrar para que o slice continue adicionando mais elementos.
}
