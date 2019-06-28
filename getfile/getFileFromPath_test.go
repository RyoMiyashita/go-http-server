package getfile

import(
	"fmt"
	"testing"
	"strings"
)

func TestGetIndex(t *testing.T) {
	act, error := GetFileFromPath("../public", "/ok.json")
	if error != nil {
		fmt.Println("err", error)
	}
	act = strings.Replace(act, "\n", "", -1)
	act = strings.Replace(act, " ", "", -1)
	exp := `{"message":"ok"}`
	if act != exp {
		t.Errorf("ok. actual %v\nwant %v", act, exp)
	}
}
