package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Err:", err, "StrErr:", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 6 {
		fmt.Printf("Usage:%s serverIP startPort endPort icmpPacketInterval portTestInterval\n", os.Args[0])
		os.Exit(0)
	}

	startPort, err := strconv.Atoi(os.Args[2])
	checkErr(err)

	endPort, err := strconv.Atoi(os.Args[3])
	checkErr(err)

	sendPacketInteval, err := strconv.Atoi(os.Args[4])
	checkErr(err)

	portTestInterval, err := strconv.Atoi(os.Args[5])
	checkErr(err)

	if startPort < 10000 || endPort > 50000 {
		fmt.Printf("Port Range is 10000 ~ 50000\n")
		os.Exit(0)
	}

	fmt.Println("ICMP Port Start Test! StartPort:", startPort, "EndPort:", endPort,
		"ICMPInterval:", sendPacketInteval, "PortInterval:", portTestInterval)

	if startPort%2 != 0 {
		startPort += 1
	}

	startTime := time.Now()
	startStrPort := strconv.Itoa(startPort)
	for {
		servAddr := os.Args[1] + ":" + startStrPort
		conn, err := net.Dial("udp", servAddr)
		checkErr(err)

		fmt.Println("Start Test Port:", startStrPort)
		for {
			testByte := "1"
			_, err = conn.Write([]byte(testByte))
			if err != nil {
				conn.Close()
				break
			}

			time.Sleep(time.Second * time.Duration(sendPacketInteval))
			TestInterval := time.Now().Unix() - startTime.Unix()
			if TestInterval >= int64(portTestInterval) {
				fmt.Printf("Port (%d) Test Finish!\n", startPort)
				startTime = time.Now()
				startPort = startPort + 2
				startStrPort = strconv.Itoa(startPort)
				conn.Close()
				fmt.Println("==================================")
				break
			}
		}

		if startPort > endPort {
			fmt.Println("All Port Test Finish...")
			break
		}
	}
}
