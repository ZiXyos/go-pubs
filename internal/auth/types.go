package auth

import (
	"net"
	"zixyos/goedges/pkg/client"
)

type Credentials struct {
  Username string;
  Password string;
}

type Auth interface {
  Authentificate(cred Credentials) (*client.Client, error);
  Register(cred Credentials, conn net.Conn) (*client.Client, error);
}

type AuthConfig struct {
  
}
