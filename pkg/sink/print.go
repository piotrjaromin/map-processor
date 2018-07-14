package sink

import (
	"fmt"
)

type PrintSink struct {
	params map[string]string
}

func NewPrint() (*PrintSink, error) {
	return &PrintSink{}, nil
}

func (ps *PrintSink) Fill(params map[string]string) error {

	for key, val := range params {
		fmt.Printf("%s: %s", key, val)
	}

	return nil
}

func (ps *PrintSink) Get() (map[string]string, error) {
	return ps.params, nil
}
