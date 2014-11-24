package aini

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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
		// fmt.Println(activeGroupName, ":", line)

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			replacer := strings.NewReplacer("[", "", "]", "")
			activeGroupName = replacer.Replace(line)

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
			host := getHost(parts)
			h.Groups[activeGroupName] = append(h.Groups[activeGroupName], host)
		}
	}
	return nil
}

func getHost(parts []string) Host {
	hostname := parts[0]
	port := 22
	if (strings.Contains(hostname, "[") &&
		strings.Contains(hostname, "]") &&
		strings.Contains(hostname, ":") &&
		(strings.LastIndex(hostname, "]") < strings.LastIndex(hostname, ":"))) ||
		(!strings.Contains(hostname, "]") && strings.Contains(hostname, ":")) {

		splithost := strings.Split(hostname, ":")
		if i, err := strconv.Atoi(splithost[1]); err == nil {
			port = i
		}
		hostname = splithost[0]
	}
	host := Host{Name: hostname, Port: port}

	return host
}
