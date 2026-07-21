/*
DESAFIO 2: Declarando Nomes Públicos e Privados
Objetivo: Compreender a regra de capitalização do Go para visibilidade.

Requisitos:
1. Fora da função main (no escopo do pacote), declare uma variável "privada" (não-exportada) chamada 'versao' com o valor "1.0.0".
2. Ainda no escopo do pacote, declare uma variável "pública" (exportada) chamada 'NomeAplicacao' com o valor "Meu Primeiro App Go".
3. Dentro da função main, imprima as duas variáveis usando o pacote "fmt".
4. Responda em formato de comentário no código: Se outro pacote tentasse importar essas duas variáveis, qual delas ele conseguiria acessar e por quê?
*/
package main

import "fmt"
var versao string = "1.0.0"
var NomeAplicacao string = "Meu Primeiro App Go"
func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 2: Público vs Privado...")
	fmt.Println(versao)
	fmt.Println(NomeAplicacao)
	// Apenas a variável com o nome de "NomeAplicacao" seria importada, pois ela está como pública. A primeira letra do nome do objeto define se ela é pública ou privada. Letra maiuscula é pública e letra minuscula é privada. 
}
