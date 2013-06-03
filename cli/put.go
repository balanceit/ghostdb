// Package cli is the command line interface for the ghostdb project
package cli

import (
	"github.com/balanceit/ghostdb/storage"
	"github.com/balanceit/ghostdb/utils"
)

//would want this to be a []byte or even a some sort of binary stream
func Puts(data string, filename string) (string, error) {

	var uuid, _ = utils.GenerateV4String()
	err := storage.Store(uuid, data, filename)

	return uuid, err
}
