package jwt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {

	h := New([]byte("secret"), "https://doudou.me", time.Hour * 3 * 24)
	token, success := h.GenToken(Token{
		BaseToken: BaseToken{
			UID: "U13413",
			LoginType: 1,
		},
		Audience:  "U13413",
		Subject:   "https://doudou.me/xiaoha",
	})
	fmt.Println(token)
	fmt.Println(success)
	if len(token) <= 0 {
		t.Fatal()
		return
	}
	if !success {
		t.Fatal()
		return
	}

	c, success := h.VerifyToken(token)
	if !success {
		t.Fatal("verify failed")
		return
	}

	bs, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(string(bs))
}