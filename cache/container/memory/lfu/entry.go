package lfu

type entry struct {
	Key   string
	Value interface{}
	Index int // index of Item in the heap
	Count int // accessed count
}

type entries []*entry

func (es entries) Len() int {
	return len(es)
}

func (es entries) Less(x, y int) bool {
	return es[x].Count < es[y].Count
}

func (es entries) Swap(x, y int) {
	es[x], es[y] = es[y], es[x]
	es[x].Index, es[y].Index = x, y
}

func (es *entries) Push(x interface{}) {
	index := len(*es)
	entry := x.(*entry)
	entry.Index = index
	*es = append(*es, entry)
}

func (es *entries) Pop() interface{} {
	old := *es
	index := len(old)
	entry := old[index-1]
	entry.Index = -1 // for safety
	*es = old[0 : index-1]

	return entry
}
