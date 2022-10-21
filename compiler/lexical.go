package compiler

import (
	"fmt"
)

func Lexical(content string) {

	for i := 0; i < len(content); i++ {
		// 영문, 언어바 판별
		if (content[i] >= 'A' && content[i] <= 'Z') || content[i] == '_' || (content[i] >= 'a' && content[i] <= 'z') {
			fmt.Printf("Ident: %c", content[i])

			for true {
				i += 1
				if (content[i] >= 'A' && content[i] <= 'Z') || content[i] == '_' || (content[i] >= 'a' && content[i] <= 'z') {
					fmt.Printf("%c", content[i])
				} else {
					break
				}
			}
			fmt.Println()
		}

		// 상수 판별
		if content[i] >= '0' && content[i] <= '9' {
			fmt.Printf("Constant: %c", content[i])

			for true {
				i += 1
				if content[i] >= '0' && content[i] <= '9' {
					fmt.Printf("%c", content[i])
				} else {
					break
				}
			}
			fmt.Println()
		}

		// 덧셈 연산자 판별
		if content[i] == '+' {
			fmt.Println("Operator: +")
		}

		// 뺄셈 연산자 판별
		if content[i] == '-' {
			fmt.Println("Operator: -")
		}

		// 곱셈 연산자 판별
		if content[i] == '*' {
			fmt.Println("Operator: *")
		}

		// 나눗셈 연삼자 판별
		if content[i] == '/' {
			fmt.Println("Operator: /")
		}

		// 치환 연산자 판별
		if content[i] == '=' {
			fmt.Println("Operator: =")
		}

		// 구분자 판별
		if content[i] == ';' {
			fmt.Println("Seperator: ;")
		}

		if content[i] == '(' {
			fmt.Println("Seperator: (")
		}

		if content[i] == ')' {
			fmt.Println("Seperator: )")
		}

		// 텍스트 판별
		if content[i] == '"' {
			fmt.Printf("String: ")
			i += 1
			for content[i] != '"' {
				fmt.Printf(string(content[i]))
				i += 1
			}
			fmt.Println()
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
