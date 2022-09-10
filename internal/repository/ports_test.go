package repository

import (
	"github.com/danieeelfr/my-large-json-file-reader/internal/config"
	"github.com/stretchr/testify/suite"
)

type PortsRepoTestSuite struct {
	suite.Suite
	config      *config.Config
	repo        PortsRepo
	mockedPorts map[string]interface{}
}

func (s *PortsRepoTestSuite) BeforeTest(suiteName, testName string) {
	s.config = &config.Config{
		Redis: &config.RedisConfig{},
	}
	s.repo = NewRepository(s.config)
}

func (s *PortsRepoTestSuite) TestGetWithSuccess() {
	result, err := s.repo.Get("", 100)
	s.NoError(err)
	s.Greater(len(result), 0)
}

func (s *PortsRepoTestSuite) TestSetWithSuccess() {
	err := s.repo.Set(s.mockedPorts)
	s.NoError(err)
}
