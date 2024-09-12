package types

import (
	"net"
	"zixyos/goedges/pkg/client"
)


type InternalCommandFunc func([]string, net.Conn) error;
type CommandFunc func(*client.Client, []string) string;
