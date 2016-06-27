package main

import "fmt"

func main() {
  values := make([]string, 0)
  values = []string{ "hola", "mundo" }

  for _, v := range values {
    fmt.Println( v )
  }
}
