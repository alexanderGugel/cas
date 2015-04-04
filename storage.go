package cas

import (
    "crypto/sha1"
    "encoding/hex"
    "io"
    "io/ioutil"
    "path/filepath"
    "os"
)

// [Content-Addressed Storage](http://de.wikipedia.org/wiki/Content-Addressed_Storage)
type Storage struct {
    KeyToPath map[string]string
    PathToKey map[string]string
}

// imports a file
func (c *Storage) ImportFile(path string) (key string, err error) {
    path, err = filepath.Abs(path)
    if err != nil {
        return "", err
    }
    c.DelByPath(path)
    file, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer file.Close()
    hash := sha1.New()
    _, err = io.Copy(hash, file)
    if err != nil {
        return "", err
    }
    key = hex.EncodeToString(hash.Sum(nil))
    c.KeyToPath[key] = path
    c.PathToKey[path] = key
    return key, nil
}

// recursively imports a directory
func (c *Storage) ImportDir(path string) (err error) {
    fis, err := ioutil.ReadDir(path)
    if err != nil {
        return err
    }
    for _, fi := range fis {
        if fi.IsDir() {
            err = c.ImportDir(path + string(os.PathSeparator) + fi.Name())
        } else {
            _, err = c.ImportFile(path + string(os.PathSeparator) + fi.Name())
        }
        if err != nil {
            return err
        }
    }
    return nil
}

// retrieves a file by the key returned when importing it
func (c *Storage) Get(key string) (*os.File, error) {
    return os.Open(c.KeyToPath[key])
}

func (c *Storage) DelByPath(path string) (err error) {
    path, err = filepath.Abs(path)
    if err != nil {
        return err
    }
    key := c.PathToKey[path]
    delete(c.PathToKey, path)
    delete(c.KeyToPath, key)
    return nil
}

func (c *Storage) DelByKey(key string) () {
    path := c.KeyToPath[key]
    delete(c.PathToKey, path)
    delete(c.KeyToPath, key)
}

func (c *Storage) String() string {
    s := ""
    for key, value := range c.KeyToPath {
        s = s + key + value + "\n"
    }
    return s
}

func New() *Storage {
    return &Storage{
        KeyToPath: make(map[string]string),
        PathToKey: make(map[string]string),
    }
}
