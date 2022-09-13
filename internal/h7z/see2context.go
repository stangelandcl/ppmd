package h7z

type see2context struct {
	summ, shift, count int
}

func newSee2Context(val int) see2context {
	//fmt.Printf("newsee2: %v\n", val)
	s := see2context{}
	s.shift = (periodBits - 4) & 0xff
	s.summ = (val << s.shift) & 0xffff
	s.count = 4
	return s
}

func (s *see2context) Mean() int {
	r := urshift(s.summ, s.shift)
	//fmt.Printf("mean: %v\t%v\t%v\n", r, s.summ, s.shift)
	s.summ -= r
	if r == 0 {
		r++
	}
	return r
}

func (s *see2context) Count() int {
	return s.count
}

func (s *see2context) SetCount(c int) {
	s.count = c & 0xff
}

func (s *see2context) Shift() int {
	return s.shift
}

func (s *see2context) SetShift(shift int) {
	//fmt.Printf("shift: %v\n", shift)
	s.shift = shift & 0xff
}

func (s *see2context) Summ() int {
	return s.summ
}

func (s *see2context) SetSumm(summ int) {
	s.summ = summ & 0xffff
}

func (s *see2context) IncSumm(dsum int) {
	s.summ += dsum
}

func (s *see2context) Update() {
	//fmt.Printf("update: %v\n", s.shift)
	if s.shift < periodBits {
		s.count--
		if s.count == 0 {
			s.summ += s.summ
			s.count = 3 << s.shift
			s.shift++
		}
	}
	s.summ &= 0xffff
	s.count &= 0xff
	s.shift &= 0xff
}
