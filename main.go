package main

import (
  "flag"
  "fmt"
  "os"
  "bufio"

  "github.com/karchx/newape/lexer"
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

  fmt.Println(lex)
}
