package Json

import (
	"encoding/json"
	"os"
)

func JsonToStruct(namefile string, structure any) error {
	dataJson, err := os.ReadFile(namefile)
	// fmt.Println(string(dataJson))
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(dataJson), &structure)
	if err != nil {
		return err
	}
	// fmt.Printf("%#v\n", structure)
	return nil
}

func StructToJson(structure any) (string, error) {
	u, err := json.Marshal(structure)
	if err != nil {
		return "", nil
	}
	return string(u), nil // {"Name":"Bob","Age":10,"Active":true}
}

func StructToJsonByte(structure any) ([]byte, error) {
	u, err := json.Marshal(structure)
	if err != nil {
		return nil, nil
	}
	return u, nil // {"Name":"Bob","Age":10,"Active":true}
}

//  rewrite configure file
// func StructToConfig(namefile string, structure any) error {
// 	data, err := StructToJson(structure)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = os.WriteFile(namefile, data)

// 	return err
// }
