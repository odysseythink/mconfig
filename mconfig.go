package mconfig

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"mlib.com/mlog"
)

var (
	ErrKeyNotExist      = errors.New("key not exist")
	ErrInvalidValueType = errors.New("invalid value type")
)

var v *Viper

func init() {
	v = New()
}

// New returns an initialized Viper instance.
func New() *Viper {
	v := new(Viper)
	v.keyDelim = "."
	v.configName = "config"
	v.configPermissions = os.FileMode(0644)
	v.config = make(map[string]interface{})
	v.isChanged = false
	return v
}

// Viper is a prioritized configuration registry. It
// maintains a set of configuration sources, fetches
// values to populate those, and provides them according
// to the source's priority.
// The priority of the sources is the following:
// 1. overrides
// 2. flags
// 3. env. variables
// 4. config file
// 5. key/value store
// 6. defaults
//
// For example, if values from the following sources were loaded:
//
//  Defaults : {
//  	"secret": "",
//  	"user": "default",
//  	"endpoint": "https://localhost"
//  }
//  Config : {
//  	"user": "root"
//  	"secret": "defaultsecret"
//  }
//  Env : {
//  	"secret": "somesecretkey"
//  }
//
// The resulting config will have the following values:
//
//	{
//		"secret": "somesecretkey",
//		"user": "root",
//		"endpoint": "https://localhost"
//	}
type Viper struct {
	// Delimiter that separates a list of keys
	// used to access a nested value in one go
	keyDelim string

	// A set of paths to look for the config file in
	configPaths []string

	// Name of file to look for inside the path
	configName        string
	configFile        string
	configType        string
	configPermissions os.FileMode
	envPrefix         string

	config    map[string]interface{}
	isChanged bool
}

// SetConfigFile explicitly defines the path, name and extension of the config file.
// Viper will use this and not check any of the config paths.
func SetConfigFile(in string) { v.SetConfigFile(in) }
func (v *Viper) SetConfigFile(in string) {
	if in != "" {
		v.configFile = in
	}
}

func (v *Viper) getConfigFile() (string, error) {
	return v.configFile, nil
}

// ReadInConfig will discover and load the configuration file from disk
// and key/value stores, searching in one of the defined paths.
func ReadInConfig() error { return v.ReadInConfig() }
func (v *Viper) ReadInConfig() error {
	mlog.Info("Attempting to read in config file")
	filename, err := v.getConfigFile()
	if err != nil {
		return err
	}

	mlog.Debug("Reading file: ", filename)
	file, err := readFile(filename)
	if err != nil {
		return err
	}

	config := make(map[string]interface{})

	err = v.unmarshalReader(bytes.NewReader(file), config)
	if err != nil {
		return err
	}
	mlog.Debug("config=", config)
	v.config = config
	return nil
}

func (v *Viper) unmarshalReader(in io.Reader, c map[string]interface{}) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	if err := json.Unmarshal(buf.Bytes(), &c); err != nil {
		return ConfigParseError{err}
	}

	insensitiviseMap(c)
	return nil
}

// WriteConfig writes the current configuration to a file.
func WriteConfig() error { return v.WriteConfig() }
func (v *Viper) WriteConfig() error {
	filename, err := v.getConfigFile()
	if err != nil {
		return err
	}
	if v.isChanged {
		mlog.Info("changed, need uncached")
		err = v.writeConfig(filename, true)
		if err == nil {
			v.isChanged = false
		}
		return err
	}
	return nil
}

func (v *Viper) writeConfig(filename string, force bool) error {
	mlog.Info("Attempting to write configuration to file.")

	if v.config == nil {
		v.config = make(map[string]interface{})
	}

	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if !force {
		flags |= os.O_EXCL
	}
	f, err := os.OpenFile(filename, flags, v.configPermissions)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := v.marshalWriter(f, "json"); err != nil {
		return err
	}

	return f.Sync()
}

// Marshal a map into Writer.
func (v *Viper) marshalWriter(f *os.File, configType string) error {
	b, err := json.MarshalIndent(v.config, "", "  ")
	if err != nil {
		return ConfigMarshalError{err}
	}
	_, err = f.WriteString(string(b))
	if err != nil {
		return ConfigMarshalError{err}
	}
	return nil
}

// ConfigMarshalError happens when failing to marshal the configuration.
type ConfigMarshalError struct {
	err error
}

// Error returns the formatted configuration error.
func (e ConfigMarshalError) Error() string {
	return fmt.Sprintf("While marshaling config: %s", e.err.Error())
}
