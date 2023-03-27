package visitor

// 和stateMachine不同，现在是已知所有资源的状态，然后更新应用和部署任务的状态
// 这种情况适合使用观察者模式
// 观察者模式允许一个对象（观察者）自动更新其状态，以反映另一个对象（被观察者）的状态变化。
// 应用可以作为资源的观察者，而部署任务可以作为应用的观察者

// 实现：
// Observer 接口包含一个 Update 方法。Application 和 DeploymentTask 结构体实现了这个接口，以便在资源状态发生变化时更新它们的状态。
// Resource 结构体现在包含一个观察者列表，并提供了 Attach 方法来注册观察者。它还提供了 Notify 方法，该方法在资源状态发生更改时通知所有观察者。
// Application 和 DeploymentTask 的 Update 方法分别用于更新应用和部署任务的状态。当资源状态更改时，这些方法会根据资源的新状态以及其他资源和应用的状态来计算新的应用和部署任务状态。

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
