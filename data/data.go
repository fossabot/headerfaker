package data

import (
	"encoding/json"
	"github.com/aURORA-JC/headerfaker/model"
	"github.com/aURORA-JC/headerfaker/util"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

var (
	TestData []model.Test
	FlagData map[string]string
)

// ReadData Read tests' TestData from TestData.json
func ReadData() {
	// Read TestData from file
	file, readError := ioutil.ReadFile("./database/data.json")
	if readError != nil {
		color.Red("Read file error!")
		os.Exit(10)
	}

	// Unmarshal JSON TestData
	unmarshalError := json.Unmarshal([]byte(file), &TestData)
	if unmarshalError != nil {
		color.Red("Illegal json")
		os.Exit(11)
	}

	// Generate FlagData
	FlagData = make(map[string]string)
	for _, test := range TestData {
		FlagData[test.TestTag] = util.GenerateFlag(test.TestTitle)
	}
}
