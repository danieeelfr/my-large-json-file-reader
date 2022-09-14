package service

import (
	"testing"

	"github.com/danieeelfr/my-large-json-file-reader/internal/config"
	"github.com/danieeelfr/my-large-json-file-reader/internal/repository"
	"github.com/stretchr/testify/suite"
)

type PortServiceTestSuite struct {
	suite.Suite
	config  config.Config
	service PortService
}

func (s *PortServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.config = config.Config{
		FileReader: &config.JSONReaderConfig{
			FilePath: "./../../ports.json",
		},
	}
	s.service = NewPortService(&s.config, &repository.Repository{})
}

// func (s *PortServiceTestSuite) TestReadWithSuccess() {
// 	err := s.service.ReadAndDecode()
// 	s.NoError(err)
// 	// s.True(len(result.Records) > 0)
// }

func TestJsonReaderTestSuite(t *testing.T) {
	// suite.Run(t, new(PortServiceTestSuite))
}
