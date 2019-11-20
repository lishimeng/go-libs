package file

import (
	"io/ioutil"
	"os"
	"time"
)

func RemoveSrc(file string) error {
	err := os.Remove(file)
	time.Sleep(2 * time.Second)
	return err
}

func Copy(from string, to string) (err error) {

	in, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(to, in, 0666)
	if err != nil {
		return err
	}
	return err
}
