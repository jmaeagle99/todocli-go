package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Store[T any] struct {
	nextId   int
	rootPath string
}

type StoreIndex struct {
	NextId int `json:"nextId"`
}

// Creates a data store backed by files on disk
func CreateStore[T any]() (*Store[T], error) {
	userPath, err := os.UserHomeDir()
	if nil != err {
		return nil, err
	}

	todopath := filepath.Join(userPath, ".todocli")
	err = os.MkdirAll(todopath, 0644)
	if nil != err {
		return nil, err
	}

	var nextId = 1
	indexpath := filepath.Join(todopath, "index.json")
	if _, err = os.Stat(indexpath); nil == err {
		indexFile, err := os.OpenFile(indexpath, os.O_RDONLY, 0644)
		if nil != err {
			return nil, err
		}
		defer indexFile.Close()

		var index StoreIndex
		decoder := json.NewDecoder(indexFile)
		err = decoder.Decode(&index)
		if nil != err {
			return nil, err
		}
		nextId = index.NextId
	}

	store := &Store[T]{
		rootPath: todopath,
		nextId:   nextId,
	}

	return store, nil
}

func (store *Store[T]) writeIndex() error {
	indexpath := filepath.Join(store.rootPath, "index.json")
	indexFile, err := os.OpenFile(indexpath, os.O_CREATE|os.O_WRONLY, 0644)
	if nil != err {
		return err
	}
	defer indexFile.Close()

	index := StoreIndex{NextId: store.nextId}

	encoder := json.NewEncoder(indexFile)
	return encoder.Encode(index)
}

func (store *Store[T]) AddItem(item T) error {
	fileName := fmt.Sprintf("%d.json", store.nextId)
	filePath := filepath.Join(store.rootPath, fileName)

	store.nextId++
	err := store.writeIndex()
	if nil != err {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if nil != err {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(item)
}
