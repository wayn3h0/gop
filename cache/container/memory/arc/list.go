package arc

import (
	golist "container/list"
)

type entry struct {
	Key   string
	Value interface{}
}

type list struct {
	*golist.List
	Table map[string]*golist.Element
}

func (l *list) Initialize() *list {
	l.List = golist.New()
	l.Table = make(map[string]*golist.Element)

	return l
}

func (l *list) Count() int {
	return l.List.Len()
}

func (l *list) Contains(key string) bool {
	_, ok := l.Table[key]
	return ok
}

func (l *list) Get(key string) interface{} {
	if element, ok := l.Table[key]; ok {
		l.List.MoveToFront(element)
		return element.Value.(*entry).Value
	}

	return nil
}

func (l *list) Save(key string, value interface{}) {
	e := &entry{
		Key:   key,
		Value: value,
	}
	if element, ok := l.Table[key]; ok {
		element.Value = e
		l.List.MoveToFront(element)
	} else {
		l.Table[key] = l.List.PushFront(e)
	}
}

func (l *list) Discard() (string, interface{}) {
	element := l.List.Back()
	if element == nil {
		return "", nil
	}
	e := element.Value.(*entry)
	l.List.Remove(element)
	delete(l.Table, e.Key)

	return e.Key, e.Value
}

func (l *list) Remove(key string) interface{} {
	if element, ok := l.Table[key]; ok {
		l.List.Remove(element)
		delete(l.Table, key)
		e := element.Value.(*entry)
		return e.Value
	}

	return nil
}
