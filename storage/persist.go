package storage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

func Persist(in string, filename string) error {
	//this needs to be opened for append only
	fo, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = io.WriteString(fo, in)

	if err != nil {
		panic(err)
	}

	return nil
}

func Read() string {
	type RedIn struct {
		entries []string
	}

	// open input file
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(fi)
	s, err := r.ReadString(byte('|'))
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Printf("found a = %v\n", s)

	return s
}

func Index(uuid string, location int, filename string) error {
	var locationString = strconv.Itoa(location)

	err := Persist(uuid+locationString, filename)

	return err
}

func Store(uuid string, data string, filename string) error {
	var len = strconv.Itoa(utf8.RuneCountInString(data))

	err := Persist(uuid+len+data, filename)
	return err
}
