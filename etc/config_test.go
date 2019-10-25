package etc

import (
	"fmt"
	"os"
	"testing"
)

type testConfig struct {
	A string `toml:"a"`
	B string `toml:"b"`
}

func TestSaveEnvs(t *testing.T) {

	filename := "out.toml"
	defer func() {
		_ = os.Remove(filename)

	}()
	config := testConfig{
		A: "this is a",
		B: "this is b",
	}
	err := SaveEnvs(filename, &config)
	if err != nil {
		t.Fatal(err)
		return
	}

	read := testConfig{}
	path, err := LoadEnvs("out.toml", []string{"."}, &read)

	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(path)
	fmt.Printf("A:%s\n", read.A)
	fmt.Printf("B:%s\n", read.B)
}
