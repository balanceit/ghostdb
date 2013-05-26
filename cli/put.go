// Package cli is the command line interface for the ghostdb project
package cli

import (
	"github.com/balanceit/ghostdb/storage"
	"github.com/balanceit/ghostdb/utils"
	"strconv"
	"unicode/utf8"
)

//would want this to be a []byte or even a some sort of binary stream
func Put(data string, filename string) (string, error) {

	var uuid, _ = utils.GenerateV4String()
	var len = strconv.Itoa(utf8.RuneCountInString(data))
	var writeData = uuid + len + data
	storage.Persist(writeData, filename)

	return uuid, nil
}
