package tests

import (
	"github.com/balanceit/ghostdb/cli"
	"github.com/balanceit/ghostdb/utils"
	"testing"
)

func Benchmark_Get(b *testing.B) {
	conn, _ := cli.CreateConnction("testfile")
	for i := 0; i < b.N; i++ {
		conn.Get("testuuid")
	}
}

func Benchmark_Put(b *testing.B) {
	conn, _ := cli.CreateConnction("testfile")
	for i := 0; i < b.N; i++ {
		conn.Put("tests")
	}
}

func Benchmark_Post(b *testing.B) {
	var uuid, _ = utils.GenerateV4String()
	conn, _ := cli.CreateConnction("testfile")
	for i := 0; i < b.N; i++ {
		conn.Post(uuid, "tests")
	}
}

func Benchmark_Delete(b *testing.B) {
	conn, _ := cli.CreateConnction("testfile")
	for i := 0; i < b.N; i++ {
		conn.Delete("testuuid")
	}
}
