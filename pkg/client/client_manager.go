package client

import (
	"net"
	"sync"
)

type Client struct {
  Id string
  Conn net.Conn
  Mut sync.Mutex
}

func NewClient(id string, conn net.Conn) *Client {
  return &Client { 
    Id: id,
    Conn: conn, 
  }
}

