package server

import "fmt"

func (s *Server) AuthenticateCommand(input []string) error {
	fmt.Println(input);
	return nil
}
