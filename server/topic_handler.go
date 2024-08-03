package server

import (
	"errors"
	"fmt"
)

func (s *Server) FindTopic(topicId string) (*Topic, error) {
  fmt.Println("[LOG::TOPIC::FIND_TOPIC]: ->", topicId)
  fmt.Println(s.topic)
	if v, ok := s.topic[topicId]; ok {
    fmt.Println("Topic: ", v, ok, " Foud")
    return v, nil
	}
  fmt.Println("NOT FOUND")
  return nil, errors.New("No topic found with the topic: " + topicId) 
}
