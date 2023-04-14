package lexer

import (
  "bufio"
  "errors"
  "io"
)

type Lexer struct {
  input        *bufio.Reader
  position     int
  readPosition int
  fileName     string
  line         int
  lines        []string
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

  return l
}

// isDigit returns true if the given byte is a digit
func (l *Lexer) isDigit(ch byte) bool {
  if ch >= '0' && ch <= '9' {
    return true
  }
  return false
}
