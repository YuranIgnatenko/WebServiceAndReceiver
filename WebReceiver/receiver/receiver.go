package receiver

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

type ReceiverBox struct {
	Logs     *Logger
	Type     string
	Port     string
	SymSplit string
	SizePipe int
}

func NewReceicerBox(conf *Config, logger *Logger) *ReceiverBox {
	rb := ReceiverBox{}
	rb.Port = conf.Port
	rb.Type = conf.TypeConnect
	rb.SymSplit = conf.SymSplit
	rb.SizePipe = 1
	rb.Logs = logger
	return &rb
}

// Принимает интерфейс: "12,3\n15,4"
// и возвращает строку: "36\n60*"
func (rb *ReceiverBox) Dot(line string) (string, error) {
	nums := make([]float64, 0)
	data := strings.Split(line, "\n")

	for _, linedata := range data {
		linevars := strings.Split(linedata, ",")
		f_total := 0.0
		for _, value := range linevars {
			fmt.Println("value from Dot::::::[", value, "]")
			if value == "" {
				continue
			}
			f_value, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return "", err
			}
			if f_total == 0 {
				f_total += f_value
			} else {
				f_total *= f_value
			}
		}
		nums = append(nums, f_total)
	}

	result_line := ""
	for _, elem := range nums {
		result_line += fmt.Sprintf("%v", elem) + "\n"
	}
	result_line = result_line[:len(result_line)-1]
	// result_line += rb.SymSplit
	return result_line, nil
}

func (rb *ReceiverBox) handle(conn net.Conn) {
	// defer conn.Close()
	for {
		data := make([]byte, (1024 * 4))
		// fmt.Println("data get receiver : ", string(data))
		n, err := conn.Read(data)
		if n == 0 || err != nil {
			return
		}
		fmt.Println("ok - get handler")
		source := string(data[0 : n-1])
		result, err := rb.Dot(source)
		if err != nil {
			panic(err)
		}
		fmt.Println("result reseiver : ", result)
		conn.Write([]byte(result))
	}

}

// слушатель сервиса: принимает - отправляет
func (rb *ReceiverBox) RunReceiver() {
	rb.Logs.INFO(fmt.Sprintf("Start server (%v:%v)", rb.Type, rb.Port))
	ln, err := net.Listen(rb.Type, ":"+rb.Port)
	if err != nil {
		rb.Logs.ERROR(err)
		panic(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		rb.Logs.ERROR(err)
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		rb.handle(conn)
		wg.Done()
	}()
	wg.Wait()
}
