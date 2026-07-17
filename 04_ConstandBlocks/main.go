/*
DESAFIO 2: Trabalhando com Constantes
Objetivo: Entender como valores imutáveis funcionam em Go.

Requisitos:
1. Declare uma constante chamada 'Pi' com o valor 3.14159.
2. Crie um bloco de constantes (usando apenas um 'const' e parênteses) declarando:
   - 'DiasNaSemana' igual a 7
   - 'HorasNoDia' igual a 24
3. Tente reatribuir um novo valor para 'Pi' dentro do main (ex: Pi = 3.14). 
   Veja o erro que o compilador do Go (através do gopls no NVChad) vai te dar, depois comente essa linha errada e explique o erro em um comentário.
4. Imprima as constantes validas.
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 2: Constantes...")
	const Pi float64 = 3.14159
	const (
		DiasNaSemana int = 7
		HorasNoDia int = 24
	)
	// Pi = 3.14 O valor de uma constante é inalteravel (com excessões). Mesmo tentando atribuir um novo valor, a linguagem não faz a alteração e aponta o erro.
	fmt.Printf("Constantes válidas:\nPi: %.5f\nDiasNaSemana: %d\nHorasNoDia: %d\n",Pi,DiasNaSemana,HorasNoDia)
}
