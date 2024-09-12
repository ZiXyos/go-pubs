package utils

import (
	"bytes"
	"strings"
)

func MessageParser(message []string) (string, []string) {
	if len(message) == 0 {
		return "", nil
	}

	fullMsg := []byte(strings.Join(message, " "))

	var start, end int
	inQuote := false
	inBraces := false
	braceCount := 0
	escapeNext := false

	for k, v := range fullMsg {
		switch v {
		case '"':
			if !inBraces && !inQuote {
				start = k + 1
				inQuote = true
			} else if inQuote && !escapeNext {
				end = k
				if end < len(fullMsg) {
					return string(fullMsg[start:end]), remainingMessage(fullMsg[end+1:])
				}
				return string(fullMsg[start:end]), nil
			}
		case '{':
			if !inQuote && !inBraces {
				start = k
				inBraces = true
				braceCount++
			}
		case '}':
			if inBraces {
				braceCount--
				if braceCount == 0 {
					end = k + 1
					if end < len(fullMsg) {
						return string(fullMsg[start:end]), remainingMessage(fullMsg[end+1:])
					}
					return string(fullMsg[start:end]), nil
				}
			}
		case '\\':
			escapeNext = !escapeNext
		default:
			escapeNext = false
			if !inBraces && !inQuote && v != ' ' {
				end = bytes.IndexByte(fullMsg[k:], ' ')
				if end == -1 {
					end = len(fullMsg)
				} else {
					end += k
				}
				if end < len(fullMsg) {
					return string(fullMsg[k:end]), remainingMessage(fullMsg[end+1:])
				}
				return string(fullMsg[k:end]), nil
			}
		}
	}

	if inBraces || inQuote {
		return string(fullMsg[start:]), nil
	}

	return "", nil
}

func remainingMessage(fullMsg []byte) []string {
	trimmed := bytes.TrimSpace(fullMsg)
	if len(trimmed) == 0 {
		return nil
	}
	return strings.Fields(string(trimmed))
}
