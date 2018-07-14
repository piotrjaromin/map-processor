package pkg

import (
	"fmt"
	"reflect"
)

func getSource(name string, args []interface{}, registry Registry) (Source, error) {

	sourceVal, err := createFromReflect(name, args, registry.Sources)
	if err != nil {
		return nil, err
	}

	return sourceVal.Interface().(Source), nil
}

func getSink(name string, args []interface{}, registry Registry) (Sink, error) {

	sourceVal, err := createFromReflect(name, args, registry.Sinks)
	if err != nil {
		return nil, err
	}

	return sourceVal.Interface().(Sink), nil
}

// createFromReflect generic function which should create instance
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
