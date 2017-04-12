package main

import (
  "log"
  "os"
)

func main() {
  // update this to read a config file
  if len(os.Args) != 2 {
    log.Fatal("Missing token string arugment")
    log.Fatal("Usage: go run src/main.go tokenstring")
    os.Exit(1)
  }

  var token = os.Args[1]
  getStatus(token)
}
