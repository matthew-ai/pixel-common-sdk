package state

import (
	"fmt"
	"testing"
)

func TestDeploymentTask_updateTaskState(t *testing.T) {
	task := DeploymentTask{
		state: Pending,
		applications: []Application{
			{
				state: Pending,
				resources: []Resource{
					{state: Pending},
					{state: Pending},
				},
			},
			{
				state: Pending,
				resources: []Resource{
					{state: Pending},
				},
			},
		},
	}

	task.updateResourceState(0, 0, Deploying)
	fmt.Println(task.state) // deploying

	task.updateResourceState(0, 1, Success)
	fmt.Println(task.state) // deploying

	task.updateResourceState(0, 0, Success)
	fmt.Println(task.state) // deploying

	task.updateResourceState(1, 0, Success)
	fmt.Println(task.state) // success

	task.updateResourceState(0, 1, Failed)
	fmt.Println(task.state) // success (不会变为failed，因为任务状态已经是success)
}
