// This is client file tcp client,
// Server program is tcpEchoServer.go
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println("IO Copy Failed!")
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s Port", os.Args[0])
	}
	serviceAddr := "localhost" + ":" + os.Args[1]
	conn, err := net.Dial("tcp", serviceAddr)
	if err != nil {
		log.Fatalf("Connect Failed! port:%s", os.Args[1])
	}

	ch := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Recv Server Data Finish!")
		ch <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-ch
}
