/*
DESAFIO 3: Importações Múltiplas e Aliases
Objetivo: Aprender a organizar importações e criar apelidos (aliases) para pacotes.

Requisitos:
1. Importe o pacote "fmt" e crie um alias para ele chamado "f" (ex: import f "fmt").
2. Importe o pacote "strings" normalmente.
3. Crie uma variável com uma frase em letras minúsculas.
4. Use a função pública (exportada) ToUpper do pacote "strings" para transformar a frase em maiúsculas.
5. Imprima o resultado usando o seu alias "f" no lugar de "fmt".
*/
package main

import (
	f "fmt"
	"strings"
)

func main() {
	// Sua implementação começa aqui
	f.Println("Iniciando Desafio 3...")
	texto := "todas as palavras em letra minúsculas"
	f.Println(strings.ToUpper(texto))
}
