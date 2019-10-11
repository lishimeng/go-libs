package script

import (
	"fmt"
	"testing"
)

func TestJsEngine_Invoke(t *testing.T) {
	testContent := `function decode(fport, data) {
return {"a": 12, "b": "ffdasf"}
	}`

	var vm, err = Create(testContent)
	if err != nil {
		return
	}
	raw, err := vm.Invoke("decode", "", "")
	if err != nil {
		return
	}
	ras, err := raw.Export()
	if err != nil {
		return
	}
	switch r := ras.(type) {
	case map[string]interface{}:
		result := r
		fmt.Println(result)
	default:
		fmt.Println("type err")
	}
}
