package sink

import (
	"fmt"
)

type PrintSink struct {
	params    map[string]string
	onlyVales bool
}

// NewPrint creates sink which prints data with which it was filled
func NewPrint() (*PrintSink, error) {
	return &PrintSink{onlyVales: false}, nil
}

func NewPrintOnlyVals() (*PrintSink, error) {
	return &PrintSink{onlyVales: true}, nil
}

func (ps *PrintSink) Fill(params map[string]string) error {
	fmt.Println("Printing data")

	for key, val := range params {
		if ps.onlyVales {
			fmt.Printf("\t\tserver %s:8000;\n", val)
		} else {
			fmt.Printf("%s: %s", key, val)
		}
	}

	return nil
}

func (ps *PrintSink) Get() (map[string]string, error) {
	return ps.params, nil
}
