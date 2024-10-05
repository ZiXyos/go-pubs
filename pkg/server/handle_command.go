package server

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
	"zixyos/goedges/pkg/client"
  "zixyos/goedges/utils"
)

func (s *Server) receive_command(client *client.Client) {
    defer client.Conn.Close();
    client.Conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
    client.Conn.Write([]byte("you can now send a command\n"));
    reader := bufio.NewReader(client.Conn);

    for {
      client.Conn.SetReadDeadline(time.Now().Add(5 * time.Minute));
      message, err := reader.ReadString('\n');
      if err != nil {
        if err == io.EOF {
          s.logger.Infof("Client %s disconnected \n", client.Id);
        }
        break
      }

      message = strings.TrimSpace(message);
      response := s.handle_command(client, message) 
      client.Conn.SetWriteDeadline(time.Now().Add(5 * time.Minute));
      _, err = client.Conn.Write([]byte(response + "\n"));
      if err != nil {
        s.logger.Errorf("Error writing to client %s: %v\n", client.Id, err);
        break
      }
    }
}

func (s *Server) handle_command(client *client.Client, entry string) string {
  commands := strings.Fields(entry);
  if len(commands) == 0 {
    return "Error: Empty command"
  }

  command := strings.ToUpper(commands[0]);
  action := *s.commandsList[command]
  return action(client, commands);
}

func (s *Server) handle_subscribe(client *client.Client, command []string) string {
  if len(command) != 2 {
    return "Error SUB command require 1 argument: topic_name. check usage with the 'HELP' command."
  }
  topicId := command[1]

  client.Mut.Lock()
  topic, err := s.FindTopic(topicId);
  client.Mut.Unlock()

  if err != nil {
    s.logger.Errorf("%v", err)
    return fmt.Sprintf("Error: %v\n", err)
  }

  client.Mut.Lock()
  if err := topic.addSubscriber(client.Id); err != nil {
    s.logger.Errorf("%v", err)
    return fmt.Sprintf("Error: %v\n", err)
  }

  s.logger.Infof("Client %s subscribed to topic %s\n", client.Id, topic.TopicId);
  client.Mut.Unlock();

  return fmt.Sprintf("Client %s, Subscribed to topic: %s\n", client.Id, topic.TopicId)
}

func (s *Server) handle_create(client *client.Client, command []string) string {
  if len(command) != 2 {
    return "Error: CREATE command require 1 argument: topic_name. Check usage with the 'HELP' command."
  }
  topicId := command[1]

  s.mutex.RLock();
  topic, err := s.FindTopic(topicId);
  s.mutex.RUnlock();
  if err == nil && topic != nil {
    return fmt.Sprintf("Error: Topic '%s' already exists\n", topic.TopicId)
  }

  topic, err = s.createTopic(client.Id, topicId);
  if err != nil {
    return fmt.Sprintf("Error creating topic: %v\n", err)
  }

  return fmt.Sprintf("Topic: '%s' created successfully\n", topic.TopicId)
}

func (s *Server) handle_publish(client *client.Client, commands []string) string {
  if len(commands) < 3 {
    return "Error: PUB command require 2 arguments: topic_name, message. Check usage with the 'HELP' command."
  }
  topicId := commands[1];
  message, others := utils.MessageParser(commands[2:]);
  s.logger.Debug(message, others);

  s.mutex.RLock();
  topic, err := s.FindTopic(topicId);
  s.mutex.RUnlock();
  if err != nil {
    return fmt.Sprintf("Error: Topic '%s' not found. check the topic list using the 'LIST' command\n", topicId)
  }

  res, err := s.publishMessage(client.Id, topic.TopicId, message);
  if err != nil {
    return fmt.Sprintf("Error: '%v' ", err)
  }

  return fmt.Sprintf(res)
}
