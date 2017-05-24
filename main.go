package main

import (
	"HttpMutt/agent"
	"flag"
)

func main() {
	flag.Parse()

	configLocation := flag.String("o", "./options.yml", "Location of config file")
	agent.Proxy(*configLocation)
}
