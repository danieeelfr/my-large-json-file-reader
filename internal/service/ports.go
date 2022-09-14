package service

import (
	"fmt"
	"log"
	"time"

	"github.com/danieeelfr/my-large-json-file-reader/internal/config"
	"github.com/danieeelfr/my-large-json-file-reader/internal/repository"
	stream "github.com/danieeelfr/my-large-json-file-reader/internal/utils"
)

// PortService is the service interface
type PortService interface {
	ReadAndDecode() error
}

// Service is the implementation of the PortsService interface
type Service struct {
	filePath string
	Records  map[string]interface{}
	Repo     *repository.Repository
}

// NewPortService returns a PortService implementation
func NewPortService(cfg *config.Config, repo *repository.Repository) PortService {
	return &Service{
		filePath: cfg.FileReader.FilePath,
		Repo:     repo,
	}
}

// ReadAndDecode is resposible for read and decode the given json file...
func (s *Service) ReadAndDecode() error {
	fmt.Println("starting the read process...")
	start := time.Now()

	stream := stream.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			if data.Error != nil {
				log.Println(data.Error)
				// break
			}
			e := s.Repo.SetPort(data.Title, data.Port)
			if e != nil {
				log.Println(e)
				// break
			}
			// log.Println(len(data.Ports))
		}
	}()
	stream.Start(s.filePath)

	// f, err := os.Open(s.filePath)
	// if err != nil {
	// 	return nil, err
	// }

	// rd := bufio.NewReader(f)

	// ports, err := decode(rd)
	// if err != nil {
	// 	return nil, err
	// }

	elapsed := time.Since(start)

	fmt.Printf("The read process took [%v] and found [] records\n", elapsed)

	return nil
}
