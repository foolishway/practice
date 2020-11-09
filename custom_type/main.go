package main

import "fmt"

type Handler func(s string) error

func (h Handler) Log(s string) error {
	return h(s)
}

type Ihandler interface {
	log() error
}

func handler(s string) error {
	fmt.Println("handler log " + s)
	return nil
}

func main() {
	Handler(handler).Log("error info")
}
