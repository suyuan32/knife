// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// ContainFile returns true if there is the file under the directory path.
func ContainFile(dirPath, fileName string) (bool, error) {
	fileNameData, err := GetFilesPathFromDir(dirPath, true)
	if err != nil {
		return false, err
	}

	for _, v := range fileNameData {
		if v == fileName {
			return true, err
		}
	}

	return false, err
}

// FileExist returns true if the specified file exists.
func FileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// CreateIfNotExist creates a file if it does not exist.
func CreateIfNotExist(filePath string) (*os.File, error) {
	if FileExist(filePath) {
		return nil, fmt.Errorf("%s already exist", filePath)
	}

	return os.Create(filePath)
}

// RemoveIfExist deletes the specified file if exists.
func RemoveIfExist(filename string) error {
	if !FileExist(filename) {
		return nil
	}

	return os.Remove(filename)
}

// GetSubDir returns subdirectory path slice from a directory.
// If onlyName is true, it will only return name slice.
func GetSubDir(path string, onlyName bool) (result []string, err error) {
	if len(path) == 0 {
		return nil, errors.New("the path can not be empty")
	}
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	_ = filepath.WalkDir(path, func(pathData string, d fs.DirEntry, err error) error {
		if d.IsDir() && filepath.Dir(pathData) == path {
			if onlyName {
				result = append(result, filepath.Base(pathData))
			} else {
				result = append(result, pathData)
			}
			return err
		}
		return err
	})
	return result, nil
}

// ReadFileString returns string from a file
func ReadFileString(path string) (string, error) {
	if FileExist(path) {
		fileData, _ := os.ReadFile(path)
		return string(fileData), nil
	} else {
		return "", errors.New("file does not exist")
	}
}

// WriteFileString writes string to a file, if file does not exist, it will
// be created.
func WriteFileString(path, data string, perm int) (err error) {
	var fileData *os.File
	if FileExist(path) {
		fileData, err = os.OpenFile(path, os.O_RDWR, fs.FileMode(perm))
		if err != nil {
			return err
		}
	} else {
		fileData, _ = CreateIfNotExist(path)
	}
	defer fileData.Close()

	_, err = fileData.Write([]byte(data))
	return err
}

// AppendFileString append the string to the target file
func AppendFileString(path, data string, perm int) error {
	if FileExist(path) {
		fileData, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, fs.FileMode(perm))
		if err != nil {
			return err
		}
		defer fileData.Close()

		_, err = fileData.Write([]byte(data))
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("file does not exist")
	}
}
