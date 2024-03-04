package singleton

import "sync"

var tml *taskManagerLazy
var lock sync.Mutex

type taskManagerLazy struct {
}

func (t *taskManagerLazy) displayProcesses() {

}

func (t *taskManagerLazy) displayServices() {

}

// GetTaskManagerLazyInstance 懒汉式，通过doubleCheck方式
func GetTaskManagerLazyInstance() *taskManagerLazy {
	if tml == nil {
		lock.Lock()
		defer lock.Unlock()
		if tml == nil { // 在锁里面，doubleCheck下。
			tml = &taskManagerLazy{}
		}
	}
	return tml
}
