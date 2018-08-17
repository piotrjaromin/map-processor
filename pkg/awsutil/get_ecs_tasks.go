package awsutil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"

	"fmt"
)

func GetSession(region *string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: region,
	})

	if err != nil {
		return nil, fmt.Errorf("Could not create session %s", err.Error())
	}

	return sess, nil
}

func GetEcsTasks(region, clusterName, serviceName *string) (*ecs.DescribeTasksOutput, error) {

	sess, err := GetSession(region)
	if err != nil {
		return nil, err
	}

	svc := ecs.New(sess)

	tasksARNs, err := getTasks(svc, clusterName, serviceName)
	if err != nil {
		return nil, err
	}

	input := ecs.DescribeTasksInput{
		Cluster: clusterName,
		Tasks:   tasksARNs,
	}

	out, err := svc.DescribeTasks(&input)
	return out, err
}

func getTasks(svc *ecs.ECS, clusterName, serviceName *string) ([]*string, error) {

	input := ecs.ListTasksInput{
		ServiceName: serviceName,
		Cluster:     clusterName,
	}

	out, err := svc.ListTasks(&input)
	if err != nil {
		return make([]*string, 0), err
	}

	return out.TaskArns, nil
}

func StopTask(svc *ecs.ECS, clusterName, taskName *string) error {

	input := ecs.StopTaskInput{
		Task:    taskName,
		Cluster: clusterName,
	}

	_, err := svc.StopTask(&input)
	if err != nil {
		return err
	}

	return nil
}
