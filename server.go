package main

import (
	"fmt"
	"net"
	"sync"
	"zixyos/goedges/utils"
)

type Server struct {
  wg sync.WaitGroup
  mutex sync.Mutex
  conn chan net.Conn
  shutdown chan struct{}
  Listener net.Listener
  client (map[string] *Client)
}

func (s *Server) Start() {
  defer s.Listener.Close()
  s.wg.Add(1)
  go s.handleConnection()

  <- s.shutdown
  close(s.conn)
  s.wg.Wait()
}

func (s *Server) handleConnection() {
  defer s.wg.Done()

  for {
    conn, err := s.Listener.Accept()
    if err != nil {
      break
    }
    s.wg.Add(1)
    go s.handleClient(conn)
  }
}

func (s *Server) addClient(con net.Conn) *Client {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  fmt.Println("Creating Client")

  clientId := utils.GenerateRandom()
  client := NewClient(clientId)

  go func() {
    client.conn <- con
  }()

  s.client[clientId] = client
  fmt.Println("Client created")

  return client
}

func (s *Server) removeClient(id string) {
  s.mutex.Lock();
  defer s.mutex.Unlock();
  delete(s.client, id);
  fmt.Println("client: ", id, "removed!")
}


func (s *Server) handleClient(conn net.Conn) {
  defer s.wg.Done()
  fmt.Println("New Connections here")
  client := s.addClient(conn)
  defer s.removeClient(client.id)
  fmt.Println("NEW CLIENT CONNECTED: ", client.id)

  conn.Close()

}

func NewServer(port string) (*Server, error) {
  l, err := net.Listen("tcp", port)
  if err != nil {
    return nil, err
  }

  return &Server{
    Listener: l,
    conn: make(chan net.Conn),
    shutdown: make(chan struct{}),
    client: make(map[string] *Client),
  }, nil
}
