package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func echoTime(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			loggerPrint("TCP Conn Echo Msg Failed, Maybe Client close!")
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func echoInput(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		log.Println("Recv Data:", input.Text())
		_, err := io.WriteString(conn, input.Text())
		if err != nil {
			loggerPrint("Maybe Client Close.")
			break
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s Port", os.Args[0])
		os.Exit(0)
	}

	serviceAddr := "localhost" + ":" + os.Args[1]
	listener, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		log.Println("tcp listen failed!")
		os.Exit(0)
	}

	for {
		connect, err := listener.Accept()
		if err != nil {
			loggerPrint("Accept TCP Connect Failed!")
			continue
		}

		//go echoTime(connect)
		go echoInput(connect)
	}
}
