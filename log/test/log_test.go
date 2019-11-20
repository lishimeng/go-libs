package test

import (
	"github.com/lishimeng/go-libs/log"
	"testing"
)

func Test_FormatLevel001(t *testing.T) {
	lvl, err := log.FormatLevel("DEBUG")
	if err != nil {
		t.Fatal(err)
		return
	}
	if lvl != log.DEBUG {
		t.Fatalf("expect %d, but %d", log.DEBUG, lvl)
	}
}

func Test_ChangeLevel001(t *testing.T) {
	log.SetLevelAll(log.DEBUG)
	log.Debug("this is debug log")
	log.SetLevelAll(log.INFO)
	log.Debug("if you see this test maybe failed")
}

func Test_SetLevel001(t *testing.T) {
	log.SetLevel("stdout", log.DEBUG)
	log.Debug("this is debug log")
	log.SetLevel("stdout", log.INFO)
	log.Debug("if you see this test maybe failed")
}
