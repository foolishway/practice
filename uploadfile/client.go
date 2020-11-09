package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const serverAddr = ":51081"

func main() {
	uploadFile, err := os.Open("./testData/logs")
	defer uploadFile.Close()
	if err != nil {
		log.Fatalf("Open file error %v", err)
	}

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Dial %s error %v", serverAddr, err)
	}

	fLine := fmt.Sprintf("UPLOAD_FILE logs\n\r")
	conn.Write([]byte(fLine))
	io.Copy(conn, uploadFile)
}
