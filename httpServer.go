package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"./httprequest"
	"./httpresponse"
	"./getfile"
	"./extension"
	// "time"
)

const listen_host = "0.0.0.0:8080"
const roorDir = "./public"

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
		go func() {
			for {
				buffer := make([]byte, 1500)
				n, error := connection.Read(buffer)
				if error != nil {
					if error == io.EOF {
						break
					}
					panic(error)
				}
				isHttp, request := httprequest.GetHTTPRequest(buffer[:n])
				// fmt.Println(isHttp)
				// fmt.Printf("%+v\n", request)
				// os.Stdout.Write(buffer[:n])
				if isHttp {
					isKeepAlive := false
					contentBytes, err := getfile.GetFileFromPath(roorDir, request.Path)
					if err != nil {
						break
					}
					fileName := extension.GetFileNameFromPath(request.Path)
					contentType :=  extension.GetContentType(fileName)
					var hedders []string
					hedders = append(hedders, "Content-Type:" + contentType)
					if request.Connection == "keep-alive" {
						isKeepAlive = true
						hedders = append(hedders, "Connection: keep-alive")
					}
					res, err := httpresponse.CreateHttpResponse(request.Protocol, 200, hedders, contentBytes)
					if err != nil {
						break
					}
					_, error = connection.Write(res)

					if !isKeepAlive {
						connection.Close()
						break
					}

				} else {
					os.Stdout.Write([]byte("unknown"))
					_, error = connection.Write([]byte("unknown"))
				}
				_, error = connection.Write([]byte("\n"))
				if error != nil {
					break
				}
			}
		}()
	}
}
