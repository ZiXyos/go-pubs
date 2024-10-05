package utils

import (
	"encoding/json"
	"errors"
	"zixyos/goedges/pkg/types"
)

// this should be an internal thing i guess
func Format_message[K string](
  sender string,
  message string,
) ([]byte, error) {
  msg := types.MessageModel[K]{
    PublisherId: sender,
    Message: K(message),
  };

  jsonBytes, err := json.Marshal(msg);
  if err != nil {
    return nil, err 
  }
  return jsonBytes, nil
}

func parse_response[K comparable](message []byte) (struct{}, error) {
  return struct{}{}, errors.New("method not implemented")
}
