package core

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

type FileStore interface {
	Save(fileId string, fileType string, binaryData bytes.Buffer) (string, error)
}

type DiskStore struct {
	mutex  sync.RWMutex
	folder string
	files  map[string]*FileInfo
}

type FileInfo struct {
	FileId string
	Type   string
	Path   string
}

func NewDiskStore(folder string) *DiskStore {
	return &DiskStore{
		folder: folder,
		files:  make(map[string]*FileInfo),
	}
}

func (store *DiskStore) Save(fileId string, fileType string, binaryData bytes.Buffer) (string, error) {
	var filePath string

	if fileType == "public" {
		filePath = fmt.Sprintf("%s/public/%s", store.folder, fileId)
	} else {
		filePath = fmt.Sprintf("%s/%s", store.folder, fileId)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot create file: %w", err)
	}

	_, err = binaryData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.files[fileId] = &FileInfo{
		FileId: fileId,
		Type:   fileType,
		Path:   filePath,
	}
	return fileId, nil
}
