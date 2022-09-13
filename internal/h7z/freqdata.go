package h7z

type freqData struct {
	heap    *heap
	Address int
}

func newFreqData(buf *heap) freqData {
	return freqData{heap: buf}
}

func (f *freqData) Stats() int {
	return f.heap.Int32(f.Address + 2)
}

func (f *freqData) SetStats(state int) {
	f.heap.PutInt32(f.Address+2, state)
}

func (f *freqData) SummFreq() int {
	return f.heap.UInt16(f.Address)
}

func (f *freqData) SetSummFreq(freq int) {
	f.heap.PutUInt16(f.Address, freq)
}

func (f *freqData) IncrementSummFreq(fr int) {
	f.SetSummFreq(f.SummFreq() + fr)
}
