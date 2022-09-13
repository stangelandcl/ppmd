package h7z

type rarNode struct {
	heap    *heap
	Address int
}

func newRarNode(buf *heap) rarNode {
	return rarNode{heap: buf}
}

func (r *rarNode) Next() int {
	return r.heap.Int32(r.Address)
}

func (r *rarNode) SetNext(next int) {
	r.heap.PutInt32(r.Address, next)
}
