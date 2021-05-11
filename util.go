package mconfig

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ConfigParseError denotes failing to parse configuration file.
type ConfigParseError struct {
	err error
}

// Error returns the formatted configuration error.
func (pe ConfigParseError) Error() string {
	return fmt.Sprintf("While parsing config: %s", pe.err.Error())
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open file=%s fail:%v", filename, err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read file=%s fail:%v", filename, err)
	}

	return buf, nil
}

func insensitiviseMap(m map[string]interface{}) {
	for key, val := range m {
		switch val.(type) {
		case map[string]interface{}:
			// nested map: recursively insensitivise
			insensitiviseMap(val.(map[string]interface{}))
		}

		// update map
		m[key] = val
	}
}
