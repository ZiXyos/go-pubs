package auth

import (
	"net"
	"zixyos/goedges/pkg/client"
)

type BasicAuth struct {
  config struct{}
}

func (auth *BasicAuth) Authentificate(cred Credentials) (*client.Client, error) {
  return nil, nil
} 

func (auth *BasicAuth) Register(cred Credentials, conn net.Conn) (*client.Client, error) {
  return client.NewClient(cred.Username, conn), nil;
}
