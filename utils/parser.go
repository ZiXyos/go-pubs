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
    if len(message) == 0 {
        return "", nil
    }

    var quotedParts []string
    insideQuotes := false

    for k, v := range message {
        if strings.Contains(v, "\"") {
            if insideQuotes {
                quotedParts = append(quotedParts, strings.TrimRight(v, "\""))
                joinedQuote := strings.Join(quotedParts, " ")
                return strings.TrimLeft(joinedQuote, "\""), message[k+1:]
            } else {
                insideQuotes = true
                quotedParts = append(quotedParts, strings.TrimLeft(v, "\""))
            }
        } else if insideQuotes {
            quotedParts = append(quotedParts, v)
        }
    }

    return message[0], message[1:]
}
