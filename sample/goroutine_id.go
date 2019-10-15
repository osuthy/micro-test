package main

import (
    "bytes"
    "fmt"
    "runtime"
    "strconv"
		"sync"
)

func main() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			fmt.Println(getGID())
			wg.Done()
		}()
		wg.Wait()
    fmt.Println(getGID())
}

func getGID() uint64 {
    buff := make([]byte, 10000)
		runtime.Stack(buff, false)
    buff = bytes.TrimPrefix(buff, []byte("goroutine "))
    buff = buff[:bytes.IndexByte(buff, ' ')]
    id, _ := strconv.ParseUint(string(buff), 10, 64)
    return id
}
