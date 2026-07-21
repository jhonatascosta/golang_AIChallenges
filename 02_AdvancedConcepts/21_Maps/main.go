/*
DESAFIO 1: Manipulando Maps
Objetivo: Criar um dicionário e usar o idioma "comma ok" para evitar erros.

Requisitos:
1. Usa a função 'make' para criar um map chamado 'dicionario' onde as chaves são strings e os valores também são strings (map[string]string).
2. Adiciona 3 palavras ao teu dicionário (ex: "apple": "maçã", "car": "carro", "book": "livro").
3. Elimina uma das palavras usando a função nativa 'delete(dicionario, "chave")'.
4. Tenta procurar uma palavra que NÃO existe no map. Usa a sintaxe: valor, ok := dicionario["chave_inexistente"].
5. Usa um 'if' para verificar a variável 'ok'. Se for falso, imprime "Palavra não encontrada". Se for verdadeiro, imprime a tradução.
*/
package main

import "fmt"

func main() {
	// A tua implementação começa aqui
	fmt.Println("A iniciar Desafio 1: Maps...")
	dicionario := make(map[string]string)
	dicionario["apple"] = "maçã"
	dicionario["car"] = "carro"
	dicionario["book"] = "livro"
	delete(dicionario, "car")
	valor, ok := dicionario["house"]
	if ok == false {
		fmt.Println("Palavra não encontrada")
	} else {
		fmt.Println(valor)
	}
}
