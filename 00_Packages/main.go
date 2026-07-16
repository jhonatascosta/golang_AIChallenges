/*
DESAFIO 1: Usando Pacotes Nativos
Objetivo: Importar pacotes existentes e utilizar seus nomes exportados (públicos).

Requisitos:
1. Importe os pacotes "fmt" e "math".
2. Tente imprimir o valor de Pi fornecido pelo pacote math. 
   (Lembre-se da regra: para acessar algo de outro pacote, a primeira letra deve ser maiúscula. É math.Pi ou math.pi?)
3. Imprima também o resultado de uma operação matemática usando a função pública de raiz quadrada do pacote math (Sqrt).
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// Sua implementação começa aqui
	fmt.Println("Iniciando Desafio 1: Pacotes e Nomes Exportados...")
	fmt.Println("Imprimindo o valor de pi: ", math.Pi)
	fmt.Println("Imprimindo o valor de uma raiz quadrada: ", math.Sqrt(64))
}
