package main

import (
	"flag"

	"github.com/muslim-teachings/api-gateway/src/main/config"
)

var args string

func main() {
	flag.StringVar(&args, "config-files", "", "Used to parse config from url or path")
	flag.Parse()
	paths := flag.Args()

	for _, path := range paths {
		servs, err := config.ReadServerConfigFromYaml(path)
		if err == nil {
			for _, serv := range servs {
				go config.RegisterService(serv)
			}
		}
	}
	// forever loop
	for {
	}
}
