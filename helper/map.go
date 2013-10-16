//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2013-10-17
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	sync mutex map
//
package SFHelper

import (
	"sync"
)

type Map struct {
	rwmutex sync.RWMutex
	thisMap map[interface{}]interface{}
}

func NewMap() Map {
	return Map{thisMap: make(map[interface{}]interface{})}
}

func (m *Map) Get(key interface{}) (interface{}, bool) {
	m.rwmutex.RLock()
	defer m.rwmutex.RUnlock()
	v, ok := m.thisMap[key]
	return v, ok
}

func (m *Map) Set(key, value interface{}) {
	m.rwmutex.Lock()
	defer m.rwmutex.Unlock()
	m.thisMap[key] = value
}

func (m *Map) Delete(key interface{}) {
	m.rwmutex.Lock()
	defer m.rwmutex.Unlock()
	if _, ok := m.thisMap[key]; ok {
		delete(m.thisMap, key)
	}

}
