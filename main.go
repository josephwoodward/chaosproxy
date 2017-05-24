package main

import (
	"chaos-kitten/agent"
	"flag"
)

func main() {
	flag.Parse()

	configLocation := flag.String("o", "./options.yml", "Location of config file")
	chaoskitten.Proxy(*configLocation)
}
