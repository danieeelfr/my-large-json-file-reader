package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/danieeelfr/my-large-json-file-reader/internal/config"
	"github.com/danieeelfr/my-large-json-file-reader/internal/models"
)

// PortService is the service interface
type PortService interface {
	ReadAndDecode() (*models.Ports, error)
}

// Service is the implementation of the PortsService interface
type Service struct {
	filePath string
	Records  map[string]interface{}
}

// NewPortService returns a PortService implementation
func NewPortService(cfg *config.Config) PortService {
	return &Service{
		filePath: cfg.FileReader.FilePath,
	}
}

// ReadAndDecode is resposible for read and decode the given json file...
func (s *Service) ReadAndDecode() (*models.Ports, error) {
	fmt.Println("starting the read process...")
	start := time.Now()

	f, err := os.Open(s.filePath)
	if err != nil {
		return nil, err
	}

	rd := bufio.NewReader(f)

	ports, err := decode(rd)
	if err != nil {
		return nil, err
	}

	elapsed := time.Since(start)

	fmt.Printf("The read process took [%v] and found [%d] records\n", elapsed, len(ports.Records))

	return ports, nil
}

func decode(rd io.Reader) (*models.Ports, error) {
	d := json.NewDecoder(rd)

	// This is the way that I found to read the given ports.json file without errors
	if err := expect(d, json.Delim('{')); err != nil {
		return nil, err
	}

	result := new(models.Ports)
	result.Records = make(map[string]interface{})

	for d.More() {
		titleKey, err := d.Token()
		if err != nil {
			return nil, err
		}

		mod := new(models.Ports)
		if err := d.Decode(&mod.Records); err != nil {
			return nil, err
		}
		result.Records[titleKey.(string)] = mod.Records

	}

	return result, nil
}

// expect returns an error if the next token in the document is not expected.
func expect(d *json.Decoder, expectedT interface{}) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != expectedT {
		return fmt.Errorf("got token %v, want token %v", t, expectedT)
	}
	return nil
}
