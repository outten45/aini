package aini

import (
	"bufio"
	"fmt"
	"io"
)

type Hosts struct {
	input   *bufio.Reader
	Groups  []Group
	Servers []Server
}

type Server struct {
	Name string
}

type Group struct {
	Name    string
	Servers []Server
}

func NewParser(r io.Reader) (*Hosts, error) {
	input := bufio.NewReader(r)
	hosts := &Hosts{input: input}
	return hosts, nil
}

func (h *Hosts) Parse() error {
	scanner := bufio.NewScanner(h.input)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
