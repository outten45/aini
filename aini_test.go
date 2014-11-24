package aini

import (
	"strings"
	"testing"
)

var input1 string = `
myhost1

[dbs]
dbhost1
dbhost2

[apps]
my-app-server1
my-app-server2:3000

`

func createHosts(input string) Hosts {
	testInput := strings.NewReader(input)
	v, _ := NewParser(testInput)
	v.Parse()
	return *v
}

func TestGroupExists(t *testing.T) {
	v := createHosts(input1)
	matched := false
	if _, ok := v.Groups["dbs"]; ok {
		matched = true
	}
	if !matched {
		t.Error("Expected to find the group \"dbs\"")
	}
}

func TestHostExistsInGroups(t *testing.T) {
	v := createHosts(input1)
	exportedHosts := map[string][]Host{
		"dbs": []Host{Host{Name: "dbhost1", Port: 22},
			Host{Name: "dbhost2", Port: 22}},
		"ungrouped": []Host{Host{Name: "myhost1", Port: 22}},
		"apps":      []Host{Host{Name: "my-app-server2", Port: 3000}},
	}

	for group, ehosts := range exportedHosts {
		for _, ehost := range ehosts {
			if hosts, ok := v.Groups[group]; ok {
				matched := false
				for _, host := range hosts {
					if host.Name == ehost.Name {
						matched = true
						if host.Port != ehost.Port {
							t.Errorf("Host port '%v' does not match expected port of '%v'.\n", host.Port, ehost.Port)
						}
					}
				}
				if !matched {
					t.Errorf("Server '%+v' was not found in '%+v'.\n", ehost.Name, group)
				}
			} else {
				t.Errorf("'%v' group doesn't exist.\n", group)
			}
		}

	}
}
