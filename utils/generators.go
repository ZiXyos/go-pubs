package utils

import (
	"fmt"

	"go.step.sm/crypto/randutil"
)

func GenerateRandom() string {
  rand, err := randutil.UUIDv4();
  if err != nil {
    return ""
  }

  return rand
}

func GenerateCommand(newCommand string, commandList *[]string) {
  fmt.Println("INIT SERVER COMMAND")
  *commandList = append(*commandList, newCommand);
}
