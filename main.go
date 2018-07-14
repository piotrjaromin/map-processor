package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/piotrjaromin/param-store-utils/pkg/config"
	"github.com/piotrjaromin/param-store-utils/pkg/sink"
	"github.com/piotrjaromin/param-store-utils/pkg/source"
	"log"
)

type Source interface {
	Get() (map[string]string, error)
}

type Sink interface {
	Source
	Fill(map[string]string) error
}

type Registry struct {
	Sources map[string]interface{}
	Sinks   map[string]interface{}
}

func (r *Registry) registerSources() {
	r.Sources["ssm"] = source.NewSSM
}

func (r *Registry) registerSinks() {
	r.Sinks["print"] = sink.NewPrint
}

var registry = Registry{
	Sources: map[string]interface{}{},
	Sinks:   map[string]interface{}{},
}

func init() {
	registry.registerSinks()
	registry.registerSources()
}

func main() {

	confFile := flag.String("conf", "./pipe.yaml", "Path to pipe.yaml file")
	flag.Parse()

	conf, err := config.Read(*confFile)

	if err != nil {
		log.Panicf("unable to read config file in path %s, reason: %s", *confFile, err.Error())
	}

	log.Printf("Config is %+v", conf)

	var source Source
	for name, params := range conf.Source {
		log.Printf("Creating source with name: %s and params: %+v", name, params)
		source, err = getSource(name, params)
		if err != nil {
			log.Panicf("Unable to create source, reason: %s", err.Error())
		}
	}

	var sinks = map[string]Sink{}
	for name, params := range conf.Sinks {
		log.Printf("Creating sink with name: %s and params: %+v", name, params)
		sink, err := getSink(name, params)
		if err != nil {
			log.Panicf("Unable to create sink, reason: %s", err.Error())
		}
		sinks[name] = sink
	}

	// Pipe data through sinks
	params, err := source.Get()
	for name, sink := range sinks {
		if err := sink.Fill(params); err != nil {
			log.Panicf("Sink: %s, on Fill returned error: %s", name, err.Error())
		}
		if params, err = sink.Get(); err != nil {
			log.Panicf("Sink: %s, on Get returned error: %s", name, err.Error())
		}
	}

}

func getSource(name string, args []interface{}) (Source, error) {

	sourceVal, err := createFromReflect(name, args, registry.Sources)
	if err != nil {
		return nil, err
	}

	return sourceVal.Interface().(Source), nil
}

func getSink(name string, args []interface{}) (Sink, error) {

	sourceVal, err := createFromReflect(name, args, registry.Sinks)
	if err != nil {
		return nil, err
	}

	return sourceVal.Interface().(Sink), nil
}

func createFromReflect(name string, args []interface{}, initRegistry map[string]interface{}) (*reflect.Value, error) {

	sourceInit, ok := initRegistry[name]
	if !ok {
		return nil, fmt.Errorf("Inalid name: %s", name)
	}

	values := []reflect.Value{}
	for _, arg := range args {
		values = append(values, reflect.ValueOf(arg))
	}

	sourceVals := reflect.ValueOf(sourceInit).Call(values)
	if len(sourceVals) != 2 {
		return nil, fmt.Errorf("Unable to create from reflection, wrong number of arguments")
	}

	errorVal := sourceVals[1]
	if !errorVal.IsNil() {
		err := errorVal.Interface().(error)
		return nil, fmt.Errorf("Create from reflect returned error: %s", err.Error())
	}

	return &sourceVals[0], nil
}
