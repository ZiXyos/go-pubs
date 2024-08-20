package server

import (
	"fmt"
	"net"
	"sync"
	"time"
	. "zixyos/goedges/client"
	"zixyos/goedges/utils"
)

type Server struct {
  wg sync.WaitGroup
  mutex sync.RWMutex
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
    select {
    case<-s.shutdown:
      return
    default:
      conn, err := s.Listener.Accept()
      if err != nil {
        break
      }
      s.wg.Add(1)
      go s.handleClient(conn)
    }
  }
}

func (s *Server) addClient(con net.Conn) *Client {
  s.mutex.Lock();
  defer s.mutex.Unlock();

  clientId := utils.GenerateRandom();
  client := NewClient(clientId, con);

  s.client[clientId] = client;
  return client
}

func (s *Server) removeClient(id string) {
  s.mutex.Lock();
  defer s.mutex.Unlock();
  delete(s.client, id);
  fmt.Println("client: ", id, "removed!")
}

func (s *Server) sendMessage(clientId string, message string) {
  s.mutex.RLock();
  client, ok := s.client[clientId];
  s.mutex.RUnlock();

  if !ok {
    fmt.Println("Client not found");
    return
  }

  client.Mut.Lock();
  defer client.Mut.Unlock();

    client.Conn.SetWriteDeadline(time.Now().Add(5 * time.Second));
  _, err := client.Conn.Write([]byte(message+"\n"));
  if err != nil {
    fmt.Println("Error sending message to client: ", clientId, err);
    return 
  }

}

func (s *Server) handleClient(conn net.Conn) {
  defer s.wg.Done();
  defer conn.Close();

  client := s.addClient(conn);
  fmt.Println("NEW CLIENT CONNECTED: ", client.Id);
  s.receive_command(client);
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
