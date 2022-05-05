// DO NOT MODIFY!

package main

import (
	"flag"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/cmu440/tribbler/libstore"

	"github.com/cmu440/tribbler/tribserver"
)

var port = flag.Int("port", 9010, "port number to listen on")
var master = flag.String("master", "", "master storage server host:port (REQUIRED)")

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
}

func main() {
	flag.Parse()
	if len(*master) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Create and start the TribServer.
	hostPort := net.JoinHostPort("localhost", strconv.Itoa(*port))
	_, err := tribserver.NewTribServer(*master, hostPort, libstore.Normal, true)
	if err != nil {
		log.Fatalln("Server could not be created:", err)
	}

	// Run the Tribbler server forever.
	select {}
}
