package cache

// ETag-based cache system implementation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const cacheDirName = "zube"

// Cache file data structure
type Cache struct {
	Etag string      `json:"etag"` // remote resource HTTP ETag
	Data interface{} `json:"data"` // resource data, taken directly from HTTP response body
}

// Initialize ETag based cache system
func Init() {
	os.Mkdir(zubeCacheDir(), 0770)
}

// Purge all cache entries manually by deleting all ETag files
func Purge() {
	fmt.Println("Purging Zube cache...")
	os.RemoveAll(zubeCacheDir())
	Init()
}

// Try to get a cache entry. Returns empty cache and falsy if does not exist, otherwise truthy.
func Get(key string) (Cache, bool) {
	file, err := os.OpenFile(filepath.Join(zubeCacheDir(), key), os.O_RDONLY, 0666)
	defer file.Close()
	if errors.Is(err, os.ErrNotExist) {
		return Cache{}, false
	}

	bytes, _ := ioutil.ReadAll(file)
	var cache Cache
	json.Unmarshal(bytes, &cache)
	return cache, true
}

// Save data under a SHA1 key hash, with an ETag and raw data
func Save(key, etag string, raw []byte) error {
	var data interface{}
	json.Unmarshal(raw, &data)
	cache := Cache{etag, data}

	cacheData, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(zubeCacheDir(), key), cacheData, 0666)
	if err != nil {
		return err
	}

	return nil
}

func zubeCacheDir() string {
	cacheDir, _ := os.UserCacheDir()
	return filepath.Join(cacheDir, cacheDirName)
}
