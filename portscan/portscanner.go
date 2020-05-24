package portscan

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sort"
)

//PortScan performs a scanning
func PortScan() {
	log.Println("Port Scanner...")

	startPort := flag.Int("startPort", 1, "Starting port to scan")
	endPort := flag.Int("endPort", 200, "End of port to scan")

	flag.Parse()

	if *startPort > *endPort {
		panic("Invalid range!")
	}

	log.Printf("Starting Port is %d\n", *startPort)
	log.Printf("Ending Port is %d\n", *endPort)

	var targetPorts []int
	var resultPorts []int
	for i := *startPort; i <= *endPort; i++ {
		//log.Println(i)
		targetPorts = append(targetPorts, i)
	}

	result := make(chan int)
	sniffPortWorker(result, targetPorts)

	for i := *startPort; i < *endPort; i++ {
		outResult := <-result
		if outResult != 0 {
			resultPorts = append(resultPorts, outResult)
		}
	}
	log.Println("Sorting result!")
	sort.Ints(resultPorts)
	for _, port := range resultPorts {
		log.Printf("Open ports is %d", port)
	}

}

func sniffPortWorker(result chan<- int, ports []int) {

	for i := range ports {
		go func(port int, out chan<- int) {
			target := fmt.Sprintf("scanme.nmap.org:%d", port)
			//log.Println("Scanning port of " + target)
			serv, err := net.Dial("tcp", target)
			if err != nil {
				out <- 0
				return
			}
			out <- port
			serv.Close()
		}(ports[i], result)
	}

}
