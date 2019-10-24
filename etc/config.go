package etc

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
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
