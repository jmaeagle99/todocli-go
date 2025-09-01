package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Store[T any] struct {
	rootPath        string
	requestChannel  chan int32
	responseChannel chan int32
}

type StoreIndex struct {
	LastId int32 `json:"lastId"`
}

// Creates a data store backed by files on disk
func CreateStore[T any]() (*Store[T], error) {
	userPath, err := os.UserHomeDir()
	if nil != err {
		return nil, err
	}

	todopath := filepath.Join(userPath, ".todocli")
	err = os.MkdirAll(todopath, 0755)
	if nil != err {
		return nil, err
	}

	store := &Store[T]{
		rootPath:        todopath,
		requestChannel:  make(chan int32, 5),
		responseChannel: make(chan int32, 5),
	}

	go store.indexWorker()

	return store, nil
}

func (store *Store[T]) indexWorker() {
	defer close(store.responseChannel)

	indexpath := filepath.Join(store.rootPath, "index.json")
	indexFile, err := os.OpenFile(indexpath, os.O_CREATE|os.O_RDWR, 0644)
	if nil != err {
		return
	}
	defer func() {
		_ = indexFile.Close()
	}()

	var index StoreIndex
	decoder := json.NewDecoder(indexFile)
	err = decoder.Decode(&index)

	var lastId int32 = 0
	if nil == err {
		lastId = index.LastId
	}

	for range store.requestChannel {
		lastId++
		store.responseChannel <- lastId

		if err = indexFile.Truncate(0); err != nil {
			return
		}
		if _, err = indexFile.Seek(0, 0); err != nil {
			return
		}

		index.LastId = lastId
		encoder := json.NewEncoder(indexFile)
		err = encoder.Encode(index)
		if err == nil {
			return
		}
	}
}

func (store *Store[T]) Add(item T) (id int32, err error) {
	store.requestChannel <- 1
	id, valid := <-store.responseChannel
	if !valid {
		id = 0
		err = errors.New("Store is closed")
		return
	}

	fileName := fmt.Sprintf("%d.json", id)
	filePath := filepath.Join(store.rootPath, fileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if nil != err {
		id = 0
		return
	}
	defer func() {
		err = errors.Join(err, file.Close())
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(item)
	return
}

func (store *Store[T]) Close() error {
	close(store.requestChannel)
	return nil
}
