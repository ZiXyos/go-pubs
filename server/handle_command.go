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
          fmt.Printf("Client %s disconnected \n", client.Id);
        }
        break
      }

      message = strings.TrimSpace(message);
      fmt.Println("[LOG::RECEIVED::MESSAGE]", message)
      response := s.handle_command(client, message) 
      conn.SetWriteDeadline(time.Now().Add(5 * time.Minute));
      _, err = conn.Write([]byte(response + "\n"));
      if err != nil {
        fmt.Printf("Error writing to client %s: %v\n", client.Id, err)
        break
      }
    }
  }(<-client.Conn)
}

func (s *Server) handle_command(client *client.Client, entry string) string {
  command := strings.Split(entry, " ");
  if command[0] == "SUB" {
    topic, err := s.FindTopic(command[1])
    if err != nil {
      return err.Error()
    }
    return "SUB TO TOPIC: " + topic.TopicId
  } else if command[0] == "CREATE" {
    topic, err := s.FindTopic(command[1])
    if err == nil && topic != nil {
      return "Topic " + topic.TopicId +" already exist"
    }
    s.createTopic(client.Id, command[1]);
    return "Topic" + command[1] + " created!"
  }
  return "PONG to: " + client.Id
}
