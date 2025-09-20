package main

import (
	"strings"
	"testing"
)

const N = 100000

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < N; j++ {
			s += "x"
		}
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for j := 0; j < N; j++ {
			sb.WriteString("x")
		}
		_ = sb.String()
	}
}

func BenchmarkBuilderGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.Grow(N) // pré-aloca espaço para N bytes
		for j := 0; j < N; j++ {
			sb.WriteByte('x')
		}
		_ = sb.String()
	}
}
