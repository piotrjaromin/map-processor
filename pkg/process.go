package pkg

import (
	"fmt"
	"github.com/piotrjaromin/map-processor/pkg/config"

	"log"
)

// Process data based on pipeline defined in config
// Config has one source as starting point
// and multiple sinks which then pass data to each other
// Registry contains definitions of source and sinks
func Process(conf config.Config, registry Registry) error {

	var source Source
	for name, params := range conf.Source {
		log.Printf("Creating source with name: %s and params: %+v", name, params)
		src, err := getSource(name, params, registry)
		if err != nil {
			return fmt.Errorf("Unable to create source, reason: %s", err.Error())
		}

		source = src
	}

	var sinks = map[string]Sink{}
	for name, params := range conf.Sinks {
		log.Printf("Creating sink with name: %s and params: %+v", name, params)
		sink, err := getSink(name, params, registry)
		if err != nil {
			return fmt.Errorf("Unable to create sink, reason: %s", err.Error())
		}
		sinks[name] = sink
	}

	// Pipe data through sinks
	params, err := source.Get()
	if err != nil {
		return fmt.Errorf("Could not start pipeline, source error: %s", err.Error())
	}

	for name, sink := range sinks {
		if err := sink.Fill(params); err != nil {
			return fmt.Errorf("Sink: %s, on Fill returned error: %s", name, err.Error())
		}
		if params, err = sink.Get(); err != nil {
			return fmt.Errorf("Sink: %s, on Get returned error: %s", name, err.Error())
		}
	}

	return nil
}
