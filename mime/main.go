//MIME
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("pid:", os.Getpid())
		res.Header().Set("content-type", "application/octet-stream")
		res.Header().Set("Content-Disposition", "attachment; filename=logs")
		f, _ := os.Open("../logs")
		defer f.Close()
		io.Copy(res, f)
	})
	http.ListenAndServe(":3000", nil)
}
