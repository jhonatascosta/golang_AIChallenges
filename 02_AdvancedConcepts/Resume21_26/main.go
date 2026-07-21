/*
DESAFIO 7: Boss Battle (Cache Genérica)
Objetivo: Juntar Maps, Structs, Methods, Interfaces e Type Assertions.

Requisitos:
1. Cria uma interface 'Cache' com os métodos: 
   - Guardar(chave string, valor any)
   - Obter(chave string) any
2. Cria uma struct 'MemoriaCache' que contenha um campo 'dados' do tipo map[string]any.
3. Cria um construtor (uma função NovaMemoriaCache() *MemoriaCache) que inicialize o map com 'make'.
4. Implementa os métodos da interface 'Cache' para a struct 'MemoriaCache' (usa ponteiros nos receivers para poderes alterar o map!).
5. Na main, guarda o teu nome (string) e a tua idade (int) na cache.
6. Recupera a tua idade da cache usando o método Obter. Como ela retorna 'any', usa uma Type Assertion direta (valor, ok := idade.(int)) para somar +10 à tua idade e imprimir o resultado!
*/
package main

import "fmt"

// Implementa as tuas interfaces, structs e métodos aqui
type Cache interface{
	Guardar(chave string, valor any)
	Obter(chave string) any
}

type MemoriaCache struct{
	dados map[string]any
}

func NovaMemoriaCache() *MemoriaCache {
	return &MemoriaCache{
		dados: make(map[string]any),
	}
}

func (m *MemoriaCache) Guardar(chave string, valor any) {
	m.dados[chave] = valor
}

func (m *MemoriaCache) Obter(chave string) any {
	return m.dados[chave]
}

func main() {
	// A tua implementação começa aqui
	fmt.Println("A iniciar Desafio 7: Boss Battle...")
	var cache Cache = NovaMemoriaCache()
	cache.Guardar("nome", "Jhonata")
	cache.Guardar("idade", 25)

	if nomeRetornado, ok := cache.Obter("nome").(string); ok {
		fmt.Printf("Dado recuperado com sucesso [string]: %s\n", nomeRetornado)
	}

	if idadeRetornada, ok := cache.Obter("idade").(int); ok {
		novaIdade := idadeRetornada + 10
		fmt.Printf("Dado recuperado [int]: %d | Somando +10 anos na cache: %d\n", idadeRetornada, novaIdade)
	}
}
