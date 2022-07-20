package server

import (
	. "WebService/app"
	"WebService/logger"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func ParsingData(line []interface{}) (string, []string) {
	line_num := ""
	keys := make([]string, 0)
	for _, elem := range line {
		switch elem.(type) {
		case map[string]interface{}:
			el := elem.(map[string]interface{})
			for k, val := range el {
				if k == "key" {
					continue
				}
				line_num += val.(string) + ","
			}
			line_num = line_num[:len(line_num)-1] + "\n"
			keys = append(keys, el["key"].(string))
		}
	}
	return line_num, keys
}

func BuilderLineAnswer(keys []string, line string) (string, error) {
	line = strings.TrimSpace(line)
	data := strings.Split(line, "\n")

	result_line := ""
	result_line += "{"

	if len(keys) != len(data) {
		return fmt.Sprintf("{error:%v}", ErrorBuilderAnswer), ErrorBuilderAnswer

	} else {
		for indkey, key := range keys {
			total := 0.0
			for ind, val := range strings.Split(data[indkey], ",") {

				f, err := strconv.ParseFloat(val, 32)
				if err != nil {
					panic(err)
				}
				if ind == 0 {
					total += f
				} else {
					total *= f
				}
			}

			result_line += key + ":" + fmt.Sprintf("%v", total) + ","
		}
	}
	result_line = result_line[:len(result_line)-1]
	result_line += "}"

	return result_line, nil
}

// example request
/*
	[
	  {
	    "a": "99",
	    "b": "2",
	    "key": "x"
	  },
	  {
	    "a": "11",
	    "b": "2",
	    "key": "y"
	  }
	]
// returned
{
	"x": "188",
	"y": "22"
}
*/
func (serverbox *ServerBox) HandlerTest3(w http.ResponseWriter, r *http.Request) {
	Log := logger.NewLogger()

	if serverbox.Conn == nil {
		conn, err := net.Dial(serverbox.TypeConnection, serverbox.HOSTreceiver+":"+serverbox.PORTreceiver)
		if err != nil {
			serverbox.Conn = nil
			Log.ERROR(ErrorConnected)
			WriteError(w, ErrorConnected)
			return
		} else {
			Log.INFO("connected tcp: ok")
			serverbox.Conn = conn
		}
	}

	d := struct {
		Array []interface{}
	}{}

	err := json.NewDecoder(r.Body).Decode(&d.Array)
	if err != nil {
		Log.ERROR(ErrorDecodeBody)
		WriteError(w, ErrorDecodeBody)
		return
	}

	keys, linejs := ParsingData(d.Array)

	w.Header().Set("Content-Type", "application/json")
	nline, err := BuilderLineAnswer(linejs, keys)
	if err != nil {
		Log.ERROR(err)
		WriteError(w, err)
		return
	}
	WriteData(w, nline)

}
