package main

import (
	"log"

	"github.com/cswank/amp"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	cmd = kingpin.Arg("command", "on or off").Required().Enum("on", "off")
)

func main() {
	kingpin.Parse()

	a, err := amp.New()
	if err != nil {
		log.Fatal(err)
	}

	switch *cmd {
	case "on":
		err = a.On()
	case "off":
		err = a.Off()
	}

	if err != nil {
		log.Println(err)
	}

	a.Close()
}
