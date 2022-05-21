package api

import (
	"fmt"
	"strconv"
	"time"
)

const (
	luaStorageCollection = "luas"
)

type Lua struct {
	ConnectionID    string        `json:"connectionID"`
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Keys            []string      `json:"keys"`
	Args            []interface{} `json:"args"`
	Script          string        `json:"script"`
	LastExecutionAt int64         `json:"lastExecutionAt"`
	Elapsed         string        `json:"elapsed"`
	Result          interface{}   `json:"result"`
}

// New add lua script
func (l *Lua) New() error {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	l.ID = id
	return GlobalStorage.Write(luaStorageCollection, fmt.Sprintf("%s_%s", l.ConnectionID, l.ID), l)
}

// Edit .
func (l *Lua) Edit() error {
	return GlobalStorage.Write(luaStorageCollection, fmt.Sprintf("%s_%s", l.ConnectionID, l.ID), l)
}

// Delete .
func (l *Lua) Delete() error {
	return GlobalStorage.Delete(luaStorageCollection, fmt.Sprintf("%s_%s", l.ConnectionID, l.ID))
}

// LoadLuas load all lua scripts
func LoadLuas(luas *[]Lua) error {
	var bytes [][]byte
	if err := GlobalStorage.ReadAll(luaStorageCollection, &bytes); err != nil {
		return err
	}

	if err := RecordsToStruct(bytes, luas); err != nil {
		return err
	}

	return nil
}
