package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const suffix = "json"

var GlobalStorage = &Storage{}

type Storage struct {
	dir string
}

// Initialize .
func (o *Storage) Initialize(dir string) error {
	o.dir = dir
	if _, err := os.Stat(o.dir); err == nil {
		log.Println(fmt.Sprintf(" Using %s alread exists", o.dir))
		return nil
	}
	return os.MkdirAll(o.dir, os.ModePerm)
}

// Write v to collection
func (o *Storage) Write(collection string, key string, value interface{}) error {
	if len(collection) <= 0 {
		return errors.New("storage write, no collection")
	}
	if len(key) <= 0 {
		return errors.New("storage write, no key")
	}
	if err := os.MkdirAll(path.Join(o.dir, collection), os.ModePerm); err != nil {
		return err
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	dest := path.Join(o.dir, collection, fmt.Sprintf("%s.%s", key, suffix))
	destTmp := fmt.Sprintf("%s.tmp", dest)
	if err := ioutil.WriteFile(destTmp, bytes, 0644); err != nil {
		return err
	}
	return os.Rename(destTmp, dest)
}

// Read record from collection
func (o *Storage) Read(collection string, key string, value interface{}) error {
	if len(collection) <= 0 {
		return errors.New("storage read, no collection")
	}
	if len(key) <= 0 {
		return errors.New("storage read, no key")
	}
	dest := path.Join(o.dir, collection, fmt.Sprintf("%s.%s", key, suffix))
	if _, err := os.Stat(dest); err != nil {
		return nil
	}
	record, err := ioutil.ReadFile(dest)
	if err != nil {
		return err
	}
	return json.Unmarshal(record, &value)
}

// Delete key from collection
func (o *Storage) Delete(collection string, key string) error {
	if len(collection) <= 0 {
		return errors.New("storage delete, no collection")
	}
	if len(key) <= 0 {
		return errors.New("storage delete, no key")
	}
	dest := path.Join(o.dir, collection, fmt.Sprintf("%s.%s", key, suffix))
	if _, err := os.Stat(dest); err != nil {
		return err
	}
	return os.RemoveAll(dest)
}

// ReadAll .
func (o *Storage) ReadAll(collection string, records *[][]byte) error {
	if len(collection) <= 0 {
		return errors.New("storage read all, no collection")
	}

	files, err := ioutil.ReadDir(path.Join(o.dir, collection))
	if err != nil {
		log.Printf("storage read all error (is empty): %s \n", err)
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		bytes, err := ioutil.ReadFile(path.Join(o.dir, collection, file.Name()))
		if err != nil {
			return err
		}
		*records = append(*records, bytes)
	}

	return nil
}

// RecordsToStruct is convert records to []T
func RecordsToStruct[T any](records [][]byte, vs *[]T) error {
	for _, bts := range records {
		var record T
		err := json.Unmarshal(bts, &record)
		if err != nil {
			return err
		}
		*vs = append(*vs, record)
	}
	return nil
}
