package main

import (
	"context"
	"fmt"
	"os/exec"
	"sync"

	"github.com/go-redis/redis"
)

var COUNTTEST = 100

func create_client() *redis.Client {
	var client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	var key = "keytest"
	var value = "0"
	err := client.Set(context.Background(), key, interface{}(value), 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
	return client
}

func LineGet(line string) string {
	re := exec.Command("bash", "-c", line)
	res, _ := re.Output()
	return string(res)
}

func post_test1() {
	r := LineGet("./curl_test1.sh keytest 1")
	fmt.Println(r)
}

func post_test2() {
	r := LineGet("./curl_test2.sh testline keysecret")
	fmt.Println(r)
}

func post_test3() {
	r := LineGet(`./curl_test3.sh`)
	fmt.Println(r)
}

func main() {
	create_client()
	var wg sync.WaitGroup
	for i := 0; i < COUNTTEST; i++ {
		wg.Add(3)
		go func() {
			post_test1()
			wg.Done()
		}()
		go func() {
			post_test2()
			wg.Done()
		}()
		go func() {
			post_test3()
			wg.Done()
		}()
	}
	wg.Wait()
}
