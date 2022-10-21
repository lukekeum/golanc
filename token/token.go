package token

type Token struct {
	TokenType TokenType
	Value     interface{}
}

type TokenType int

const (
	IDENT TokenType = iota
	CONST
	OPER
	SEPER
	STRING
)

func New(tokenType TokenType, value interface{}) *Token {
	return &Token{TokenType: tokenType, Value: value}
}
