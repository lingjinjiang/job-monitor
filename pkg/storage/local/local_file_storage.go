package local

import (
	"encoding/json"
	"fmt"
	"io/fs"
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
	metaFilePath := filepath.Join(s.dataDir, job.Id+".json")
	data, _ := json.MarshalIndent(job, "", "  ")
	metaFile, err := os.Create(metaFilePath)
	if err != nil {
		log.Println("failed to create file", metaFile, err)
	}
	metaFile.Write(data)
	return nil
}

func (s *localFileStorage) UpdateJob(job api.Job) error {
	metaFilePath := filepath.Join(s.dataDir, job.Id+".json")
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

func (s *localFileStorage) LoadJobs() []api.Job {
	files, err := os.ReadDir(s.dataDir)
	if err != nil {
		log.Println("failed to load data from", s.dataDir, ":", err)
		return nil
	}
	jobs := make([]api.Job, 0)
	for _, file := range files {
		job, err := s.loadJob(file)
		if err != nil {
			continue
		}
		jobs = append(jobs, job)
	}
	return jobs
}

func (s *localFileStorage) loadJob(file fs.DirEntry) (api.Job, error) {
	filename := file.Name()
	if file.IsDir() {
		return api.Job{}, fmt.Errorf("skip directory: %s", filename)
	}
	content, err := os.ReadFile(filepath.Join(s.dataDir, filename))
	if err != nil {
		return api.Job{}, fmt.Errorf("failed to load %s: %s", filename, err)
	}
	var job api.Job
	if err := json.Unmarshal(content, &job); err != nil {
		return api.Job{}, fmt.Errorf("failed to parse %s: %s", filename, err)
	}
	if job.Id+".json" != filename {
		return api.Job{}, fmt.Errorf("job's id is not equal with filename: %s", filename)
	}
	return job, nil
}
