package pathlib

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Exists(pathname string) bool {
	_, err := os.Stat(pathname)
	return !os.IsNotExist(err)
}

func Abs(path string) (string, error) {
	fullpath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Clean(fullpath), nil
}

func IsDir(pathname string) bool {
	stat, err := os.Stat(pathname)
	return err == nil && stat.IsDir()
}

func IsFile(pathname string) bool {
	stat, err := os.Stat(pathname)
	return err == nil && !stat.IsDir()
}

func Size(pathname string) (int64, bool) {
	stat, err := os.Stat(pathname)
	if err != nil {
		return 0, false
	}
	return stat.Size(), true
}

func Modtime(pathname string) (time.Time, error) {
	stat, err := os.Stat(pathname)
	if err != nil {
		return time.Now(), err
	}
	return stat.ModTime(), nil
}

func EnsureDirectory(directory string) (string, error) {
	fullpath, err := filepath.Abs(directory)
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(fullpath, 0o750)
	if err != nil {
		return "", err
	}
	stats, err := os.Stat(fullpath)
	if !stats.IsDir() {
		return "", fmt.Errorf("Path %s is not a directory!", fullpath)
	}
	return fullpath, nil
}

func EnsureParentDirectory(resource string) (string, error) {
	return EnsureDirectory(filepath.Dir(resource))
}
