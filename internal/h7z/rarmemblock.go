package h7z

type rarMemBlock struct {
	heap    *heap
	Address int
}

func newRarMemBlock(heap *heap) rarMemBlock {
	return rarMemBlock{heap: heap}
}

func (r *rarMemBlock) Stamp() int {
	return r.heap.UInt16(r.Address)
}

func (r *rarMemBlock) SetStamp(v int) {
	r.heap.PutUInt16(r.Address, v)
}

func (r *rarMemBlock) InsertAt(p rarMemBlock) {
	tmp := newRarMemBlock(r.heap)
	r.SetPrev(p.Address)
	tmp.Address = r.Prev()
	r.SetNext(tmp.Next())
	tmp.SetNext(r.Address)
	tmp.Address = r.Next()
	tmp.SetPrev(r.Address)
}

func (r *rarMemBlock) Remove() {
	tmp := newRarMemBlock(r.heap)
	tmp.Address = r.Prev()
	tmp.SetNext(r.Next())
	tmp.Address = r.Next()
	tmp.SetPrev(r.Prev())
}

func (r *rarMemBlock) Next() int {
	return r.heap.UInt32(r.Address + 4)
}

func (r *rarMemBlock) SetNext(next int) {
	r.heap.PutUInt32(r.Address+4, next)
}

func (r *rarMemBlock) Prev() int {
	return r.heap.UInt32(r.Address + 8)
}

func (r *rarMemBlock) SetPrev(prev int) {
	r.heap.PutUInt32(r.Address+8, prev)
}

func (r *rarMemBlock) Nu() int {
	return r.heap.UInt16(r.Address + 2)
}

func (r *rarMemBlock) SetNu(prev int) {
	r.heap.PutUInt16(r.Address+2, prev)
}
