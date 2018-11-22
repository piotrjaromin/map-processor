package sink

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/piotrjaromin/map-processor/pkg/awsutil"
)

type TaskKiller struct {
	ClusterName *string
	Region      *string
	params      map[string]string
}

func NewTaskKiller(region string, clusterName string) (*TaskKiller, error) {

	return &TaskKiller{
		ClusterName: &clusterName,
		Region:      &region,
	}, nil

}

func (ts *TaskKiller) Fill(params map[string]string) error {
	result := map[string]string{}

	sess, err := awsutil.GetSession(ts.Region)
	if err != nil {
		return err
	}

	svc := ecs.New(sess)

	for taskArn := range params {

		if err := awsutil.StopTask(svc, ts.ClusterName, &taskArn); err != nil {
			result[taskArn] = fmt.Sprintf("FAILED to kill: %s", err.Error())
		} else {
			result[taskArn] = "KILLED"
		}
	}

	ts.params = result
	return nil
}

func (ts *TaskKiller) Get() (map[string]string, error) {
	return ts.params, nil
}
