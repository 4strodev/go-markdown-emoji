package emoji

import (
	"bytes"

	"github.com/gomarkdown/markdown/ast"
)

// Node is a node containing an emoji
type Node struct {
	ast.Leaf
}

var seen map[string]bool

func init() {
	seen = make(map[string]bool)
}

// Parser is a Markdown parser hook to parse emojis.
// This operates as a pre-parser, because gomarkdown
// does not allow inline parser extensions.
func Parser(data []byte) (ast.Node, []byte, int) {
	if seen[string(data)] {
		// Already processed
		return nil, nil, 0
	}
	if bytes.Contains(data, []byte("class=\"emoji\"")) {
		// Already processed
		return nil, nil, 0
	}
	dataLen := len(data)
	if dataLen <= 1 {
		// Not long enough to be an emoji
		return nil, nil, 0
	}
	if bytes.IndexByte(data, ':') == -1 {
		// No emoji delimieters
		return nil, nil, 0
	}
	// Translate emojis to HTML
	resData := make([]byte, 0)
	startIndex := bytes.IndexByte(data, ':')
	resData = append(resData, data[0:startIndex]...)
	for {
		if startIndex >= len(data) {
			// Done
			break
		}
		endIndex := bytes.IndexByte(data[startIndex+1:], ':') + startIndex + 1
		if endIndex > startIndex {
			name := string(data[startIndex+1 : endIndex])
			if isValidEmoji([]byte(name)) {
				startIndex = endIndex + 1
				emoji := getEmoji(name)
				resData = append(resData, []byte(emoji)...)
			} else {
				resData = append(resData, data[startIndex:endIndex]...)
				startIndex = endIndex
			}
			if startIndex == dataLen {
				break
			}
		} else {
			break
		}
	}
	if startIndex < dataLen {
		resData = append(resData, data[startIndex:]...)
	}

	if !bytes.Contains(resData, []byte("class=\"emoji\"")) {
		// Processed with no changes
		seen[string(resData)] = true
	}

	return &ast.Softbreak{}, resData, dataLen
}
