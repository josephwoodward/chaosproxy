package main

import (
	"chaoskoko/agent"
	"chaoskoko/config"
	"flag"
)

func main() {
	flag.Parse()

	configLocation := flag.String("o", "./options.yml", "Location of config file")
	chaoskoko.Proxy(config.CommandLineArgs{ConfigLocation: *configLocation})
}
