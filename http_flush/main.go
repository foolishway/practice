package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		// fmt.Println("flush")
		fw.f.Flush()
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}
	cmd := exec.Command("counter")
	cmd.Stdout = &fw
	cmd.Stderr = &fw
	cmd.Run()
}

func main() {
	fmt.Printf("%s, %-30s", "hello", "world")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
