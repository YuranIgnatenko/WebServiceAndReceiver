package server

import (
	"WebService/app"
	"WebService/ctrldb"
	"WebService/logger"
	"fmt"
	"net"
	"net/http"
)

var Log = logger.NewLogger()

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err)))
}

func WriteAnswer(w http.ResponseWriter, key, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"%s\":\"%s\"}", key, data)))
}

func WriteData(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("%v", data)))
}

type ServerBox struct {
	Domain         string
	DB             *ctrldb.Ctrl
	Conf           *app.Config
	Conn           net.Conn
	TypeConnection string
	HOSTreceiver   string
	PORTreceiver   string
}

func NewServerBox(typeconn, host, port, hostreceiver, portreceiver string, db *ctrldb.Ctrl) *ServerBox {
	serverbox := ServerBox{}
	serverbox.Domain = host + ":" + port
	serverbox.DB = db
	serverbox.Conn = nil
	serverbox.TypeConnection = typeconn
	serverbox.HOSTreceiver = hostreceiver
	serverbox.PORTreceiver = portreceiver

	return &serverbox
}

func (serverbox *ServerBox) StartServer() {
	http.HandleFunc("/test1", serverbox.HandlerTest1)
	http.HandleFunc("/test2", serverbox.HandlerTest2)
	http.HandleFunc("/test3", serverbox.HandlerTest3)
	http.ListenAndServe(serverbox.Domain, nil)
}
