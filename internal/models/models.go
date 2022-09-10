package models

import "encoding/json"

// Ports holds the records extracted from the json file
type Ports struct {
	Records map[string]interface{}
}

// Unmarshal is the unmarshal implementation
func (p *Ports) Unmarshal(b []byte) error {
	return json.Unmarshal(b, p)
}
