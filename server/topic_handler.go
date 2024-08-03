package server

import (
	"errors"
)

func (s *Server) FindTopic(topicId string) (*Topic, error) {
	if v, ok := s.topic[topicId]; ok {
    return v, nil
	}

  return nil, errors.New("No topic found with the topic: " + topicId) 
}
