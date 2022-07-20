package receiver

import (
	"fmt"
	"testing"
)

var logger = NewLogger()
var conf = NewConfig()
var receivebox = NewReceicerBox(conf, logger)

type dataDot struct {
	sym        string
	inputLine  interface{}
	outputLine string
}

var arrayDataDot = []dataDot{
	// ==
	dataDot{
		"==",
		interface{}("2,2\n2,2"),
		"4\n4*",
	},
	dataDot{
		"==",
		interface{}("2,3\n3,2"),
		"6\n6*",
	},
	dataDot{
		"==",
		interface{}("12,3\n15,4"),
		"36\n60*",
	},
	// !=
	dataDot{
		"!=",
		interface{}("2,2\n2,2"),
		"4\n4",
	},
	dataDot{
		"!=",
		interface{}("2,3\n3,2"),
		"66*",
	},
	dataDot{
		"!=",
		interface{}("12,3\n15,4"),
		"",
	},
	// err
	dataDot{
		"err",
		interface{}("12,3\n15,4*"),
		"",
	},
	dataDot{
		"err",
		interface{}("*"),
		"",
	},
}

func TestDot(t *testing.T) {
	for _, s := range arrayDataDot {
		res, err := receivebox.Dot(s.inputLine)
		switch s.sym {
		case "==":
			if !(res == s.outputLine) {
				t.Errorf("Error : \n\twait result(%v)\n\thave result(%v)", s.outputLine, res)
				fmt.Printf("Error : \n\twait result(%v)\n\thave result(%v)", s.outputLine, res)
			}
		case "!=":
			if !(res != s.outputLine) {
				t.Errorf("Error : \n\twait result(%v)\n\thave result(%v)", s.outputLine, res)
				fmt.Printf("Error : \n\twait result(%v)\n\thave result(%v)", s.outputLine, res)
			}
		case "err":
			if !(err != nil) {
				t.Errorf("Error : \n\twait result(%v)\n\thave result(%v)", s.outputLine, res)
				fmt.Printf("Error : \n\twait result(%v)\n\thave result(%v)", s.outputLine, res)
			}
		}
	}
}
