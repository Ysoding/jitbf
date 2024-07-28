package jitbf

type OpKind byte

const (
	OpEOF           OpKind = 0
	OpInc           OpKind = '+'
	OpDec           OpKind = '-'
	OpLeft          OpKind = '<'
	OpRight         OpKind = '>'
	OpInput         OpKind = ','
	OpOutput        OpKind = '.'
	OpJumpIfZero    OpKind = '['
	OpJumpIfNonZero OpKind = ']'
)

type Op struct {
	Kind    OpKind
	Operand int
}

type Ops struct {
	Items []Op
}

func NewOps() *Ops {
	return &Ops{Items: make([]Op, 0)}
}

func (ops *Ops) Add(op Op) {
	ops.Items = append(ops.Items, op)
}

func (ops *Ops) Count() int {
	return len(ops.Items)
}

type AddrStack struct {
	Items []int
}

func NewAddrStack() *AddrStack {
	return &AddrStack{Items: make([]int, 0)}
}

func (s *AddrStack) Push(addr int) {
	s.Items = append(s.Items, addr)
}

func (s *AddrStack) Pop() int {
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[0 : len(s.Items)-1]
	return item
}

func (s *AddrStack) IsEmpty() bool {
	return len(s.Items) == 0
}

type Memory struct {
	Items []byte
}

func NewMemory() *Memory {
	return &Memory{Items: make([]byte, 0)}
}

func (m *Memory) Add(b byte) {
	m.Items = append(m.Items, b)
}

func (m *Memory) Count() int {
	return len(m.Items)
}
