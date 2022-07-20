package app

import "errors"

var (
	ErrorDecodeBody         = errors.New("error request body. Using: key:a value:1")
	ErrorIncrementing       = errors.New("error incrementing. Check your key")
	ErrorReadConnection     = errors.New("error reading connection. Check you symbol split")
	ErrorUserHostDB         = errors.New("error host database")
	ErrorUserPortDB         = errors.New("error port database")
	ErrorPingDB             = errors.New("error ping test database")
	ErrorCreatedFile        = errors.New("error created file")
	ErrorWriteString        = errors.New("error write string in file")
	ErrorConnectWebReceiver = errors.New("error connected to WebReceiver. Check host and port")
	ErrorDialHost           = errors.New("error dial host")
	ErrorConnected          = errors.New("error tcp connected. Check service enable")
	ErrorBuilderAnswer      = errors.New("error builder line. Check data json")
)
