package main

import (
	httpInterface "github.com/Snaptime23/snaptime-server/v2/base/interface/cmd"
	service "github.com/Snaptime23/snaptime-server/v2/base/service/cmd"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		service.Run()
	}()
	time.Sleep(time.Second * 5)
	go func() {
		defer wg.Done()
		httpInterface.Run()
	}()
	wg.Wait()
}
