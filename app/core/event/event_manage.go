package event

// EventStoreList 定义一个全局事件存储变量
var EventStoreList = make(map[string]func(args ...interface{}), 0)

// CreateEventManageFactory 创建一个事件管理工厂
func CreateEventManageFactory() *EventManage {
	return &EventManage{}
}

// EventManage 定义一个事件管理结构体
type EventManage struct {
}

// Register 1.注册事件， 强烈建议注册事件的时候，根据不同的类型添加对应的前缀
func (e *EventManage) Register(keyName string, keyNameFunc func(args ...interface{})) {
	//判断keyName下是否已有事件
	if e.keyExistsEvent(keyName) == false {
		EventStoreList[keyName] = keyNameFunc
	}
}

// Delete 2.删除事件
func (e *EventManage) Delete(keyName string) {
	delete(EventStoreList, keyName)
}

// Dispatch 3.调用事件
func (e *EventManage) Dispatch(keyName string, args ...interface{}) {

	// 调用一个确定性的事件
	if len(keyName) > 0 {
		e.callEvent(keyName, args...)
	} else {
		// 触发已经注册的全部事件去执行
		for key, _ := range EventStoreList {
			e.callEvent(key, args...)
		}
	}
}

// 4 执行事件
func (e *EventManage) callEvent(keyName string, args ...interface{}) {
	if fn, ok := EventStoreList[keyName]; ok {
		fn(args...)
	}
}

// 判断某个键是否已经存在某个事件
func (e *EventManage) keyExistsEvent(keyName string) bool {
	if _, exists := EventStoreList[keyName]; exists {
		return exists
	}
	return false
}
