package server

import (
	"errors"
	"slices"
	"zixyos/goedges/utils"
)

type Topic struct {
  CreatorId string
  TopicId string
  publisher []string   // clientId
  subscriber []string  // clientId
}

func NewTopic(clientId string, topicId string) *Topic {
  return &Topic{
    CreatorId: clientId,
    TopicId: topicId,
    publisher: make([]string, 10),
    subscriber: make([]string, 10),
  }
}

func (topic *Topic) addPublisher(clientId string) {
  utils.SortedInsert(&topic.publisher, clientId)
}

func (topic *Topic) removePublisher(clientId string) error {
  v, found := slices.BinarySearch(topic.publisher, clientId);
  if !found {
    return errors.New("publisher not found")
  }

  topic.publisher = slices.Delete(topic.publisher, v, v+1);
  return nil
}


func (s *Server) createTopic(clientId string, topicId string) {
  topic := NewTopic(clientId, topicId);
  topic.addPublisher(clientId);
  s.topic[clientId] = topic;
}

