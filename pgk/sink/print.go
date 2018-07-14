package sink

import (
	"fmt"
)

type PrintSink struct{}

func NewPrint() PrintSink {
	return PrintSink{}
}

func (ps PrintSink) Fill(params map[string]string) error {

	for key, val := range params {
		fmt.Printf("%s: %s", key, val)
	}

	return nil
}
