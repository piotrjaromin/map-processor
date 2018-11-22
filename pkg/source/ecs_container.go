package source

import (
	"github.com/piotrjaromin/map-processor/pkg/awsutil"
	"strconv"
)

type ContainerFetcher struct {
	ServiceName *string
	ClusterName *string
	Region      *string
}

func NewContainerFetcher(region string, clusterName string, serviceName string) (ContainerFetcher, error) {

	return ContainerFetcher{
		ServiceName: &serviceName,
		ClusterName: &clusterName,
		Region:      &region,
	}, nil

}

func (tf ContainerFetcher) Get() (map[string]string, error) {

	result := map[string]string{}

	out, err := awsutil.GetEcsTasks(tf.Region, tf.ClusterName, tf.ServiceName)
	if err != nil {
		return result, err
	}

	for _, t := range out.Tasks {

		for _, c := range t.Containers {
			if len(c.NetworkInterfaces) > 0 {
				result[*c.ContainerArn] = *c.NetworkInterfaces[0].PrivateIpv4Address
			} else {
				result[*c.ContainerArn] = strconv.Itoa(int(*c.NetworkBindings[0].HostPort))
			}
		}
	}

	return result, nil
}
