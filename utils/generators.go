package utils

import (
	"go.step.sm/crypto/randutil"
)

func GenerateRandom() string {
  rand, err := randutil.UUIDv4();
  if err != nil {
    return ""
  }

  return rand
}
