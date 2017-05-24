package main

import (
	"HttpMutt/agent"
	"flag"
)

func main() {

	configLocation := flag.String("o", "./options.json", "Location of config file")
	portNumber := flag.String("p", "8080", "Port number to listen on")
	outputLog := flag.Bool("l", false, "Output log to file")

	flag.Parse()

	agent.Log(*configLocation, *portNumber, *outputLog)

	//HttpLog.Log(HttpLog.ConfigurationOptions{
	//	Port:          *portNumber,
	//	LogDir:        *logDir,
	//	RunAsDetached: false,
	//})

}
