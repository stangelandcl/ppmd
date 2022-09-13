package h7z

type stateRef struct {
	Successor, symbol, freq int
}

func newStateRef(s *state) stateRef {
	r := stateRef{}
	r.SetFreq(s.Freq())
	r.SetSymbol(s.Symbol())
	r.Successor = s.Successor()
	return r
}

func (s *stateRef) IncrementFreq(dfreq int) {
	s.freq = (s.freq + dfreq) & 0xff
}

func (s *stateRef) DecrementFreq(dfreq int) {
	s.freq = (s.freq - dfreq) & 0xff
}

func (s *stateRef) Freq() int {
	return s.freq
}

func (s *stateRef) SetFreq(f int) {
	s.freq = f & 0xff
}

func (s *stateRef) Symbol() int {
	return s.symbol
}

func (s *stateRef) SetSymbol(sym int) {
	s.symbol = sym & 0xff
}
