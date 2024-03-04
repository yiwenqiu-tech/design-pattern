package singleton

var tme = &taskManagerEager{} // 饿汉式：定义式就创建实例

type taskManagerEager struct {
}

func (t *taskManagerEager) displayProcesses() {

}

func (t *taskManagerEager) displayServices() {

}

// GetTaskManagerEagerInstance 饿汉式，不需要判断是否为空，因为定义时就创建实例，直接返回即可。
func GetTaskManagerEagerInstance() *taskManagerEager {
	return tme
}
