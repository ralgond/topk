package topk

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	item_array []*Item
	n int
	capacity int
}

func NewPriorityQueueForTopK(k int) *PriorityQueue {
	return &PriorityQueue {
		item_array: make([]*Item, k),
		n: 0,
		capacity: k,
	}
}

func (pq PriorityQueue) Len() int { return pq.n }

func (pq PriorityQueue) Less(i, j int) bool {
	// fmt.Printf("Less, i:%d, j:%d, item_array.len:%d, pi:%p, pj:%p", i, j, pq.n, pq.item_array[i], pq.item_array[j])
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq.item_array[i].priority > pq.item_array[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.item_array[i], pq.item_array[j] = pq.item_array[j], pq.item_array[i]
	pq.item_array[i].index = i
	pq.item_array[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	if pq.n == pq.capacity {
		return
	}
	item := x.(*Item)
	item.index = pq.n
	pq.item_array[pq.n] = item
	pq.n += 1
}

func (pq *PriorityQueue) Pop() any {
	n := pq.n
	if n == 0 {
		return nil
	}
	item := pq.item_array[n-1]
	pq.item_array[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.n -= 1
	return item
}

func (pq *PriorityQueue) Top() *Item {
	n := pq.n
	item := pq.item_array[n-1]
	return item
}

func (pq *PriorityQueue) TryPush(x any) {
	item := x.(*Item)
	if pq.n == pq.capacity {
		top := pq.Top()
		if item.priority > top.priority {
			pq.Pop()
			pq.Push(item)
			heap.Fix(pq, pq.n-1)
		}
	} else {
		pq.Push(item)
		if pq.n > 1 {
			heap.Fix(pq, pq.n-1)
		}
	}
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type TOPK struct {
	pq *PriorityQueue
}

func NewTOPK(k int) *TOPK {
	return &TOPK {
		pq : NewPriorityQueueForTopK(k),
	}
}

func (topk *TOPK) Init() {
	heap.Init(topk.pq)
}

func (topk *TOPK) Add(item *Item) {
	topk.pq.TryPush(item)
}
