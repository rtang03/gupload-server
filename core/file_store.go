package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	// make an public/listing.txt
	publicDir := fmt.Sprintf("%s/public", store.folder)
	indexTxt := fmt.Sprintf("%s/public/index.txt", store.folder)

	files, err := ioutil.ReadDir(publicDir)
	if err != nil {
		fmt.Println(err)
		return fileId, nil
	}

	f, err := os.Create(indexTxt)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return fileId, nil
	}

	for _, file := range files {
		_, _ = fmt.Fprintln(f, file.Name())
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return fileId, nil
	}
	fmt.Println("index.txt created")

	return fileId, nil
}
