package local

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJobs(t *testing.T) {
	s := localFileStorage{
		dataDir: "testdata",
	}
	jobs := s.LoadJobs()
	assert.Equal(t, 1, len(jobs))
}
