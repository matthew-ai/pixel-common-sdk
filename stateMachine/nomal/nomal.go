package nomal

// 不使用设计模式，而是遍历所有资源然后更新应用和任务的状态

type State string

const (
	Pending   State = "pending"
	Deploying State = "deploying"
	Success   State = "success"
	Failed    State = "failed"
)

type Observer interface {
	Update(state State)
}

type Resource struct {
	state     State
	observers []Observer
}

func (res *Resource) Attach(observer Observer) {
	res.observers = append(res.observers, observer)
}

func (res *Resource) Notify() {
	for _, observer := range res.observers {
		observer.Update(res.state)
	}
}

func (res *Resource) SetState(state State) {
	res.state = state
	res.Notify()
}

type Application struct {
	state     State
	resources []*Resource
	observers []Observer
}

func (app *Application) Attach(observer Observer) {
	app.observers = append(app.observers, observer)
}

func (app *Application) Notify() {
	for _, observer := range app.observers {
		observer.Update(app.state)
	}
}

func (app *Application) Update(resourceState State) {
	if app.state == Success || app.state == Failed {
		return
	}

	if resourceState == Failed {
		app.state = Failed
	} else {
		allResourcesSuccess := true
		for _, res := range app.resources {
			if res.state != Success {
				allResourcesSuccess = false
				break
			}
		}

		if allResourcesSuccess {
			app.state = Success
		} else {
			app.state = Deploying
		}
	}
	app.Notify()
}

type DeploymentTask struct {
	state        State
	applications []*Application
}

func (dt *DeploymentTask) Update(appState State) {
	if dt.state == Success || dt.state == Failed {
		return
	}

	if appState == Failed {
		dt.state = Failed
	} else {
		allAppsSuccess := true
		for _, app := range dt.applications {
			if app.state != Success {
				allAppsSuccess = false
				break
			}
		}

		if allAppsSuccess {
			dt.state = Success
		} else {
			dt.state = Deploying
		}
	}
}
