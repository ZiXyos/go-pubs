package types

import "zixyos/goedges/pkg/client"

type InternalCommandFunc func([]string) error
type CommandFunc func(*client.Client, []string) string
