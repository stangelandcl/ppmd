package h7z

type state struct {
	heap    *heap
	Address int
}

func newState(buf *heap) state {
	return state{heap: buf}
}

func (s *state) Symbol() int {
	return s.heap.Byte(s.Address)
}

func (s *state) SetSymbol(v int) {
	s.heap.PutByte(s.Address, byte(v))
}

func (s *state) Freq() int {
	return s.heap.Byte(s.Address + 1)
}

func (s *state) SetFreq(freq int) {
	s.heap.PutByte(s.Address+1, byte(freq))
}

func (s *state) IncrementFreq(dfreq int) {
	s.heap.PutByte(s.Address+1, byte(s.heap.Byte(s.Address+1)+dfreq))
	//s.SetFreq(s.Freq() + dfreq)
}

func (s *state) Successor() int {
	return s.heap.UInt32(s.Address + 2)
}

func (s *state) SetSuccessor(successor int) {
	s.heap.PutUInt32(s.Address+2, successor)
}

func (s *state) SetRef(r stateRef) {
	s.SetSymbol(r.symbol)
	s.SetFreq(r.Freq())
	s.SetSuccessor(r.Successor)
}

func (s *state) SetValues(ptr state) {
	s.heap.Copy(s.Address, ptr.Address, stateSize)
}

func (s *state) DecrementAddress() {
	s.Address -= stateSize
}

func (s *state) IncrementAddress() {
	s.Address += stateSize
}

func ppmdSwap(p1 *state, p2 *state) {
	for i := 0; i < stateSize; i++ {
		x := p1.heap.Byte(p1.Address + i)
		p1.heap.PutByte(p1.Address+i, byte(p2.heap.Byte(p2.Address+i)))
		p2.heap.PutByte(p2.Address+i, byte(x))
	}
}
