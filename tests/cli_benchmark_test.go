package tests

import (
	"github.com/balanceit/ghostdb/cli"
	"testing"
)

func Benchmark_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cli.Get()
	}
}

func Benchmark_Put(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cli.Put("tests", "testfile")
	}
}

func Benchmark_Post(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cli.Post()
	}
}

func Benchmark_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cli.Delete()
	}
}
