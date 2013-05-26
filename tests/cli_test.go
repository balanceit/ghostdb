package tests

import (
	"github.com/balanceit/ghostdb/cli"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"unicode/utf8"
)

func Test_Get(t *testing.T) {
	cli.Get()
}

func Test_Put(t *testing.T) {
	var filename = "testfile"
	var fo, _ = os.Create(filename)
	fo.Close()

	var data = "something stored"
	var uuid, _ = cli.Put(data, filename)

	var len = strconv.Itoa(utf8.RuneCountInString(data))
	var expected = uuid + len + data

	buf, _ := ioutil.ReadFile(filename)
	s := string(buf)
	if s != expected {
		t.Errorf("expecting [%s] but got [%s]", s, expected)
	}

	defer func() {
		os.Remove(filename)
	}()
}

func Test_Post(t *testing.T) {
	cli.Post()
}

func Test_Delete(t *testing.T) {
	cli.Delete()
}
