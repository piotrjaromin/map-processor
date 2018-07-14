package source

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"log"
)

var optionRecursive = "Recursive"
var keyPath = "Path"
var trueVal = true

type ParamsFetcher struct {
	Svc            *ssm.SSM
	Path           *string
	WithDecryption *bool
}

func NewSSM(region *string, withDecryption *bool, path *string) ParamsFetcher {
	sess, err := session.NewSession(&aws.Config{
		Region: region,
	})

	if err != nil {
		log.Panicf("Could not create session %s", err.Error())
	}

	svc := ssm.New(sess)

	return ParamsFetcher{
		Svc:            svc,
		Path:           path,
		WithDecryption: withDecryption,
	}

}

func (pf ParamsFetcher) Get() (map[string]string, error) {
	return pf.FetchAllRecursive(nil, map[string]string{})
}

func (pf ParamsFetcher) FetchAllRecursive(nextToken *string, params map[string]string) (map[string]string, error) {

	queryPath := ssm.GetParametersByPathInput{
		Recursive:      &trueVal,
		Path:           pf.Path,
		WithDecryption: pf.WithDecryption,
	}

	if nextToken != nil {
		queryPath.NextToken = nextToken
	}

	res, err := pf.Svc.GetParametersByPath(&queryPath)
	if err != nil {
		log.Panicf("Could not create session %s", err.Error())
		return params, err
	}

	for _, param := range res.Parameters {
		if param != nil {
			fmt.Printf("%s: %s\n", *param.Name, *param.Value)
			params[*param.Name] = *param.Value
		} else {
			fmt.Println("Got nil parameter!")
		}
	}

	if res.NextToken != nil {
		return pf.FetchAllRecursive(res.NextToken, params)
	}

	return params, nil
}
