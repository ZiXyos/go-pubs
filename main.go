package main

import (
	"fmt"
	"os"
  "zixyos/goedges/server"
)

func main() {
  fmt.Println("[MAIN::LOG]: -> init")
  server, err := server.NewServer(":9091")

  if err != nil {
    fmt.Println(err)
    os.Exit(84)
  }

  server.Start()
}
