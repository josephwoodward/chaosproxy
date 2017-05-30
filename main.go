package main

import (
	"chaosproxy/agent"
	"chaosproxy/config"
	"flag"
)

func main() {
	flag.Parse()

	configLocation := flag.String("o", "./options.yml", "Location of config file")
	chaosproxy.Proxy(config.CommandLineArgs{ConfigLocation: *configLocation})
}
