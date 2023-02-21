package file

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	SuperPerm     = 0777
	ReadWritePerm = 0766
	ReadExecPerm  = 0755
	OnlyReadPerm  = 0744
)

// IsEmptyDir returns true if the directory is empty.
func IsEmptyDir(path string) (result bool, err error) {
	if len(path) == 0 {
		return false, errors.New("the path can not be empty")
	}

	dir, err := os.ReadDir(path)
	if err != nil {
		return false, fmt.Errorf("failed to load directory, err: %v", err)
	}

	if len(dir) > 0 {
		return false, nil
	}

	return true, nil
}

// MkdirIfNotExist creates a directory from given path if the path does not contain this directory.
// perm is the permission number for directory such as 777, we provide with SuperPerm, ReadWritePerm, ReadExecPerm,
// OnlyReadPerm.
func MkdirIfNotExist(path string, perm uint32) (err error) {
	if len(path) == 0 {
		return errors.New("the path can not be empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, fs.FileMode(perm))
	}

	return nil
}

// GetFilesPathFromDir returns file path slice from a directory.
// If onlyName is true, it will only return name slice.
func GetFilesPathFromDir(path string, onlyName bool) (result []string, err error) {
	if len(path) == 0 {
		return nil, errors.New("the path can not be empty")
	}
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	_ = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		if onlyName {
			result = append(result, filepath.Base(path))
		} else {
			result = append(result, path)
		}
		return err
	})
	return result, nil
}

// RemoveDir removes all the files in the target directory
func RemoveDir(path string) error {
	if len(path) == 0 {
		return errors.New("the path can not be empty")
	}

	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, d := range dir {
		err = os.RemoveAll(filepath.Join([]string{path, d.Name()}...))
		if err != nil {
			return err
		}
	}

	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}
