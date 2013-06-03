package tests

import (
	"github.com/balanceit/ghostdb/cli"
	"github.com/balanceit/ghostdb/utils"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"unicode/utf8"
)

func setup(filename string) {
	var fo, _ = os.Create(filename)
	fo.Close()

}

func Test_Get(t *testing.T) {
	var filename = "testfile"
	setup(filename)
	defer func() {
		os.Remove(filename)
	}()

	var expected = "data"
	conn, _ := cli.CreateConnction(filename)
	var ret, _ = conn.Get("testuuid")

	if ret != expected {
		t.Errorf("expecting [%s] but got [%s]", ret, expected)
	}
}

func Test_Put(t *testing.T) {
	var filename = "testfile"
	setup(filename)
	defer func() {
		os.Remove(filename)
	}()

	var data = "something stored"

	conn, _ := cli.CreateConnction(filename)
	var uuid, _ = conn.Put(data)

	var len = strconv.Itoa(utf8.RuneCountInString(data))
	var expected = uuid + len + data

	var buf, _ = ioutil.ReadFile(filename)
	s := string(buf)
	if s != expected {
		t.Errorf("expecting [%s] but got [%s]", s, expected)
	}
}

func Test_Post(t *testing.T) {
	var filename = "testfile"
	setup(filename)
	defer func() {
		os.Remove(filename)
	}()

	var data = "something stored"
	var uuid, _ = utils.GenerateV4String()

	conn, _ := cli.CreateConnction(filename)
	var _ = conn.Post(uuid, data)

	var len = strconv.Itoa(utf8.RuneCountInString(data))
	var expected = uuid + len + data

	var buf, _ = ioutil.ReadFile(filename)
	s := string(buf)
	if s != expected {
		t.Errorf("expecting [%s] but got [%s]", s, expected)
	}
}

func Test_Delete(t *testing.T) {
	var filename = "testfile"
	setup(filename)
	defer func() {
		os.Remove(filename)
	}()

	conn, _ := cli.CreateConnction(filename)
	_ = conn.Delete("testuuid")

}
