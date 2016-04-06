package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var Builddate = "No build date provided"
var Githash = "No githasg provided"

func usage(name string) {
	fmt.Printf("Usage: %s [unix|tcp]:[/path/to/socket|ip:port]\n", os.Args[0])
	fmt.Printf("Build date:%s\n", Builddate)
	fmt.Printf("Githash:%s\n", Githash)
}

func main() {

	version := flag.Bool("version", false, "Print current version")
	flag.Parse()
	if *version {
		fmt.Printf("Build date:%s\n", Builddate)
		fmt.Printf("Githash:%s\n", Githash)
		os.Exit(0)
	}

	if len(os.Args) != 2 {
		usage(os.Args[0])
		return

	}
	args := strings.SplitN(os.Args[1], ":", 2)
	if len(args) != 2 {
		usage(os.Args[0])
		return

	}
	log.Println("=> Dialing", os.Args[1])
	conn, err := net.Dial(args[0], args[1])
	if err != nil {
		panic(err)

	}
	defer conn.Close()

	buf := make([]byte, 4096)
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))
		if err != nil {
			log.Fatal("=> write error:", err)

		}
		log.Println("=> Reading result")
		status, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return

		}
		log.Printf("==> Read %v bytes", status)
		log.Println("==> Result:")
		fmt.Printf("%v", string(buf[:]))

	}

}
