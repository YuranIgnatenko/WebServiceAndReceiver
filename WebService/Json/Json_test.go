package Json

import (
	"go/build"
	"testing"
)

func TestF(t *testing.T) {
	type Data struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		Flag bool   `json:"flag"`
	}

	filename := build.Default.GOPATH + "src/WebService/Json/data.json"
	data := Data{}
	err := JsonToStruct(filename, &data)
	if err != nil {
		t.Errorf("error: %v", err)
	}

}
