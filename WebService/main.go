package main

import (
	. "WebService/app"
	"WebService/ctrldb"
	"WebService/logger"
	"WebService/server"
	"flag"
)

func init() {
	flag.StringVar(&HostBD, "host", "", "host redis database (xxxx)")
	flag.StringVar(&PortBD, "port", "", "port redis database (xxx.xxx.xxx.xxx or localhost)")
	flag.Parse()
}

func main() {
	Log := logger.NewLogger()
	Log.INFO("start service")

	if err := ValidateHostPortDB(HostBD, PortBD); err != nil {
		Log.ERROR(err)
		panic(err)
		return
	}

	db := ctrldb.NewCtrl(HostBD, PortBD, PassBD, NumBD)
	serverbox := server.NewServerBox(TypeConnection, Host, Port, HOSTreceiver, PORTreceiver, db)
	serverbox.StartServer()
}
