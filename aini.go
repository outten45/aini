package aini

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/flynn/go-shlex"
)

type Hosts struct {
	input  *bufio.Reader
	Groups map[string][]Host
}

type Host struct {
	Name string
	Port int
}

// type Group struct {
// 	Name    string
// 	Servers []string
// }

func NewParser(r io.Reader) (*Hosts, error) {
	input := bufio.NewReader(r)
	hosts := &Hosts{input: input}
	return hosts, nil
}

func (h *Hosts) Parse() error {
	scanner := bufio.NewScanner(h.input)

	activeGroupName := "ungrouped"
	h.Groups = make(map[string][]Host)
	h.Groups[activeGroupName] = make([]Host, 0)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		// fmt.Println(line)

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			activeGroupName = strings.Replace(strings.Replace(line, "[", "", -1), "]", "", -1)
			if _, ok := h.Groups[activeGroupName]; !ok {
				h.Groups[activeGroupName] = make([]Host, 0)
			}
		} else if strings.HasPrefix(line, ";") || line == "" {
			// do nothing
		} else if activeGroupName != "" {
			parts, err := shlex.Split(line)
			if err != nil {
				fmt.Println("couldn't tokenizer ", line)
			}
			host := Host{Name: parts[0]}
			h.Groups[activeGroupName] = append(h.Groups[activeGroupName], host)
		}
	}
	return nil
}
