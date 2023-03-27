package state

// 状态模式允许一个对象在其内部状态改变时改变其行为。
// 部署任务、应用和资源可以根据状态有不同的行为

type State string

const (
	Pending   State = "pending"
	Deploying State = "deploying"
	Success   State = "success"
	Failed    State = "failed"
)

type Resource struct {
	state State
}

type Application struct {
	state     State
	resources []Resource
}

type DeploymentTask struct {
	state        State
	applications []Application
}

func (dt *DeploymentTask) updateTaskState() {
	if dt.state == Success || dt.state == Failed {
		return
	}

	allAppsSuccess := true
	for _, app := range dt.applications {
		if app.state == Failed {
			dt.state = Failed
			return
		}
		if app.state != Success {
			allAppsSuccess = false
		}
	}

	if allAppsSuccess {
		dt.state = Success
	} else {
		dt.state = Deploying
	}
}

func (app *Application) updateAppState() {
	if app.state == Success || app.state == Failed {
		return
	}

	allResourcesSuccess := true
	for _, res := range app.resources {
		if res.state == Failed {
			app.state = Failed
			return
		}
		if res.state != Success {
			allResourcesSuccess = false
		}
	}

	if allResourcesSuccess {
		app.state = Success
	} else {
		app.state = Deploying
	}
}

func (dt *DeploymentTask) updateResourceState(appIndex, resIndex int, newState State) {
	if dt.applications[appIndex].resources[resIndex].state == Success || dt.applications[appIndex].resources[resIndex].state == Failed {
		return
	}

	dt.applications[appIndex].resources[resIndex].state = newState
	dt.applications[appIndex].updateAppState()
	dt.updateTaskState()
}
