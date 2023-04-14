package main

import (
  "flag"
  "fmt"
  "os"
  "bufio"

  "github.com/karchx/newape/lexer"
  "github.com/karchx/newape/tokens"
)

func main() {
  f := flag.String("f", "", "File for the interpreter")
  flag.Parse()
 
  file, err := os.Open(*f)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  reader := bufio.NewReader(file)
  lex := lexer.New(reader, *f)

  var toks []tokens.Token

  for {
    tok, err := lex.NextToken()

    toks = append(toks, tok)
    if err != nil {
      panic(err)
    }
    if tok.Type == tokens.EOF {
      break
    }
  }

  fmt.Println(toks)
}
