package storage

import (
	"log"
	"os"
	"path/filepath"
)

type storageRepo struct {
	dirPath string
}

func NewStorageRepo(dirPath string) *storageRepo {
	return &storageRepo{
		dirPath: dirPath,
	}
}

func (repo *storageRepo) Init() {
	if err := os.MkdirAll(repo.dirPath, os.ModePerm); err != nil {
		log.Fatalf("error occured while initialzing the local storage, Err: %s\n", err.Error())
	}
	log.Println("local storage initialized successfully")
}

func (repo *storageRepo) UpdateMoistureSettings(data []byte) error {
	p := filepath.Join(repo.dirPath, "moisture_settings.json")
	err := os.WriteFile(p, data, 0644)
	return err
}

func (repo *storageRepo) GetMoistureSettings() ([]byte, error) {
	p := filepath.Join(repo.dirPath, "moisture_settings.json")
	d, err := os.ReadFile(p)
	return d, err
}
