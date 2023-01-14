package loader

import (
	"bytes"
	"encoding/json"

	"github.com/piprate/json-gold/ld"
	"github.com/sirupsen/logrus"
	"github.com/suutaku/ld-loader/pkg/storage"
)

type CachingDocumentLoader struct {
	nextLoader ld.DocumentLoader
	cache      storage.Storage
}

func NewCachingDocumentLoader(nextLoader ld.DocumentLoader, sType storage.StorageType) *CachingDocumentLoader {
	return &CachingDocumentLoader{
		nextLoader: nextLoader,
		cache:      storage.GetStorage(sType),
	}
}

func (cld *CachingDocumentLoader) createCache(key string, val interface{}) error {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		return err
	}
	return cld.cache.Put([]byte(key), buf.Bytes())
}

// LoadDocument returns a RemoteDocument containing the contents of the JSON resource
// from the given URL.
func (cld *CachingDocumentLoader) LoadDocument(u string) (ret *ld.RemoteDocument, err error) {
	// try caching first
	raw, err := cld.cache.Get([]byte(u))
	if err != nil {
		logrus.Warn("use default remove loader")
		// load with remote loader
		ret, err = cld.nextLoader.LoadDocument(u)
		if err == nil {
			nerr := cld.createCache(u, ret)
			if nerr != nil {
				logrus.Error(nerr)
			}
		}
		return ret, err
	}
	buf := bytes.NewBuffer(raw)
	err = json.NewDecoder(buf).Decode(&ret)
	return ret, err

}
