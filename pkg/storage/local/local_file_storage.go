package local

import (
	"encoding/json"
	"job-monitor/pkg/api"
	"job-monitor/pkg/message"
	"log"
	"os"
	"path/filepath"
)

type localFileStorage struct {
	queue   message.Queue
	dataDir string
}

func NewLocalFileStorage(queue message.Queue, dataDir string) *localFileStorage {
	s := localFileStorage{queue: queue, dataDir: dataDir}
	message.Register("add", s.AddJob)
	message.Register("update", s.UpdateJob)
	message.Register("delete", s.DeleteJob)

	dirInfo, err := os.Stat(s.dataDir)
	if os.IsNotExist(err) {
		log.Fatalln("the data directory", s.dataDir, "doesn't exist.")
	}
	if !dirInfo.IsDir() {
		log.Fatalln(s.dataDir, "is not a directory")
	}

	return &s
}

func (s *localFileStorage) AddJob(job api.Job) error {
	dir := filepath.Join(s.dataDir, job.Type, job.Kind, job.Id)
	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		log.Println(err)
	}
	metaFilePath := filepath.Join(dir, "meta.json")
	data, _ := json.MarshalIndent(job, "", "  ")
	metaFile, err := os.Create(metaFilePath)
	if err != nil {
		log.Println("failed to create file", metaFile, err)
	}
	metaFile.Write(data)
	return nil
}

func (s *localFileStorage) UpdateJob(job api.Job) error {
	dir := filepath.Join(s.dataDir, job.Type, job.Kind, job.Id)
	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		log.Println(err)
	}
	metaFilePath := filepath.Join(dir, "meta.json")
	data, _ := json.MarshalIndent(job, "", "  ")
	metaFile, err := os.Create(metaFilePath)
	if err != nil {
		log.Println("failed to create file", metaFile, err)
	}
	metaFile.Write(data)
	return nil
}

func (s *localFileStorage) DeleteJob(job api.Job) error {
	return nil
}
