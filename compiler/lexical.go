package compiler

import (
	"fmt"
	"strconv"

	"github.com/lukekeum/golanc/token"
)

func Lexical(content string, store []token.Token) {

	for i := 0; i < len(content); i++ {
		// 영문, 언어바 판별

		if (content[i] >= 'A' && content[i] <= 'Z') || content[i] == '_' || (content[i] >= 'a' && content[i] <= 'z') {
			value := ""
			value += fmt.Sprintf("%c", content[i])

			for true {
				i += 1
				if (content[i] >= 'A' && content[i] <= 'Z') || content[i] == '_' || (content[i] >= 'a' && content[i] <= 'z') {
					value += fmt.Sprintf("%c", content[i])
				} else {
					break
				}
			}
			store = append(store, *token.New(token.IDENT, value))
		}

		// 상수 판별
		if content[i] >= '0' && content[i] <= '9' {
			value := fmt.Sprintf("%c", content[i])

			for true {
				i += 1
				if content[i] >= '0' && content[i] <= '9' {
					value += fmt.Sprintf("%c", content[i])
				} else {
					break
				}
			}
			value_int, _ := strconv.Atoi(value)

			store = append(store, *token.New(token.CONST, value_int))
		}

		// 덧셈 연산자 판별
		if content[i] == '+' {
			store = append(store, *token.New(token.OPER, '+'))
		}

		// 뺄셈 연산자 판별
		if content[i] == '-' {
			store = append(store, *token.New(token.OPER, '-'))
		}

		// 곱셈 연산자 판별
		if content[i] == '*' {
			store = append(store, *token.New(token.OPER, '*'))
		}

		// 나눗셈 연삼자 판별
		if content[i] == '/' {
			store = append(store, *token.New(token.OPER, '/'))
		}

		// 치환 연산자 판별
		if content[i] == '=' {
			store = append(store, *token.New(token.OPER, '='))
		}

		// 구분자 판별
		if content[i] == ';' {
			store = append(store, *token.New(token.SEPER, ';'))
		}

		if content[i] == '(' {
			store = append(store, *token.New(token.SEPER, '('))
		}

		if content[i] == ')' {
			store = append(store, *token.New(token.SEPER, ')'))
		}

		if content[i] == '"' {
			store = append(store, *token.New(token.SEPER, '"'))
		}

		if content[i] == '\'' {
			store = append(store, *token.New(token.SEPER, '\''))
		}

		// 텍스트 판별
		if content[i] == '"' {
			value := ""

			i += 1
			for content[i] != '"' {
				value += fmt.Sprintf("%c", content[i])
				i += 1
			}
			store = append(store, *token.New(token.STRING, value))
		}

		// /* */ 주석 판별
		if content[i] == '/' {
			if content[i+1] == '*' {
				i += 1
				for content[i] != '/' {
					for content[i] != '*' {
						i += 1
					}
				}
			}

			if content[i+1] == '/' {
				i += 1
				for content[i] != '\n' {
					i += 1
				}
			}
		}
	}

}
