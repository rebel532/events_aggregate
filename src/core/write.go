package core

import (
	"encoding/json"
	"fmt"
	"os"
)

// It will wirte reports to the output file
func writeReports(outputFile string, reports []*Report) error {
	data := []map[string]interface{}{}
	for _, v := range reports {
		x := map[string]interface{}{}
		x["userId"] = v.UserId
		x["date"] = v.Date
		for key, val := range v.EventCounts {
			x[key] = val
		}
		data = append(data, x)
	}
	reportJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(outputFile, reportJSON, 0644)
	if err != nil {
		fmt.Println("error in writting report")
	}
	return err
}
