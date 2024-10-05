package server

import "zixyos/goedges/pkg/client"

func (serv *Server) ListClients(client *client.Client, command []string) string {
  clientCounter := 0;
  for k, _ := range serv.client {
    serv.logger.Infof("Client: %s is Connected", k)
    clientCounter++;
  }
  serv.logger.Infof("There is currentlu %d, clients connected on the server", clientCounter);
  return ""
}

func (s *Server) ListSubbedClient(client *client.Client, command []string) string {
  if len(command) != 2 {
    s.logger.Error("Missing Argument")
    return "Not so much argument"
  }

  s.logger.Info(command)
	topic, err := s.FindTopic(command[1])
	if err != nil {
    return ""
	}

	for _, v := range topic.subscriber {
		s.logger.Infof("Client: %s, is subbed to topic: %s", v, topic.TopicId)
	}
  return ""
}

func (s *Server) ListTopic(client *client.Client, command []string) string {
	for k, _ := range s.topic {
		s.logger.Infof("Topic: %s\n", k)
	}
  return ""
}
