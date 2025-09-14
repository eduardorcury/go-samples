package main

import "testing"

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindromeFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeFaster(`A man, a plan, a canal: Panama`)
	}
}
