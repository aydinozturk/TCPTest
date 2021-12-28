package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {

		}
	}(l)

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		file, err := os.Open("data.txt")

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var gpsReadLines []string

		for scanner.Scan() {
			gpsReadLines = append(gpsReadLines, scanner.Text())
		}

		err = file.Close()
		if err != nil {
			return
		}

		for _, readLine := range gpsReadLines {
			_, err := c.Write([]byte(readLine + "\n"))
			if err != nil {
				return
			}

			time.Sleep(100 * time.Millisecond)
		}

	}
}
