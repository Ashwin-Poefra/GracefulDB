package ecowriter

import (
	"bytes"
	"encoding/json"
	"os"
)

var (
	encoder *json.Encoder
	decoder *json.Decoder
)

// Saving the structure to a JSON file.
func WriteJSON(name string, data any) (err error) {
	// This function is complete
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		if err2 := os.Remove(name); err2 != nil {
			return err2
		}
	}

	wFile, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer wFile.Close()

	encoder = json.NewEncoder(wFile)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// Loading a structure from a JSON file.
func ReadJSON(name string, data any) (err error) {
	// This function is complete
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return err
	}

	rFile, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer rFile.Close()

	decoder = json.NewDecoder(rFile)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	return nil
}

func EncodeString(data any) string {
	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return ""
	}

	return buf.String()
}
