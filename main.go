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
    /*sockErr = syscall.SetsockoptInt(
      int(fd), 
      syscall.IPPROTO_TCP,
      syscall.SO_RCVTIMEO,
      60000,
    );*/
     // sockErr = utils.SetCustomTCPOption(int(fd), 200, "new_client")
     /*sockErr = syscall.SetsockoptString(int(fd), syscall.IPPROTO_TCP, 200, "random_id")
      if sockErr != nil {
        fmt.Errorf("Error: unable to set SocketOptionString");
      }*/
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
