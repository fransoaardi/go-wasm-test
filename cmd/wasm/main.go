package main

import (
	"context"
	"io"
	"net/http"
	"sync"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

func goHTTP(this js.Value, args []js.Value) interface{} {
	c := http.Client{}

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8000/", nil)
	if err != nil {
		println(err)
	}

	var rtn string

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		resp, err := c.Do(req)
		if err != nil {
			println(err)
		}
		defer resp.Body.Close()

		read, err := io.ReadAll(resp.Body)
		if err != nil {
			println(err)
		}

		rtn = string(read)
	}()
	wg.Wait()

	return js.ValueOf(rtn)
}

func goSumTo(this js.Value, args []js.Value) interface{} {
	sum := 0
	for i := 0; i < args[0].Int(); i++ {
		sum += i
	}
	return js.ValueOf(sum)
}

func registerCallbacks() {
	js.Global().Set("goSumTo", js.FuncOf(goSumTo))
	js.Global().Set("goHTTP", js.FuncOf(goHTTP))
}
