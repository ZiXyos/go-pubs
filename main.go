package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"zixyos/goedges/pkg/server"
)


func initListener(network, address string, conn syscall.RawConn) error {
  var sockErr error;
  if err := conn.Control(func(fd uintptr) {
  }); err != nil {
    return err
  }
  return sockErr
}

func main() {
  fmt.Println("[MAIN::LOG]: -> init")
  listenerConfig := net.ListenConfig{
    Control: initListener,
  }
  server, err := server.NewServer(":9091", listenerConfig, nil);

  if err != nil {
    fmt.Println(err)
    os.Exit(84)
  }

  server.Start()
}
