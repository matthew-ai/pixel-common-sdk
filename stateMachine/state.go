package stateMachine

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
	dt.applications[appIndex].resources[resIndex].state = newState
	dt.applications[appIndex].updateAppState()
	dt.updateTaskState()
}
