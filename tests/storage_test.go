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
	var filename = "testfile"
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

func Test_Index(t *testing.T) {
	var indexfilename = "testindex"

	var fo, _ = os.Create(indexfilename)
	fo.Close()

	var uuid, _ = utils.GenerateV4String()
	var location = 0
	storage.Index(uuid, location, indexfilename)
	var expected = uuid + strconv.Itoa(location)

	buf, _ := ioutil.ReadFile(indexfilename)
	s := string(buf)

	if s != expected {
		t.Errorf("expecting [%s] but got [%s]", s, expected)
	}

	defer func() {
		os.Remove(indexfilename)
	}()
}
