package cas

import (
	"path/filepath"
	"testing"
)

func TestImportFile(t *testing.T) {
	storage := New()
	key, err := storage.ImportFile("./storage.go")
	if err != nil {
		t.Errorf("ImportFile(\"./storage.go\") %v", err)
	}
	f, ok, err := storage.Get(key)
	if !ok {
		t.Errorf("Get(\"%v\") should be ok, got \"%v\"", key, ok)
	}
	if err != nil {
		t.Errorf("Get(\"%v\") %v", key, err)
	}
	if !filepath.IsAbs(f.Name()) {
		t.Errorf("Storage should store absolute filepath of files, got \"%v\"", f.Name())
	}
}

func TestImportDir(t *testing.T) {
	storage := New()
	err := storage.ImportDir("./")
	if err != nil {
		t.Errorf("ImportDir(\"../\") %v", err)
	}
}

func TestString(t *testing.T) {
	storage := New()
	err := storage.ImportDir("./")
	if err != nil {
		t.Errorf("ImportDir(\"../\") %v", err)
	}
}
