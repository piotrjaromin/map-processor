package source

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JsonReader struct {
	Path string
}

func NewJsonFileReader(filePath string) (JsonReader, error) {

	return JsonReader{
		Path: filePath,
	}, nil

}

func (jr JsonReader) Get() (map[string]string, error) {

	result := map[string]string{}

	jsonFile, err := os.Open(jr.Path)
	if err != nil {
		return result, err
	}
	defer jsonFile.Close()

	bytesJson, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return result, err
	}

	json.Unmarshal(bytesJson, &result)

	return result, nil
}
