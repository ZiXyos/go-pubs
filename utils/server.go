package utils

import (
	"errors"
	"fmt"
	"net"
	"syscall"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func ExtractClientId(conn net.Conn) (string, error) {
  buff := make([]byte, 4096)
  n, err := conn.Read(buff);
  if err != nil {
    return "", err
  }

  packet := gopacket.NewPacket(buff[:n], layers.LayerTypeTCP, gopacket.Default);
  if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
    tcp, ok := tcpLayer.(*layers.TCP);
    if !ok {
      return "", errors.New("Error: failed to convert TCP layer");
    }

    for _, opt := range tcp.Options {
      if opt.OptionType == layers.TCPOptionKindMSS {
        return string(opt.OptionData), nil
      }
      if opt.OptionType == 200 {
        return string(opt.OptionData), nil
      }
    }
  }
  for _, layer := range packet.Layers() {
    fmt.Println("PACKET LAYER:", layer.LayerType());
  }
  return "", errors.New("client_id not found in TCP options")
}

func SetCustomTCPOption(fd, optType int, optData string) error {
  optBytes := []byte(optData);

  err := syscall.SetsockoptString(
    fd, syscall.IPPROTO_TCP,
    optType,
    string(optBytes),
  );
   if err != nil {
        return err
    }
    return nil
}
