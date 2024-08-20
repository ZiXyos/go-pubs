package server

import (
	"errors"
	"fmt"
	"slices"
	"sync"
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
    publisher: make([]string, 0,  10),
    subscriber: make([]string, 0, 10),
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

func (t *Topic) isPublihser(clientId string) (int, bool) {
  pos, ok := slices.BinarySearch(t.publisher, clientId);
  if !ok {
    return -1, false
  }
  return pos, ok
}

func (s *Server) createTopic(clientId string, topicId string) (*Topic, error){
  topic := NewTopic(clientId, topicId);
  err := topic.addPublisher(clientId);
  if err != nil {
    return nil, errors.New("error on publisher adding")
  }

  s.topic[topicId] = topic;
  fmt.Println("topic created: ", s.topic[topicId])
  return s.topic[topicId], nil
}

func (s * Server) publishMessage(
  clientId string,
  topicId string,
  message string,
) (string, error) {
  s.mutex.RLock();
  topic, err := s.FindTopic(topicId);
  s.mutex.RUnlock();
  if err != nil {
    return "", err
  }

  fmt.Println("Topic founded ", topic.TopicId);
  if _, ok := topic.isPublihser(clientId); !ok {
    return "", errors.New("Current client: " + clientId + " is not a publisher of topic " + topicId)
  }

  formatted_message, err := utils.Format_message(clientId, message);
  if err != nil {
    return "", err
  }
  var localWg sync.WaitGroup;
  for _, v := range topic.subscriber {
    localWg.Add(1);
    go func(subId string) {
      defer localWg.Done();
      s.sendMessage(subId, string(formatted_message));
    }(v)
  }

  localWg.Wait();
  return "Message sent to all Subscribers!", nil
}
