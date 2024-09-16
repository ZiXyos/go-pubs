package server

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
	. "zixyos/goedges/internal/auth"
	"zixyos/goedges/pkg/client"
	. "zixyos/goedges/pkg/client"
	"zixyos/goedges/pkg/types"
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
  commandsList (map[string] *types.CommandFunc)
  internalCommandsList (map[string] *types.InternalCommandFunc)
  authentificator Auth
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
      client, err := s.authenticateConn(conn);
      s.wg.Add(1)
      go s.handleClient(client)
    }
  }
}

func (s *Server) authenticateConn(conn net.Conn) (*Client, error) {
  reader := bufio.NewReader(conn);
  for {

    line, err := reader.ReadString('\n');
    if err != nil {
      return nil, err
    }

    command := strings.Split(strings.TrimSpace(line), " ");
    fun, ok := s.internalCommandsList[command[0]];
    if ok {
      deffered := *fun;
      err = deffered(command, conn)
      if err != nil {
        fmt.Println(err);
        fmt.Fprintf(conn, "%s\n", err);
        continue
      }

      fmt.Println(err);
      return s.client[command[1]], nil
    }
  }
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

func (s *Server) handleClient(client *client.Client) {
  defer s.wg.Done();
  defer client.Conn.Close();

  s.receive_command(client);
}

func NewServer(port string, config net.ListenConfig, auth Auth) (*Server, error) {
  if auth == nil {
    auth = &BasicAuth{}
  }
  l, err := config.Listen(context.Background(), "tcp", ":9091")
  if err != nil {
    return nil, err
  }

  commandList := make([]string, 0, 10)
  utils.GenerateCommand("PUB", &commandList);
  utils.GenerateCommand("SUB", &commandList);
  utils.GenerateCommand("AUTH", &commandList);

  serv := &Server{
    Listener: l,
    conn: make(chan net.Conn),
    shutdown: make(chan struct{}),
    client: make(map[string] *Client),
    topic: make(map[string] *Topic),
    commandList: commandList,
    commandsList: make(map[string]*types.CommandFunc),
    internalCommandsList: make(map[string] *types.InternalCommandFunc), 
    authentificator: auth,
  }
  utils.GenerateInternalCommandMap("AUTH", serv.AuthenticateCommandWrapper, &serv.internalCommandsList);

  utils.GenerateCommandMap("CREATE", serv.handle_create, &serv.commandsList);
  utils.GenerateCommandMap("PUB", serv.handle_publish, &serv.commandsList);
  utils.GenerateCommandMap("SUB", serv.handle_subscribe, &serv.commandsList);

  return serv, nil
}
