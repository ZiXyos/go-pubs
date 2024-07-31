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

/*func (s *Server) ClientUpscale() error {
  if len(s.client) > (cap(s.client) / 2 ) {
    newSlice := make([]Client, len(s.client), (cap(s.client) * 2 + len(s.client)));
    copy(newSlice, s.client);
    s.client = newSlice;
    return nil
  }
  return fmt.Errorf("cannot upscale content, no enough active client")
}

func (s *Server) ClientDownScale() error {
  if len(s.client) < ((cap(s.client) + 10) / 2) {
    newSlice := make([]Client, len(s.client), cap(s.client) / 2 + len(s.client));
    copy(newSlice, s.client);
    s.client = newSlice;
    return nil
  } 
  return fmt.Errorf("cannot downscale client list, too many active client")
}*/

