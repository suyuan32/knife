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
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempFiles() error {
	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")
	err := os.MkdirAll(tmpDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	if dir, _ := os.ReadDir(tmpDirPath); len(dir) == 4 {
		return err
	}

	for i := 0; i < 4; i++ {
		file, err := os.Create(filepath.Join(tmpDirPath, fmt.Sprintf("%d.txt", i)))
		if err != nil {
			return err
		}
		_, err = file.Write([]byte("hello world"))
		if err != nil {
			return err
		}
		file.Close()
	}

	return err
}

func TestIsEmptyDir(t *testing.T) {
	err := createTempFiles()
	assert.Nil(t, err)

	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")
	dir, err := IsEmptyDir(tmpDirPath)
	assert.Nil(t, err)
	assert.Equal(t, dir, false)

	err = os.MkdirAll(tmpDirPath+"_empty", os.ModePerm)
	assert.Nil(t, err)

	dir, err = IsEmptyDir(tmpDirPath + "_empty")
	assert.Nil(t, err)
	assert.Equal(t, dir, true)

	dir, err = IsEmptyDir("")
	assert.NotNil(t, err)

	dir, err = IsEmptyDir("/knife")
	assert.NotNil(t, err)
}

func TestMkdirIfNotExist(t *testing.T) {
	tmpDirPath := filepath.Join(os.TempDir(), "knife_test_mkdir")
	err := MkdirIfNotExist(tmpDirPath, SuperPerm)
	assert.Nil(t, err)

	err = RemoveDir(tmpDirPath)
	assert.Nil(t, err)

	err = MkdirIfNotExist(tmpDirPath, SuperPerm)
	assert.Nil(t, err)

	err = MkdirIfNotExist("", SuperPerm)
	assert.NotNil(t, err)
}

func TestGetFilesPathFromDir(t *testing.T) {
	err := createTempFiles()
	assert.Nil(t, err)

	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")
	dir, err := GetFilesPathFromDir(tmpDirPath, false)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(dir))

	dir, err = GetFilesPathFromDir(tmpDirPath, true)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(dir))

	dir, err = GetFilesPathFromDir("", false)
	assert.NotNil(t, err)

	dir, err = GetFilesPathFromDir("knife", false)
	assert.Nil(t, dir)
}

func TestRemoveDir(t *testing.T) {
	err := createTempFiles()
	assert.Nil(t, err)

	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")

	err = RemoveDir(tmpDirPath)
	assert.Nil(t, err)

	err = RemoveDir("")
	assert.NotNil(t, err)

	err = RemoveDir("/tmp/test_data")
	assert.NotNil(t, err)
}

func TestContainFile(t *testing.T) {
	err := createTempFiles()
	assert.Nil(t, err)

	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")

	fileExist, err := ContainFile(tmpDirPath, "1.txt")
	assert.Nil(t, err)
	assert.Equal(t, fileExist, true)

	fileExist, err = ContainFile(tmpDirPath, "10.txt")
	assert.Nil(t, err)
	assert.Equal(t, fileExist, false)

	fileExist, err = ContainFile("", "10.txt")
	assert.NotNil(t, err)
}

func TestFileExists(t *testing.T) {
	err := createTempFiles()
	assert.Nil(t, err)

	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")

	result := FileExist(filepath.Join(tmpDirPath, "1.txt"))
	assert.Equal(t, true, result)

	result = FileExist(filepath.Join(tmpDirPath, "10.txt"))
	assert.Equal(t, false, result)
}

func TestCreateIfNotExist(t *testing.T) {
	tmpFilePath := filepath.Join(os.TempDir(), "knife_test_create", "5.txt")
	err := MkdirIfNotExist(filepath.Join(os.TempDir(), "knife_test_create"), SuperPerm)
	assert.Nil(t, err)

	file, err := CreateIfNotExist(tmpFilePath)
	assert.Nil(t, err)
	file.Close()

	file2, err := CreateIfNotExist(tmpFilePath)
	assert.NotNil(t, err)
	file2.Close()

	err = RemoveIfExist(tmpFilePath)
	assert.Nil(t, err)

	err = RemoveIfExist(tmpFilePath)
	assert.Nil(t, err)
}

func TestGetSubDir(t *testing.T) {
	tmpDirPath := filepath.Join(os.TempDir(), "knife_test_mkdir")
	err := MkdirIfNotExist(tmpDirPath+"/1", SuperPerm)
	assert.Nil(t, err)

	err = MkdirIfNotExist(tmpDirPath+"/1/2", SuperPerm)
	assert.Nil(t, err)

	dir, err := GetSubDir(tmpDirPath, false)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(dir))

	dir, err = GetSubDir(tmpDirPath, true)
	assert.Nil(t, err)
	assert.Equal(t, "1", dir[0])

	dir, err = GetSubDir("", false)
	assert.NotNil(t, err)

	dir, err = GetSubDir("/knife", false)
	assert.NotNil(t, err)

	err = RemoveDir(tmpDirPath)
	assert.Nil(t, err)
}

func TestReadFileString(t *testing.T) {
	err := createTempFiles()
	assert.Nil(t, err)

	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")

	fileString, err := ReadFileString(filepath.Join(tmpDirPath, "1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, "hello world", fileString)

	fileString, err = ReadFileString("")
	assert.NotNil(t, err)
}

func TestWriteFileString(t *testing.T) {
	err := createTempFiles()
	assert.Nil(t, err)

	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")

	err = RemoveIfExist(filepath.Join(tmpDirPath, "0.txt"))
	assert.Nil(t, err)

	err = WriteFileString(filepath.Join(tmpDirPath, "0.txt"), "Simple Admin", SuperPerm)
	assert.Nil(t, err)

	err = WriteFileString(filepath.Join(tmpDirPath, "0.txt"), "Simple Admin", SuperPerm)
	assert.Nil(t, err)

	fileString, err := ReadFileString(filepath.Join(tmpDirPath, "0.txt"))
	assert.Nil(t, err)
	assert.Equal(t, "Simple Admin", fileString)
}

func TestAppendFileString(t *testing.T) {
	err := createTempFiles()
	assert.Nil(t, err)

	tmpDirPath := filepath.Join(os.TempDir(), "knife_test")

	err = AppendFileString(filepath.Join(tmpDirPath, "2.txt"), "Hi!", SuperPerm)
	assert.Nil(t, err)

	err = AppendFileString(filepath.Join(tmpDirPath, "20.txt"), "Hi!", SuperPerm)
	assert.NotNil(t, err)
}
