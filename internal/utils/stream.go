package stream

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/danieeelfr/my-large-json-file-reader/internal/models"
)

// Entry represents each stream. If the stream fails, an error will be present.
type Entry struct {
	Error error
	// Ports1 models.Ports
	Port  *models.Port
	Title string
}

// Stream helps transmit each streams withing a channel.
type Stream struct {
	stream chan Entry
}

// NewJSONStream returns a new `Stream` type.
func NewJSONStream() Stream {
	return Stream{
		stream: make(chan Entry),
	}
}

// Watch watches JSON streams. Each stream entry will either have an error or a
// User object. Client code does not need to explicitly exit after catching an
// error as the `Start` method will close the channel automatically.
func (s Stream) Watch() <-chan Entry {
	return s.stream
}

// Start starts streaming JSON file line by line. If an error occurs, the channel
// will be closed.
func (s Stream) Start(path string) {
	// Stop streaming channel as soon as nothing left to read in the file.
	defer close(s.stream)

	// Open file to read.
	file, err := os.Open(path)
	if err != nil {
		s.stream <- Entry{Error: fmt.Errorf("open file: %w", err)}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening delimiter. `[` or `{`
	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode opening delimiter: %w", err)}
		return
	}

	i := 1

	for decoder.More() {
		titleKey, err := decoder.Token()
		if err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode line %d: %w", i, err)}
			return
		}

		port := new(models.Port)

		if err := decoder.Decode(&port); err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode line %d: %w", i, err)}
			return
		}

		s.stream <- Entry{Port: port, Title: titleKey.(string)}

		i++
	}

	// Read closing delimiter. `]` or `}`
	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode closing delimiter: %w", err)}
		return
	}
}
