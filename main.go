package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"profile_parser/pkg/node_reader"
)

func main() {
	var path string

	flag.StringVar(&path, "input_path", "", "input path (txt file with nodes)")
	flag.Parse()

	if path == "" {
		err := errors.New("input path not specified")
		fmt.Println(err)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	nodes := node_reader.NodeRead(file)
	for node := range nodes {
		fmt.Println(node.Name, node.Type)
	}
	fmt.Println(len(nodes))
}
