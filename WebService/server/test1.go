package server

import (
	. "WebService/app"
	"WebService/ctrldb"
	"WebService/logger"
	"encoding/json"
	"net/http"
)

func Incrementing(key, value string) (string, error) {
	db := ctrldb.NewCtrl(HostBD, PortBD, PassBD, NumBD)
	data, err := db.IncrementingValues(key, value)
	if err != nil {
		return "", err
	}
	return data, nil
}

// example request
/*
{
	"key":"k",
	"value":"1"
}
// returned
{
	"k":"13628"
}
*/
func (serverbox *ServerBox) HandlerTest1(w http.ResponseWriter, r *http.Request) {
	Log := logger.NewLogger()
	type data struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	d := data{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		Log.ERROR(ErrorDecodeBody)
		WriteError(w, ErrorDecodeBody)
		return
	}
	result, err := Incrementing(d.Key, d.Value)
	if err != nil {
		Log.ERROR(ErrorIncrementing)
		WriteError(w, ErrorIncrementing)
		return
	}
	WriteAnswer(w, d.Key, result)
	Log.INFO(TEST1_OK)
}
