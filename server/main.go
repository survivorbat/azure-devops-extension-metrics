package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/survivorbat/azure-devops-extension-metrics/server/message"
	"net"
	"os"
	"time"
)

// checkError turns the if-statement into a one-liner to exit the program
func checkError(err error) {
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

// processor is a one-time initialization of the processor
var processor = getProcessor(os.Getenv("PROCESSOR_ID"))

func main() {
	serveUrl := os.Getenv("LISTENER_URL")
	if serveUrl == "" {
		logrus.Fatal("LISTENER_URL environment variable is not set")
	}

	listener, err := net.Listen("tcp", serveUrl)
	checkError(err)

	logrus.Infof("Listening on %s", serveUrl)
	for {
		if conn, err := listener.Accept(); err == nil {
			go handleConnection(conn)
		}
	}
}

func handleConnection(con net.Conn) {
	defer con.Close()

	data := make([]byte, 2048)
	size, err := con.Read(data)
	if err != nil {
		logrus.Debugf("Error reading data from connection: %s", err.Error())
	}

	var input message.Input
	if err := proto.Unmarshal(data[:size], &input); err != nil {
		logrus.Debugf("Error unmarshalling data: %s", err.Error())
		return
	}

	id, err := uuid.Parse(input.GetId())
	if err != nil {
		logrus.Infof("Error parsing UUID: %s", err.Error())
		return
	}

	processor.Process(&ProcessInput{Timestamp: time.Unix(input.Timestamp, 0), ID: id})
}
