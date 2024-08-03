package server

import (
	"errors"
	"fmt"
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

func (topic *Topic) addPublisher(clientId string) error {
  i, ok := slices.BinarySearch(topic.publisher, clientId);
  if ok {
    fmt.Println(i, ok)
    return errors.New("User " + clientId + " already a publisher")
  }
  utils.SortedInsert(&topic.publisher, clientId);
  return nil
}

func (topic *Topic) removePublisher(clientId string) error {
  v, found := slices.BinarySearch(topic.publisher, clientId);
  if !found {
    return errors.New("publisher not found")
  }

  topic.publisher = slices.Delete(topic.publisher, v, v+1);
  return nil
}

func (topic *Topic) addSubscriber(clientId string) (error) {
  _, ok := slices.BinarySearch(topic.subscriber, clientId);
  if ok {
    return errors.New("User " + clientId + " already subscribed");
  }
  
  utils.SortedInsert(&topic.subscriber, clientId);
  return nil
}

func (topic *Topic) removeSubscriber(clientId string) error {
  pos, ok := slices.BinarySearch(topic.subscriber, clientId);
  if !ok {
    return errors.New(
      "User " + clientId + " not found",
    )
  }

  topic.publisher = append(topic.publisher[:pos], topic.publisher[pos+1:]...)
  return nil
}

func (s *Server) createTopic(clientId string, topicId string) {
  topic := NewTopic(clientId, topicId);
  err := topic.addPublisher(clientId);
  if err != nil {
    fmt.Println("error on publisher adding")
    return
  }

  s.topic[topicId] = topic;
  fmt.Println("topic created: ", s.topic[topicId])
}

