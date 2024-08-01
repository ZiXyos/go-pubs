package server

type Topic struct {
  CreatorId string
  TopicId string
  Publisher []string   // clientId
  Subscriber []string  // clientId
}

func NewTopic(clientId string, topicId string) *Topic {
  return &Topic{
    CreatorId: clientId,
    TopicId: topicId,
    Publisher: make([]string, 10),
    Subscriber: make([]string, 10),
  }
} 
