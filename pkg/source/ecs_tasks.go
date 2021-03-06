package source

import (
	"github.com/piotrjaromin/map-processor/pkg/awsutil"
)

type TaskFetcher struct {
	ServiceName *string
	ClusterName *string
	Region      *string
}

func NewTaskFetcher(region string, clusterName string, serviceName string) (TaskFetcher, error) {

	return TaskFetcher{
		ServiceName: &serviceName,
		ClusterName: &clusterName,
		Region:      &region,
	}, nil

}

func (tf TaskFetcher) Get() (map[string]string, error) {

	result := map[string]string{}

	out, err := awsutil.GetEcsTasks(tf.Region, tf.ClusterName, tf.ServiceName)
	if err != nil {
		return result, err
	}

	for _, t := range out.Tasks {
		result[*t.TaskArn] = *t.ContainerInstanceArn
	}

	return result, nil
}
