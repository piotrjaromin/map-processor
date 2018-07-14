package sink

import (
	"log"
)

type PrintSink struct {
	params map[string]string
}

// NewPrint creates sink which prints data with which it was filled
func NewPrint() (*PrintSink, error) {
	return &PrintSink{}, nil
}

func (ps *PrintSink) Fill(params map[string]string) error {
	log.Println("Printing data")
	for key, val := range params {
		log.Printf("%s: %s", key, val)
	}

	return nil
}

func (ps *PrintSink) Get() (map[string]string, error) {
	return ps.params, nil
}
