package aini

import (
	"strings"
	"testing"
)

//
func TestTokenizer(t *testing.T) {
	i := "myhost1\n[dbs]\ndbhost1"
	testInput := strings.NewReader(i)
	v, _ := NewParser(testInput)
	v.Parse()
}
