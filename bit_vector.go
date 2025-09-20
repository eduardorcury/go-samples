package main

import (
	"fmt"
	"strings"
)

// Um bit vector, bit map ou bitset é uma estrutura de dados que armazena valores 0 ou 1
// Serve o mesmo propósito que um mapa de booleanos indicando quais índices existem ou não
// No entanto, ocupa menos memória que um mapa de booleanos (1 bit vs 8 bits)
// Num vetor de bits, um elemento igual a 1 (ligado) na posição 3 significa que o número 3 existe

type IntSet struct {
	//se o número 10 está no conjunto, então o bit 10 (dentro do words[0]) vai estar ligado.
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Nota: o uso de ponteiro é obrigatório pois sem ele o append só alteraria uma copia de words, e não
// a estrutura original.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var sb strings.Builder
	sb.WriteByte('{')
	first := true

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if !first {
					sb.WriteByte(' ')
				}
				first = false
				sb.WriteString(fmt.Sprintf("%d", 64*i+j))
			}
		}
	}

	sb.WriteByte('}')
	return sb.String()
}
