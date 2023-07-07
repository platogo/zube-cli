package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const existingKey = "20c22a082fcce4ece7a64f692d9a86fd0f9f06b2"

func setup() {
	Init()
	mockCacheData := Cache{Etag: "4b61-KvCvo1XgbWyNb1RjJY6Ci2/Z0DA", Data: "peko"}
	data, _ := json.Marshal(mockCacheData)
	os.WriteFile(filepath.Join(cacheDir(), existingKey), data, 0666)
}

func teardown() {
	os.Remove(filepath.Join(cacheDir(), existingKey))
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestGet(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  Cache
		want1 bool
	}{
		{"check returns empty cache and falsy on nonexistent cache entry", args{key: "somenonexistentkey"},
			Cache{}, false},
		{"check returns correct cache for key and truthy for existent cache entry", args{key: existingKey},
			Cache{Etag: "4b61-KvCvo1XgbWyNb1RjJY6Ci2/Z0DA", Data: "peko"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
