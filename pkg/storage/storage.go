package storage

type StorageType string

type Storage interface {
	Get([]byte) ([]byte, error)
	Exists([]byte) bool
	Put([]byte, []byte) error
	Type() StorageType
}

var storageReg = make(map[StorageType]Storage, 0)

func RegisterStorage(s Storage) {
	storageReg[s.Type()] = s
}

func GetStorage(name StorageType) Storage {
	return storageReg[name]
}
