package node_reader

import (
	"bufio"
	"os"
)

type Node struct {
	Name string
	Type string
}

func NodeRead(file *os.File) map[Node]struct{} {
	nodes := make(map[Node]struct{})
	scanner := bufio.NewScanner(file)
	cnt := 0
	var name, opType string
	for scanner.Scan() {
		line := scanner.Text()
		if cnt%2 == 0 { // check string with 'Name:'
			name = line[6:] // because line contains the data type described below 'Name: ...'
		} else { // check string with 'Type:'
			opType = line[6:] // because line contains the data type described below 'Type: ...'
			nodes[Node{Name: name, Type: opType}] = struct{}{}
		}
		cnt += 1
	}
	return nodes
}
