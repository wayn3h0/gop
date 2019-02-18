package lfu

import (
	goheap "container/heap"
)

type heap struct {
	list  entries
	table map[string]*entry
}

func (h *heap) Initialize() *heap {
	h.list = make(entries, 0)
	goheap.Init(&h.list)
	h.table = make(map[string]*entry)

	return h
}

func (h *heap) Count() int {
	return h.list.Len()
}

func (h *heap) Contains(key string) bool {
	_, ok := h.table[key]
	return ok
}

func (h *heap) Get(key string) interface{} {
	if e, ok := h.table[key]; ok {
		e.Count++
		goheap.Fix(&h.list, e.Index)
		return e.Value
	}

	return nil
}

func (h *heap) Save(key string, value interface{}) {
	if element, ok := h.table[key]; ok {
		element.Value = value
	} else {
		e := &entry{
			Key:   key,
			Value: value,
		}
		goheap.Push(&h.list, e)
		h.table[key] = e
	}
}

func (h *heap) Discard() {
	if len(h.list) == 0 {
		return
	}

	entry := goheap.Pop(&h.list).(*entry)
	delete(h.table, entry.Key)
}

func (h *heap) Remove(key string) {
	if element, ok := h.table[key]; ok {
		goheap.Remove(&h.list, element.Index)
		delete(h.table, key)
	}
}
