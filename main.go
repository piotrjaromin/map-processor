package main

import (
	"flag"
	"github.com/piotrjaromin/param-store-utils/pkg/sink"
	"github.com/piotrjaromin/param-store-utils/pkg/source"
	"log"
)

type Source interface {
	Get() (map[string]string, error)
}

type Sink interface {
	Fill(map[string]string) error
}

type SourceSink interface {
	Source
	Sink
}

func main() {

	path := flag.String("path", "", "Path for which ssm parameters should be found")
	region := flag.String("region", "eu-central-1", "Region in which search for parameters should be done")
	withDecryption := flag.Bool("withDecryption", true, "If values should be decrypted")
	flag.Parse()

	src := source.NewSSM(region, withDecryption, path)
	params, err := src.Get()
	if err != nil {
		log.Panicf("Got error while fetching parameters %s", err.Error())
	}

	dst := sink.NewPrint()
	dst.Fill(params)

	if err != nil {
		log.Panicf("Got error while printing parameters %s", err.Error())
	}

}
