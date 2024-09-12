package utils

import (
	"fmt"
	"zixyos/goedges/pkg/types"
	"go.step.sm/crypto/randutil"
)

type GenerateRandom func();
func (gen *GenerateRandom) GenerateRandomString() string {
  rand, err := randutil.UUIDv4();
  if err != nil {
    return ""
  }
  return rand
}

func (gen *GenerateRandom) GenerateRandomId() string {
  return "client_" + gen.GenerateRandomString();
}


func GenerateCommand(newCommand string, commandList *[]string) {
  fmt.Println("INIT NEW SERVER COMMAND: ", newCommand)
  SortedInsert(commandList, newCommand)
}

func GenerateInternalCommandMap(
  newCommand string,
  fun types.InternalCommandFunc,
  commandList *map[string] *types.InternalCommandFunc,
) {
  (*commandList)[newCommand] = &fun;
}

func GenerateCommandMap(
  newCommand string,
  fun types.CommandFunc,
  commandList *map[string] *types.CommandFunc,
) {
  (*commandList)[newCommand] = &fun;
}
