package main

import (
	"gopher-byakugan/portscan"
	"log"
)

type urlInfo struct {
	url     string
	content int
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Staring the web server....")

	portscan.PortScan()

}
