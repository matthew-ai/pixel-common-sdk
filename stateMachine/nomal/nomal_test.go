package nomal

import (
	"fmt"
	"testing"
)

func TestDeploymentTask_Update(t *testing.T) {
	task := DeploymentTask{state: Pending}
	app1 := Application{state: Pending}
	app2 := Application{state: Pending}

	task.applications = []*Application{&app1, &app2}

	res1 := Resource{state: Pending}
	res2 := Resource{state: Pending}
	res3 := Resource{state: Pending}

	app1.resources = []*Resource{&res1, &res2}
	app2.resources = []*Resource{&res3}

	res1.Attach(&app1)
	res2.Attach(&app1)
	res3.Attach(&app2)

	app1.Attach(&task)
	app2.Attach(&task)

	res1.SetState(Deploying)
	fmt.Println(task.state) // deploying

	res2.SetState(Success)
	fmt.Println(task.state) // deploying

	res1.SetState(Success)
	fmt.Println(task.state) // deploying

	res3.SetState(Success)
	fmt.Println(task.state) // success

	res2.SetState(Failed)
	fmt.Println(task.state) // success (不会变为failed，因为任务状态已经是success)
}
