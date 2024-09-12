package auth

type Credential struct {
  Username string;
  Password string;
}

type Auth interface {
  Authentificate(cred Credential) error;
  Register(cred Credential) error;
}

type AuthConfig struct {
  
}
type BasicAuth struct {
  config struct{}
}

func (auth *BasicAuth) Authentificate(cred Credential) error {
  return nil
} 

func (auth *BasicAuth) Register(cred Credential) error {
  return nil
}
