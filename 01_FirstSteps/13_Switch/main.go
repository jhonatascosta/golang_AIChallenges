/*
DESAFIO 2: Switch com Múltiplos Valores
Objetivo: Usar um switch para verificar uma variável contra múltiplas opções no mesmo 'case'.

Requisitos:
1. Declare uma variável 'dia' do tipo int com um valor de 1 a 7 (representando os dias da semana, onde 1 = Domingo, 7 = Sábado).
2. Crie um 'switch' avaliando a variável 'dia'.
3. Em um único 'case', verifique se é 1 ou 7. Se for, imprima "Fim de semana!".
4. Em outro 'case', verifique os dias de 2 a 6. Se for, imprima "Dia útil!".
5. Adicione um 'default' para imprimir "Dia inválido" caso o número não esteja entre 1 e 7.
*/
package main

import "fmt"

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 2: Switch Clássico...")
	dia := 2
	switch dia {
		case 1, 7:
			fmt.Println("Fim de Semana!")
		case 2,3,4,5,6:
			fmt.Println("Dia útil!")
		default: 
			fmt.Println("Dia inválido")
	}
}
