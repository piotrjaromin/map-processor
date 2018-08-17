package main

import (
	"flag"

	"github.com/piotrjaromin/map-processor/pkg"
	"github.com/piotrjaromin/map-processor/pkg/config"
	"log"
)

func main() {

	confFile := flag.String("conf", "./pipe.yml", "Path to pipe.yml file")
	flag.Parse()

	log.Printf("using conifg %s", *confFile)
	conf, err := config.Read(*confFile)
	if err != nil {
		log.Panicf("unable to read config file in path %s, reason: %s", *confFile, err.Error())
	}

	registry := pkg.NewDefaultRegistry()
	if err := pkg.Process(conf, registry); err != nil {
		log.Panic("Could not process pipeline, reason: ", err.Error())
	}
}
