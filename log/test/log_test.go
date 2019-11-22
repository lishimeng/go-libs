package test

import (
	"github.com/lishimeng/go-libs/log"
	"os"
	"testing"
	"time"
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
	time.Sleep(1 * time.Second)
}

func Test_SetLevel001(t *testing.T) {
	log.SetLevel("stdout", log.DEBUG)
	log.Debug("this is debug log")
	log.SetLevel("stdout", log.INFO)
	log.Debug("if you see this test maybe failed")
	time.Sleep(1 * time.Second)
}

func Test_Finest001(t *testing.T) {
	log.SetLevel("stdout", log.FINEST)
	log.Finest("this is Finest log")
	log.SetLevel("stdout", log.FINE)
	log.Finest("if you see this test maybe failed")
	time.Sleep(1 * time.Second)
}

func Test_Critical001(t *testing.T) {

	err := log.Critical("this is Critical log")
	if err == nil {
		t.Fatalf("expect an err")
	}
	time.Sleep(1 * time.Second)
}

func Test_AddFileLog(t *testing.T) {
	fileName := "a.log"
	log.AddFileLog("a_file", log.DEBUG, fileName, false, false, "%M", 100, 1024)
	log.Debug("this will be written into a file")

	time.Sleep(1 * time.Second)
	log.Close()

	fd, err := os.Open(fileName)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = fd.Close()
	if err != nil {
		t.Fatal(err)
		return
	}
	err = os.Remove(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
