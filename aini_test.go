package aini

import (
	"strings"
	"testing"
)

func input1() string {
	return `
myhost1
[dbs]
dbhost1
dbhost2

[apps]
my-app-server1
my-app-server2

`
}

func createHosts(input string) Hosts {
	testInput := strings.NewReader(input)
	v, _ := NewParser(testInput)
	v.Parse()
	return *v
}

//
func TestGroupExists(t *testing.T) {
	v := createHosts(input1())
	matched := false
	if _, ok := v.Groups["dbs"]; ok {
		matched = true
	}
	if !matched {
		t.Error("Expected to find the group \"dbs\"")
	}
}

func TestServerExistsInGroup(t *testing.T) {
	v := createHosts(input1())
	if hosts, ok := v.Groups["dbs"]; ok {
		matched := false
		for _, host := range hosts {
			if host.Name == "dbhost2" {
				matched = true
			}
		}
		if !matched {
			t.Error("Server dbhost2 was not found in dbs")
		}
	} else {
		t.Error("\"dbs\" group didn't exists")
	}
}
