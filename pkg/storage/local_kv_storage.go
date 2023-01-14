package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
)

const defaultDBName = "schema_caching"

const LocalKVStorageType = StorageType("LocalKVStorage")

type LocalKVStorage struct {
	db *leveldb.DB
}

func NewLocalKVStorage(name string) *LocalKVStorage {
	if name == "" {
		name = defaultDBName
	}
	db, err := leveldb.OpenFile(name, nil)
	if err != nil {
		return nil
	}
	return &LocalKVStorage{
		db: db,
	}
}

func (lkv *LocalKVStorage) Get(key []byte) ([]byte, error) {
	return lkv.db.Get(key, nil)
}

func (lkv *LocalKVStorage) Put(key, val []byte) error {
	return lkv.db.Put(key, val, nil)
}

func (lkv *LocalKVStorage) Exists(key []byte) bool {
	_, err := lkv.db.Get(key, nil)
	return err != nil
}

func (lkv *LocalKVStorage) Type() StorageType {
	return LocalKVStorageType
}

func init() {
	RegisterStorage(NewLocalKVStorage(""))
}
