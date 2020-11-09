// go1.15-examples/runtime/tinyint2interface.go

package main

import (
	"io"
	"log"
	"os"
	"time"
)

type Observer struct {
}

func (o *Observer) Write(b []byte) (int, error) {
	log.Printf("%s: %s", time.Now(), string(b))
	return len(b), nil
}

type ObserverWapper struct {
	w io.Writer
}

func (ob *ObserverWapper) Write(b []byte) (int, error) {
	return ob.w.Write(b)
}

func newObserverWapper(des io.Writer, ob ...io.Writer) ObserverWapper {
	writes := make([]io.Writer, 0)
	writes = append([]io.Writer{des})
	writes = append(writes, ob...)
	return ObserverWapper{w: io.MultiWriter(writes...)}
}

func main() {
	f, err := os.OpenFile("./test.yml", os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Open file error %v", err)
	}

	ob := &Observer{}
	//mul ob...
	obWapper := newObserverWapper(f, ob)

	obWapper.Write([]byte("hello world"))

}
