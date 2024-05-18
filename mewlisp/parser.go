package mewlisp

import (
	"fmt"
	"strings"
)

const MIN_TOK_CAP = 32

type TokenPos struct {
	line uint
	startpos uint
	endpos uint
}

type MewlToken struct {
	lexeme string
	position TokenPos
}

func (t * MewlToken) String() string {
	return fmt.Sprintf("[%s][%d,[%d,%d]]", t.lexeme, t.position.line, t.position.startpos, t.position.endpos)
}

type MewlParser struct {
	source string
	tokens []MewlToken
}

func NewParser(src string) *MewlParser {
	a := strings.ReplaceAll(src, "(", " ( ")
	b := strings.ReplaceAll(a, ")", " ) ")
	return &MewlParser{
		source: b,
		tokens: make([]MewlToken, 0),
	}
}

func (p * MewlParser) tokenize() []MewlToken{
	chs := []rune(p.source)
	
	var output []MewlToken
	var curtok []rune
	pos := 0
	lineno := 0

	for pos < len(chs) {
		for chs[pos] != ' ' && chs[pos] != '\t'{
			if chs[pos] == '\n'{
				lineno += 1
				pos += 1
				continue
			}

			if chs[pos] == '/' && chs[pos + 1] == '/' {
				lineno += 1
				pos += 2

				for chs[pos] != '\n' {
					pos += 1
				}

				pos += 1
				continue
			}

			curtok = append(curtok, chs[pos])
			pos += 1

		}

		if len(curtok) > 0 {
			output = append(output, MewlToken{
				lexeme: string(curtok),
				position: TokenPos{
					line: uint(lineno),
					startpos: uint(pos - len(curtok)),
					endpos: uint(pos),

				},
			})


			curtok = []rune{}
			
		}

		pos += 1


	}

	return output

	
}


func (p * MewlParser) Debug() {
	toks := p.tokenize()


	for _, t := range toks {
		fmt.Println(t.String())
	}
}

