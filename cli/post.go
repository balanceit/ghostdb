// Package cli is the command line interface for the ghostdb project
package cli

import (
	"github.com/balanceit/ghostdb/storage"
)

func Posts(uuid string, data string) error {

	storage.Store(uuid, data, "ff")
	return nil
}
