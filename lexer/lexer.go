package lexer

import (
	"strings"

	"github.com/Ysoding/jitbf/jitbf"
)

type Lexer struct {
	input    string
	position int
}

func New(input string) *Lexer {
	return &Lexer{input: input, position: 0}
}

func (l *Lexer) Next() jitbf.OpKind {
	for l.position < len(l.input) && !l.isBFCmd(l.input[l.position]) {
		l.position += 1
	}

	if l.position >= len(l.input) {
		return jitbf.OpEOF
	}

	res := jitbf.OpKind(l.input[l.position])
	l.position += 1
	return res
}

func (l *Lexer) isBFCmd(ch byte) bool {
	return strings.Contains("+-<>,.[]", string(ch))
}

func (l *Lexer) Pos() int {
	return l.position
}
