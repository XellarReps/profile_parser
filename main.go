package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"profile_parser/pkg/csv_writer"
	"profile_parser/pkg/node_reader"
	"profile_parser/pkg/parser"
)

func main() {
	var txtPath, jsonPath, outputPath string

	flag.StringVar(&txtPath, "input_txt_path", "", "input path (txt file with nodes)")
	flag.StringVar(&jsonPath, "input_json_path", "", "input path (json file with nodes info)")
	flag.StringVar(&outputPath, "output_path", "", "output path (csv file)")
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
	if outputPath == "" {
		err := errors.New("output path not specified")
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

	timer := make(map[node_reader.Node]int)
	for node := range nodes {
		nodeTime := 0
		for _, info := range nodesInfo {
			if len(node.Name) <= len(info.Name) &&
				node.Name == info.Name[:len(node.Name)] &&
				node.Type == info.Args.Type {
				nodeTime += info.Time
			}
		}
		timer[node] = nodeTime
	}

	outFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = outFile.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	err = csv_writer.WriteTimeCsv(outFile, timer)
	if err != nil {
		fmt.Println(err)
		return
	}
}
