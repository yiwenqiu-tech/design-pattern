package singleton

var tm *taskManager

type taskManager struct {
}

func (t *taskManager) displayProcesses() {

}

func (t *taskManager) displayServices() {

}

// GetTaskManagerInstance 可能会有并发问题，多个客户端同时访问时，可能会返回多个不同实例
func GetTaskManagerInstance() *taskManager {
	if tm == nil {
		tm = &taskManager{}
	}
	return tm
}
