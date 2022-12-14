package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"profile_parser/pkg/node_reader"
	"profile_parser/pkg/parser"
)

func main() {
	var txtPath, jsonPath string

	flag.StringVar(&txtPath, "input_txt_path", "", "input path (txt file with nodes)")
	flag.StringVar(&jsonPath, "input_json_path", "", "input path (json file with nodes info)")
	flag.Parse()

	if txtPath == "" {
		err := errors.New("txt input path not specified")
		fmt.Println(err)
		return
	}
	if jsonPath == "" {
		err := errors.New("json input path not specified")
		fmt.Println(err)
		return
	}

	file, err := os.Open(txtPath)
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
	fmt.Println(nodes)

	jsonBuf, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	nodesInfo, err := parser.ParseJSON(jsonBuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nodesInfo)

	timer := make(map[node_reader.Node]int)
	for node := range nodes {
		timer[node] = 0
	}
}
