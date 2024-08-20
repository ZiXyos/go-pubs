package utils

import (
	"strings"
)

func MessageParser(message []string) string {
  first, last := -1, -1;
  
    for _, v := range message {
      for i, c := range v {
        if c == '\'' && first == -1 {
          first = i;
        }
        if c == '\'' && first != -1 {
          last = i;
        }
      }
    }

  if first != -1 && last != -1  && first < last {
    res := strings.Join(message[first:last+1], " ");
    return res[1:len(res)-1]
  }

  return ""
}
