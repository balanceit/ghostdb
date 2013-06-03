package cli

import (
	"github.com/balanceit/ghostdb/storage"
	"github.com/balanceit/ghostdb/utils"
)

type Connection struct {
	filename string
}

func CreateConnction(filename string) (Connection, error) {
	return Connection{filename}, nil
}

//would want this to be a []byte or even a some sort of binary stream
func (c Connection) Put(data string) (string, error) {

	var uuid, _ = utils.GenerateV4String()
	err := storage.Store(uuid, data, c.filename)

	return uuid, err
}

func (c Connection) Post(uuid string, data string) error {

	storage.Store(uuid, data, c.filename)
	return nil
}

func (c Connection) Get(uuid string) (string, error) {
	return "data", nil
}

func (c Connection) Delete(uuid string) error {
	return nil
}
