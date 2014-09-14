package io

import (
	"os"
	"path/filepath"
)

func IsFile(path string) bool {
	f, err := os.Stat(path)
	return err != nil && !f.IsDir()
}

func IsDir(path string) bool {
	f, err := os.Stat(path)
	return err != nil && f.IsDir()
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func EnsureDirExists(dirPaths ...string) {
	for _, dirPath := range dirPaths {
		if !IsExists(dirPath) {
			parentPath := filepath.Dir(dirPath)
			parentInfo, err := os.Stat(parentPath)
			if err != nil {
				panic(err)
			}
			os.MkdirAll(dirPath, parentInfo.Mode())
		}
	}
}

func RemoveFile(path string) bool {
	if IsExists(path) {
		err := os.Remove(path)
		if err != nil {
			return false
		}
	}
	return true
}

func WriteText(path string, text string) (int, error) {
	f, err := os.Create(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.WriteString(text)
}

func GetText(path string) (string, error) {
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

type WalkFileFunc func(fileInfo os.FileInfo)

func WalkDirFiles(root string, handler WalkFileFunc) {
	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		handler(fi)
		return nil
	})
}
