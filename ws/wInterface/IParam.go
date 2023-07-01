package wInterface

import . "github.com/Roninchen/OKEX_V5SDK_GO/ws/wImpl"

// 请求数据
type WSParam interface {
	EventType() Event
	ToMap() *map[string]string
}
