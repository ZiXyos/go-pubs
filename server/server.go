package server

import (
	"fmt"
	"net"
	"sync"
	"zixyos/goedges/utils"
  . "zixyos/goedges/client"
)

type Server struct {
  wg sync.WaitGroup
  mutex sync.Mutex
  conn chan net.Conn
  shutdown chan struct{}
  Listener net.Listener
  client (map[string] *Client)
  topic (map[string] *Topic)
  commandList []string
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
    client.Conn <- con
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
  fmt.Println("NEW CLIENT CONNECTED: ", client.Id)
  s.receive_command(client)
}

func NewServer(port string) (*Server, error) {
  l, err := net.Listen("tcp", port)
  if err != nil {
    return nil, err
  }

  commandList := make([]string, 0, 10)
  utils.GenerateCommand("PUB", &commandList);
  utils.GenerateCommand("SUB", &commandList);

  return &Server{
    Listener: l,
    conn: make(chan net.Conn),
    shutdown: make(chan struct{}),
    client: make(map[string] *Client),
    topic: make(map[string] *Topic),
    commandList: commandList,
  }, nil
}
