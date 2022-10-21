package parser

import (
	"github.com/lukekeum/golanc/token"
)

type LParser struct {
	Store []token.Token
	begin int
}

func New(store []token.Token) *LParser {
	return &LParser{Store: store, begin: 0}
}

func (p *LParser) Init() {} // TODO: Implement this
