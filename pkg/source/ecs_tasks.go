package source

import (
	"github.com/piotrjaromin/map-processor/pkg/source/awsutil"
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

		for _, c := range t.Containers {
			result[*c.ContainerArn] = *c.NetworkInterfaces[0].PrivateIpv4Address
		}
	}

	return result, nil
}
