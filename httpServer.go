package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const listen_host = "0.0.0.0:12345"

func main() {
	fmt.Printf("Open %s\n", listen_host)
	// open 12345 port
	listener, error := net.Listen("tcp", listen_host)
	if error != nil {
		panic(error)
	}

	// この関数の終了時にCloseを実行
	defer listener.Close()

	for {
		connection, error := listener.Accept()
		if error != nil {
			panic(error)
		}
		defer connection.Close()
		buffer := make([]byte, 1500)
		go func() {
			for {
				n, error := connection.Read(buffer)
				if error != nil {
					if error == io.EOF {
						break
					}
					panic(error)
				}
				os.Stdout.Write(buffer[:n]) // echo
				_, error = connection.Write(buffer[:n])
				if error != nil {
					break
				}
			}
		}()
	}
}
