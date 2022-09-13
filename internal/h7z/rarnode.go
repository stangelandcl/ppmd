package h7z

type rarNode struct {
	heap    *heap
	Address int
}

func newRarNode(buf *heap) rarNode {
	return rarNode{heap: buf}
}

func (r *rarNode) Next() int {
	return r.heap.UInt32(r.Address)
}

func (r *rarNode) SetNext(next int) {
	r.heap.PutUInt32(r.Address, next)
}
