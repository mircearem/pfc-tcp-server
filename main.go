package main

import (
	"log"
	"os"

	"github.com/mircearem/pfc-tcp-server/network"
	"github.com/sirupsen/logrus"
)

// Set up logging
func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	s := network.NewTransport(":3000")
	if err := s.ListenAndAccept(); err != nil {
		log.Fatalln(err)
	}
}
