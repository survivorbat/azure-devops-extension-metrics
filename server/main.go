package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/survivorbat/azure-devops-extension-metrics/listener/message"
	"log"
	"net"
	"os"
	"time"
)

// checkError turns the if-statement into a one-liner to exit the program
func checkError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// processor is a one-time initialization of the processor
var processor = getProcessor(os.Getenv("PROCESSOR_ID"))

func main() {
	serveUrl := os.Getenv("LISTENER_URL")
	if serveUrl == "" {
		log.Fatalln("LISTENER_URL environment variable is not set")
	}

	listener, err := net.Listen("tcp", serveUrl)
	checkError(err)

	log.Printf("Listening on %s\n", serveUrl)
	for {
		if conn, err := listener.Accept(); err == nil {
			log.Printf("Accepted connection from %s\n", conn.RemoteAddr())
			go handleConnection(conn)
		}
	}
}

func handleConnection(con net.Conn) {
	defer con.Close()

	data := make([]byte, 2048)
	size, err := con.Read(data)
	if err != nil {
		log.Printf("Error reading data from connection: %s", err.Error())
	}

	var input message.Input
	if err := proto.Unmarshal(data[:size], &input); err != nil {
		log.Printf("Error unmarshalling data: %s", err.Error())
		return
	}

	id, err := uuid.Parse(input.GetId())
	if err != nil {
		log.Printf("Error parsing UUID: %s", err.Error())
		return
	}

	processor.Process(&ProcessInput{Timestamp: time.Unix(input.Timestamp, 0), ID: id})
}
