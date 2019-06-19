package heap

import "container/heap"

// ArrayItem tracks three items. The values of the array entries, the respective
// array-id and the respective item-id in the array.
type ArrayItem struct {
	Value, ArrayId, ItemId int
}

type ArrayItemHeap []ArrayItem

func (this ArrayItemHeap) Len() int {
	return len(this)
}

func (this ArrayItemHeap) Less(i, j int) bool {
	return this[i].Value < this[j].Value
}

func (this ArrayItemHeap) Swap(i, j int) {
	this[i].Value, this[j].Value = this[j].Value, this[i].Value
	this[i].ItemId, this[j].ItemId = this[j].ItemId, this[i].ItemId
	this[i].ArrayId, this[j].ArrayId = this[j].ArrayId, this[i].ArrayId
}

func (this *ArrayItemHeap) Push(item interface{}) {
	*this = append(*this, item.(ArrayItem))
}

func (this *ArrayItemHeap) Pop() (item interface{}) {
	length := len(*this)
	item = (*this)[length-1]

	slice := (*this)[:length-1]
	*this = slice

	return
}

func NewArrayItemHeap() *ArrayItemHeap {
	var this ArrayItemHeap
	this = make([]ArrayItem, 0)
	return &this
}

func MergeSortedArrays(arrays [][]int) []int {
	h := NewArrayItemHeap()
	heap.Init(h)


	// Track total entries for the final sorted list.
	totalEntries := 0

	// Push first element of each array.
	for i := 0; i < len(arrays); i++ {
		heap.Push(h, ArrayItem{arrays[i][0], i, 0})
		totalEntries += len(arrays[i])
	}

	// Initialize the final sorted list.
	sorted := make([]int, 0, totalEntries)

	// Pop the smallest element from the heap. Push the next item
	// of the same array, if available, into the heap.
	for h.Len() > 0 {
		item := heap.Pop(h).(ArrayItem)
		sorted = append(sorted, item.Value)

		if item.ItemId >= len(arrays[item.ArrayId])-1 {
			continue
		}

		newEntry := ArrayItem{arrays[item.ArrayId][item.ItemId+1],
				item.ArrayId, item.ItemId+1}
		heap.Push(h, newEntry)
	}

	return sorted
}
