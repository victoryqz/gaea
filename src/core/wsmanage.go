package core

import (
	"fmt"
	"strings"
	"sync"
)

type WsManage struct {
	*sync.Map
}

var wsManageInstance *WsManage
var wsManageOnce sync.Once

func init() {
	wsManageOnce.Do(func() {
		wsManageInstance = &WsManage{
			Map: &sync.Map{},
		}
	})
}

func GetWsManage() *WsManage {
	return wsManageInstance
}

func (wm *WsManage) Join(wc *WsConnection, callback func(wm *WsMessage) interface{}) *WsManage {
	wm.Store(wc.id, wc)
	go wc.receive(func(id string) {
		if c, ok := wm.LoadAndDelete(id); ok {
			c.(*WsConnection).Close()
		}
	})
	go wc.loop(callback)
	go wc.send()

	return wm
}

func (wm *WsManage) Notice(namespace string, resource string, action WsAction, entity interface{}) {
	key := namespace + "-" + resource

	if temp, ok := wm.Load(key); ok {
		for _, value := range temp.([]string) {
			fmt.Println(value)
			if wsc, ok := wm.Load(value); ok {
				wsc.(*WsConnection).WriteJSON(entity)
			}
		}
	}
}

func (wm *WsManage) Subscribe(namespace string, resource string, id string) {
	key := strings.ToLower(namespace + "-" + resource)
	if temp, ok := wm.Load(key); ok {
		temp = append(temp.([]string), id)
		wm.Store(key, temp)
	} else {
		wm.Store(key, []string{id})
	}
}
