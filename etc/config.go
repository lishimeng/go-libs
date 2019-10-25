package etc

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

func LoadEnvs(configName string, envPath []string, Config interface{}) (f string, err error) {

	for _, path := range envPath {
		file := fmt.Sprintf("%s/%s", path, configName)
		f, err = filepath.Abs(file)
		if err != nil {
			continue
		} else {
			if _, err := toml.DecodeFile(f, Config); err != nil {
				return f, err
			} else {
				return f, err
			}
		}
	}
	return f, errors.New("can't find config file")
}

func SaveEnvs(filename string, config interface{}) (err error) {

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	enc := toml.NewEncoder(f)
	err = enc.Encode(&config)
	return err
}
