package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"testOnce/utils"
)

const(
	SocketPath = "/host/etc/csi-tool/connector.sock"
)

func main(){
	if utils.IsFileExisting(SocketPath) {
		os.Remove(SocketPath)
	} else {
		pathDir := filepath.Dir(SocketPath)
		if !utils.IsFileExisting(pathDir) {
			os.MkdirAll(pathDir, os.ModePerm)
		}
	}
	fmt.Printf("Socket path is ready: %s\n", SocketPath)
	ln, err := net.Listen("unix", SocketPath)
	if err != nil {
		log.Fatalf("Server Listen error: %s", err.Error())
	}
	log.Print("Daemon Started ...")
	defer ln.Close()

	// Handler to process the command
	for {
		fd, err := ln.Accept()
		if err != nil {
			log.Printf("Server Accept error: %s", err.Error())
			continue
		}
		go echoServer(fd)
	}

}

func echoServer(c net.Conn) {
	buf := make([]byte, 2048)
	nr, err := c.Read(buf)
	if err != nil {
		log.Print("Server Read error: ", err.Error())
		return
	}

	cmd := string(buf[0:nr])
	log.Printf("Server Receive OSS command: %s", cmd)

	// Run command
	if out, err := utils.Run(cmd); err != nil {
		reply := "Fail: " + cmd + ", error: " + err.Error()
		_, err = c.Write([]byte(reply))
		log.Print("Server Fail to Run cmd:", reply)
	} else {
		out = "Success:" + out
		_, err = c.Write([]byte(out))
		log.Printf("Success: %s", out)
	}
}