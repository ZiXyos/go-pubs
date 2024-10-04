package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"time"
	"zixyos/goedges/pkg/server"

	"github.com/charmbracelet/log"
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
  loggerOption := log.Options {
    ReportCaller: true,
    ReportTimestamp: true,
    Prefix: "Goedges 🛫",
    TimeFormat: time.Kitchen,
  }
  server, err := server.NewServer(":9091", listenerConfig, nil, loggerOption);

  if err != nil {
    fmt.Println(err)
    os.Exit(84)
  }

  server.Start()
}
