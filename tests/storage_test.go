package tests

import (
	"github.com/balanceit/ghostdb/storage"
	"github.com/balanceit/ghostdb/utils"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"unicode/utf8"
)

func Test_Persist(t *testing.T) {
	filename := "testfile2"
	fo, _ := os.Create(filename)

	fo.Close()
	var data = "something stored"
	var uuid, _ = utils.GenerateV4String()
	var len = strconv.Itoa(utf8.RuneCountInString(data))
	var expected = uuid + len + data

	storage.Persist(expected, filename)
	buf, _ := ioutil.ReadFile(filename)
	s := string(buf)

	if s != expected {
		t.Errorf("expecting [%s] but got [%s]", s, expected)
	}

	defer func() {
		os.Remove(filename)
	}()
}
