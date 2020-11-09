package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Listen error %v", err)
	}

	var bf *bufio.Reader
	addr := l.Addr().String()
	log.Printf("Server listen at %s", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Accept error %v", err)
		}

		bf = bufio.NewReader(conn)
		newName := strconv.FormatInt(time.Now().Unix(), 10)

		line, _, err := bf.ReadLine()
		if err != nil {
			log.Fatalf("Read first line error %v", err)
		}

		s := strings.SplitN(string(line), " ", 2)
		flag, fileName := s[0], s[1]
		if flag != "UPLOAD_FILE" {
			continue
		}
		newName = newName + "_" + fileName

		uploadPath := fmt.Sprintf("./upload_files/%s", newName)

		dir := filepath.Dir(uploadPath)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("Create %s error %v", dir, err)
		}

		f, err := os.Create(uploadPath)
		if err != nil {
			log.Fatalf("Create %s error %v", uploadPath, err)
		}

		io.Copy(f, conn)
		f.Close()
	}
}
