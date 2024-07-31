package main

import (
	"fmt"
	"os"
)


func main() {
  fmt.Println("[MAIN::LOG]: -> init")
  server, err := NewServer(":9091")

  if err != nil {
    fmt.Println(err)
    os.Exit(84)
  }

  server.Start()
}
