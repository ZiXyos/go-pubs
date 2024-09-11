package utils

import (
	"strings"
)

/*
** check if message list contain double quotes
** send content only in double quote
** send firs arg if not
*/

// ["this, is, as, test"]
func MessageParser(message []string) (string, []string) {
  if (strings.Contains(message[0], "\"")) {
    mark := 0;
    response := make([]string, 0, 4096);
    for k, v := range message {
      if (strings.Contains(v, "\"")) {
        mark++;
      }
      response = append(response, v);
      if mark == 2 {
        return strings.Join(response, " "), message[k:]
      } 

    }
  }
  return message[0], message[1:] 
}
