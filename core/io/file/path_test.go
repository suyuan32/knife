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
		_, err = file.Write([]byte("hello word"))
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

	err = RemoveDir("////")
	assert.NotNil(t, err)
}
