package kvlm

import (
	"fmt"
	"strings"
)

// Key-Value List with Message
type KWLM struct {
	Tree     string
	Parent   string
	Author   string
	Comitter string
	GPGSIG   string
	Message  string
}

func validKeys() []string {
	return []string{"tree", "parent", "author", "committer", "gpgsig"}
}

func (k *KWLM) Deserialize(data string) {
	for data != "" {
		space := strings.Index(data, " ")
		newLine := strings.Index(data, "\n")

		if newLine == 0 || space > newLine {
			k.Message = strings.TrimSpace(data[1:])
			data = ""
			continue
		}
		key := data[0:space]
		data = data[space:]

		value := ""
		for {
			end := strings.Index(data, "\n")
			value = value + data[:end+1]
			if data[end+1] != byte(' ') {
				data = data[end+1:]
				break
			}
			data = data[end+1:]
		}

		value = strings.TrimSpace(value)
		switch key {
		case "tree":
			k.Tree = value
		case "parent":
			k.Parent = value
		case "author":
			k.Author = value
		case "committer":
			k.Comitter = value
		case "gpgsig":
			k.GPGSIG = strings.ReplaceAll(value, "\n ", "\n")
		}
	}
}

func (k *KWLM) Serialize() string {
	value := fmt.Sprintf("%s %s\n", "tree", k.Tree)
	value = value + fmt.Sprintf("%s %s\n", "parent", k.Parent)
	value = value + fmt.Sprintf("%s %s\n", "author", k.Author)
	value = value + fmt.Sprintf("%s %s\n", "committer", k.Comitter)
	value = value + fmt.Sprintf("%s %s\n", "gpgsig", strings.ReplaceAll(k.GPGSIG, "\n", "\n "))
	value = value + fmt.Sprintf("\n%s\n", k.Message)
	return value
}
