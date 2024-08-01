package client

import (
	"net"
)

type Client struct {
  Id string
  Conn chan net.Conn
}

func NewClient(id string) *Client {
  return &Client { 
    Id: id,
    Conn: make(chan net.Conn), 
  }
}

