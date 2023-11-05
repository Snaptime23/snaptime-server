package main

import (
	baseHttpInterface "github.com/Snaptime23/snaptime-server/v2/base/interface/cmd"
	baseService "github.com/Snaptime23/snaptime-server/v2/base/service/cmd"

	videoService "github.com/Snaptime23/snaptime-server/v2/video/service/cmd"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		baseService.Run()
	}()
	go func() {
		defer wg.Done()
		videoService.Run()
	}()

	time.Sleep(time.Second * 5)
	go func() {
		defer wg.Done()
		baseHttpInterface.Run()
	}()

	wg.Wait()
}
