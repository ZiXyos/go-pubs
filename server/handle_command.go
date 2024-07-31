package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
	"zixyos/goedges/client"
)

func (s *Server) receive_command(client *client.Client) {
  fmt.Println("ready to receive command")
  go func (conn net.Conn) {
    defer conn.Close();
    conn.Write([]byte("you can now send a command\n"));
    reader := bufio.NewReader(conn);

    for {
      conn.SetReadDeadline(time.Now().Add(5 * time.Minute));
      message, err := reader.ReadString('\n');
      if err != nil {
        if err == io.EOF {
          fmt.Println("Client %s disconnected", client.Id);
        }
        break
      }

      message = strings.TrimSpace(message);
      response := s.handle_command(client) 

      conn.SetWriteDeadline(time.Now().Add(5 * time.Minute));
      n, err := conn.Write([]byte(response + "\n"));
      if err != nil {
        fmt.Printf("Error writing to client %s: %v\n", client.Id, err)
        break
      }

      fmt.Println(n);
    }
  }(<-client.Conn)
}

func (s *Server) handle_command(client *client.Client) string {
  return "PONG to: " + client.Id
}
