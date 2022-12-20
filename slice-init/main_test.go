package main

import "testing"

const n = 1_000_000

var global []Doller

func BenchmarkConvert_EmptySlice(b *testing.B) {
	var local []Doller
	euros := make([]Euro, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertEmptySlice(euros)
	}
	global = local
}

func BenchmarkConvert_GivenCapacity(b *testing.B) {
	var local []Doller
	euros := make([]Euro, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertGivenCapacity(euros)
	}
	global = local
}

func BenchmarkConvert_GivenLength(b *testing.B) {
	var local []Doller
	euros := make([]Euro, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertGivenLength(euros)
	}
	global = local
}
