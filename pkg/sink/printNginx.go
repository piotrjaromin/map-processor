package sink

import (
	"fmt"
	"sort"
)

type PrintNginxSink struct {
	params    map[string]string
	onlyVales bool
}

// PrintNginx creates sink which prints data with which it was filled
func NewPrintNginx() (*PrintNginxSink, error) {
	return &PrintNginxSink{}, nil
}

func (ps *PrintNginxSink) Fill(params map[string]string) error {
	fmt.Println("Printing data")

	arr := []string{}
	for _, val := range params {
		arr = append(arr, val)
	}

	sort.Strings(arr)
	for _, val := range arr {
		fmt.Printf("\tserver %s:8000 max_fails=0;\n", val)
	}

	return nil
}

func (ps *PrintNginxSink) Get() (map[string]string, error) {
	return ps.params, nil
}
