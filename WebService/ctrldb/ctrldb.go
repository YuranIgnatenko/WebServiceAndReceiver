package ctrldb

import (
	// . "WebService/app"
	// "WebService/Json"
	"context"
	"strconv"

	"github.com/go-redis/redis"
)

type Ctrl struct {
	Host     string
	Port     string
	Password string
	DB       int
	Client   *redis.Client
}

// new struct
// "localhost", "6379", "", 0
func NewCtrl(host, port, password string, db int) *Ctrl {
	ctrl := Ctrl{}
	ctrl.Host = host
	ctrl.Port = port
	ctrl.Password = password
	ctrl.DB = db
	ctrl.Client = redis.NewClient(&redis.Options{
		Addr:     ctrl.Host + ":" + ctrl.Port,
		Password: ctrl.Password,
		DB:       ctrl.DB,
	})
	return &ctrl
}


// ping test database
func (ctrl *Ctrl) Ping() (string, error) {
	pong, err := ctrl.Client.Ping(context.Background()).Result()
	return pong, err
}

// get record from key
func (ctrl *Ctrl) Get(key string) (string, error) {
	val, err := ctrl.Client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// set new record key,value
func (ctrl *Ctrl) Set(key, value string) {
	err := ctrl.Client.Set(context.Background(), key, interface{}(value), 0).Err()
	if err != nil {
		panic(err)
	}
}

// get key, increment value, save in database, return new value
func (ctrl *Ctrl) IncrementingValues(key, value string) (string, error) {
	val, err := ctrl.Get(key)
	if err != nil {
		return "", err
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		return "", err
	}
	v2, err := strconv.Atoi(value)
	if err != nil {
		return "", err
	}
	r := v + v2
	res := strconv.Itoa(r)
	ctrl.Set(key, res)
	return res, nil
}
