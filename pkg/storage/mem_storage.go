package storage

import (
	"fmt"
)

const MemStorageType = StorageType("MemStorage")

type MemStorage struct {
	mem map[string][]byte
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		mem: make(map[string][]byte),
	}
}

func (mem *MemStorage) Get(key []byte) ([]byte, error) {
	ret := mem.mem[string(key)]
	if ret == nil {
		return ret, fmt.Errorf("not exist")
	}
	return ret, nil
}

func (mem *MemStorage) Put(key []byte, val []byte) error {
	mem.mem[string(key)] = val
	return nil
}

func (mem *MemStorage) Exists(key []byte) bool {
	_, contains := mem.mem[string(key)]
	return contains
}

func (mem *MemStorage) Type() StorageType {
	return MemStorageType
}

func init() {
	RegisterStorage(NewMemStorage())
}
