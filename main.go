package main

import (
	"chaos-kitten/agent"
	"chaos-kitten/config"
	"flag"
)

func main() {
	flag.Parse()

	configLocation := flag.String("o", "./options.yml", "Location of config file")
	chaoskitten.Proxy(config.CommandLineArgs{ConfigLocation: *configLocation})
}
