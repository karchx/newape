package lexer

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/karchx/newape/tokens"
)

type Lexer struct {
	input        *bufio.Reader
	position     int
	readPosition int
	ch           byte
	fileName     string
	line         int
	lines        []string
	lineHadNonWS bool
	col          int
}

func New(is *bufio.Reader, filename string) *Lexer {
	l := &Lexer{input: is, fileName: filename}

	for {
		line, err := l.input.ReadString('\n')
		if errors.Is(err, io.EOF) {
			l.lines = append(l.lines, line)
			break
		}
		if err != nil {
			break
		}
		l.lines = append(l.lines, line)
	}

	l.input.Reset(strings.NewReader(strings.Join(l.lines, "")))

	l.readChar()

	l.line = 1
	l.col = 1

	return l
}

func newToken(tokenType tokens.TokenType, ch byte) tokens.Token {
	return tokens.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() (tokens.Token, error) {
	var tok tokens.Token

	switch l.ch {
  case 0:
    tok = newToken(tokens.EOF, '\x00')
	case '+':
		if !l.lineHadNonWS {
			tok = l.handleComment()
		} else {
			l.lineHadNonWS = true
			tok = newToken(tokens.PLUS, l.ch)
		}
	}

  l.readChar()
	return tok, nil
}

// readChar reads the next character in the input
func (l *Lexer) readChar() {
	l.ch, _ = l.input.ReadByte()
	l.position = l.readPosition
	l.readPosition += 1
	l.col++
}

// isDigit returns true if the given byte is a digit
func (l *Lexer) isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func (l *Lexer) handleComment() tokens.Token {
	var tok tokens.Token
	currLine := l.lines[l.line-1]
	if currLine[len(currLine)-1] == '\r' || currLine[len(currLine)-1] == '\n' {
		currLine = currLine[:len(currLine)-1]
	}
	tok.Literal = strings.TrimSpace(currLine)
	l.consumeLine()
	l.line++
	l.col = 0
	return tok
}

func (l *Lexer) consumeLine() {
	for l.ch != '\n' && l.ch != '\r' {
		l.readChar()
		if l.peekChar() == 0 {
			break
		}
	}
}

func (l *Lexer) peekChar() byte {
	ch, err := l.input.Peek(1)
	if err != nil {
		return 0
	}

	return ch[0]
}
