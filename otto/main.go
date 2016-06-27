package main

import(
  "io/ioutil"

  "github.com/robertkrimen/otto"
)

func main() {
  f, err := ioutil.ReadFile( "./node/index.js" )

  if err != nil {
    panic(err)
  }

  vm := otto.New()
  vm.Run(string(f))
}
