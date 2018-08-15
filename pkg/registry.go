package pkg

import (
	"github.com/piotrjaromin/map-processor/pkg/sink"
	"github.com/piotrjaromin/map-processor/pkg/source"
)

// Source only returns data and start pipeline
type Source interface {
	Get() (map[string]string, error)
}

// Sink takes in data and returns it as output
// Many sinks can be connected in pipeline
type Sink interface {
	Source
	Fill(map[string]string) error
}

// Registry contains names of sinks and sources along with functions that construct them
type Registry struct {
	Sources map[string]interface{}
	Sinks   map[string]interface{}
}

func (r *Registry) registerDefaultSources() {
	r.registerSource("ssm", source.NewSSM)
	r.registerSource("fetchTasks", source.NewTaskFetcher)
}

func (r *Registry) registerDefaultSinks() {
	r.registerSink("print", sink.NewPrint)
	r.registerSink("printOnlyVals", sink.NewPrintOnlyVals)
	r.registerSink("printNginx", sink.NewPrintNginx)
}

func (r *Registry) registerSink(name string, constructorFn interface{}) {
	r.Sinks[name] = constructorFn
}

func (r *Registry) registerSource(name string, constructorFn interface{}) {
	r.Sources[name] = constructorFn
}

// NewDefaultRegistry creates registry with default sinks and sources
func NewDefaultRegistry() Registry {
	registry := NewEmptyRegistry()
	registry.registerDefaultSinks()
	registry.registerDefaultSources()
	return registry
}

// NewEmptyRegistry creates new empty registry
func NewEmptyRegistry() Registry {
	return Registry{
		Sources: map[string]interface{}{},
		Sinks:   map[string]interface{}{},
	}
}
