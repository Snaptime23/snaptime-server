package main

import (
	baseHttpInterface "github.com/Snaptime23/snaptime-server/v2/base/interface/cmd"
	baseService "github.com/Snaptime23/snaptime-server/v2/base/service/cmd"

	videoHttpInterface "github.com/Snaptime23/snaptime-server/v2/video/interface/cmd"
	videoService "github.com/Snaptime23/snaptime-server/v2/video/service/cmd"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		baseService.Run()
	}()
	time.Sleep(time.Second * 5)
	go func() {
		defer wg.Done()
		baseHttpInterface.Run()
	}()

	go func() {
		defer wg.Done()
		videoService.Run()
	}()
	time.Sleep(time.Second * 5)
	go func() {
		defer wg.Done()
		videoHttpInterface.Run()
	}()

	wg.Wait()
}
