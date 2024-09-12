package types;

type MessageModel[K comparable] struct {
  PublisherId string `json:"publisher_id"`
  Message K `json:"message"`
}
