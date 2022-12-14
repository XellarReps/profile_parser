package csv_writer

import (
	"encoding/csv"
	"fmt"
	"os"
	"profile_parser/pkg/node_reader"
)

func WriteTimeCsv(file *os.File, times map[node_reader.Node]int) error {
	records := [][]string{
		{"Name", "Op_type", "Time"},
	}

	for node, time := range times {
		timeStr := fmt.Sprintf("%d", time)
		records = append(records, []string{node.Name, node.Type, timeStr})
	}

	w := csv.NewWriter(file)
	err := w.WriteAll(records)
	if err != nil {
		return err
	}
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
