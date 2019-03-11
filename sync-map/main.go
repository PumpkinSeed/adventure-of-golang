package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type logger struct {
	Fields sync.Map
}

func main() {
	mes := time.Now()
	l := &logger{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			l.AddField("test", "data")
			l.Log("test_msg")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(mes))
}

func (l *logger) AddField(k string, value interface{}) {
	l.Fields.Store(k, value)
}

func (l *logger) Log(msg string) {
	l.Fields.Store("msg", msg)
	l.Fields.Store("time", time.Now().UnixNano())

	out := make(map[string]interface{})

	l.Fields.Range(func(key, value interface{}) bool {
		if v, ok := key.(string); ok {
			out[v] = value
		}
		return true
	})

	serialized, err := json.Marshal(out)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println(string(serialized))
}
