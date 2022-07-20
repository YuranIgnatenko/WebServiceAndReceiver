package app

import (
	"WebService/ctrldb"
	"strconv"
)

// check line is nil
func IsNotNil(s string) bool {
	if s == "" {
		return false
	}

	return true
}


func ValidateValue(s string) error {
	_, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	return nil
}

// check is nil params started service
func ValidateHostPortDB(host, port string) error {
	if host == "" {
		return ErrorUserHostDB
	}
	if port == "" {
		return ErrorUserPortDB
	}
	db := ctrldb.NewCtrl(host, port, "", 0)
	_, err := db.Ping()
	if err != nil {
		return ErrorPingDB
	}
	db = nil
	return nil
}
