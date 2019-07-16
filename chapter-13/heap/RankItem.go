package heap

import "container/heap"

// A min-heap for items with their ranks.
type RankItem struct {
	Value interface{}
	Rank int
}

type RankItemHeap []*RankItem

func (this RankItemHeap) Len() int {
	return len(this)
}

func (this RankItemHeap) Less(i, j int) bool {
	return this[i].Rank < this[j].Rank
}

func (this RankItemHeap) Swap(i, j int) {
	this[i].Rank, this[j].Rank = this[j].Rank, this[i].Rank
	this[i].Value, this[j].Value = this[j].Value, this[i].Value
}

func (this *RankItemHeap) Push(item interface{}) {
	*this = append(*this, item.(*RankItem))
}

func (this *RankItemHeap) Pop() (item interface{}) {
	length := len(*this)
	item = (*this)[length-1]
	*this = (*this)[:length-1]

	return
}

func NewRankItemHeap() *RankItemHeap {
	var this RankItemHeap
	this = make([]*RankItem, 0)

	heap.Init(&this)

	return &this
}
