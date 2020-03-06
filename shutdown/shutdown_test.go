package shutdown

import (
	"fmt"
	"testing"
	"time"
)

func TestWaitExit(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			t.Fatal(err)
			return
		}
	}()

	go func() {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("ctx done")
				return
			default:
				time.Sleep(time.Millisecond * 200)
			}
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		Exit("call exit in TestWaitExit")
	}()
	WaitExit(&Configuration{
		BeforeExit: func(s string) {
			fmt.Printf("exit message:%s\n", s)
		},
	})
}
