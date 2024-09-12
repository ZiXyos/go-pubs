package server

import (
	"errors"
	"fmt"
	"net"
	"zixyos/goedges/internal/auth"
	"zixyos/goedges/pkg/client"
)

func (s *Server) AuthenticateCommand(input []string, conn net.Conn) (*client.Client, error) {
  // AUTH <client-id> <server-password>
  if len(input[1:]) < 2 {
    return nil, errors.New("Error trying to authenticate not so much arg")
  }

  fmt.Println(input[1])
  if client, ok := s.client[input[1]]; ok {
    c, err := s.authentificator.Authentificate(
      auth.Credentials{Username: client.Id, Password: input[2]},
    );
    if err != nil {
      return nil, err
    }
    fmt.Println(c);
    return c, nil;
  } else {
    // will return a new client
    c, err := s.authentificator.Register(
      auth.Credentials{Username: input[1], Password: input[2]},
      conn,
    )

    if err != nil {
      return nil, err; 
    }

    return c, nil
  }
}

func (serv *Server) AuthenticateCommandWrapper(input []string, conn net.Conn) error {
    client, err := serv.AuthenticateCommand(input, conn)
    if err != nil {
        return err
    }
    serv.mutex.Lock()
    serv.client[client.Id] = client
    serv.mutex.Unlock()
    return nil
}
