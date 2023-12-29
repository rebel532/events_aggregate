package core

import (
	"encoding/json"
	"os"

	"github.com/spf13/cast"
)

func readEvents(inputFile string) ([]*Event, error) {
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}
	var events []*Event
	err = json.Unmarshal(fileContent, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func readReports(reportPath string) (reports []*Report, err error) {
	existingReports := []map[string]interface{}{}
	fileContent, err := os.ReadFile(reportPath)
	if err != nil {
		return reports, err
	}

	err = json.Unmarshal(fileContent, &existingReports)
	if err != nil {
		return nil, err
	}
	reports = []*Report{}
	for _, v := range existingReports {
		rp := Report{}
		rp.Counts = map[string]int{}
		for key, val := range v {
			if key == "userId" {
				rp.UserId = cast.ToInt(val)
			} else if key == "date" {
				rp.Date = cast.ToString(val)
			} else {
				rp.Counts[key] = cast.ToInt(val)
			}
		}
		reports = append(reports, &rp)
	}
	return reports, nil
}
