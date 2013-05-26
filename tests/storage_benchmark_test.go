package tests

import (
	"github.com/balanceit/ghostdb/storage"
	"os"
	"testing"
)

func Benchmark_Persist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var filename = "testfile"
		var randomString = "random string"
		storage.Persist(randomString, filename)
		os.Remove(filename)
	}
}
