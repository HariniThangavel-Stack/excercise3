package utility

import (
	"csv-to-html-converter/model"
	"reflect"
	"testing"
)

func TestIndexOf(t *testing.T) {
	type FindIndex struct {
		name   string
		find   string
		arr    []string
		result int
	}
	tt := []FindIndex{
		{"Index of two should be 1", "two", []string{"one", "two", "threee"}, 1},
		{"Index of four should be -1", "four", []string{"one", "two", "threee"}, -1},
		{"Index of one should be 0", "one", []string{"one", "two", "threee"}, 0},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res := IndexOf(ti.find, ti.arr)
			if ti.result != res {
				t.Errorf("Expected %d actual %d \n", ti.result, res)
			}
		})
	}
}

func TestExceuteRegexPattern(t *testing.T) {
	type ExecuteRegex struct {
		name   string
		val    string
		reg    string
		result bool
	}
	tt := []ExecuteRegex{
		{"Pattern should match and return true", "sample.csv", `.csv$`, true},
		{"Pattern should not match and return false", "sample.csvv", `.csv$`, false},
		{`Word ends with " and return true`, `test"`, `(\w+)"$`, true},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res := ExceuteRegexPattern(ti.val, ti.reg)
			if ti.result != res {
				t.Errorf("Expected %t actual %t \n", ti.result, res)
			}
		})
	}
}

func TestFormatString(t *testing.T) {
	type ExecuteRegex struct {
		name   string
		str    string
		result string
	}
	tt := []ExecuteRegex{
		{"Format string and return as OPEN", `","OPEN`, `OPEN`},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res := FormatString(ti.str)
			if ti.result != res {
				t.Errorf("Expected %s actual %s \n", ti.result, res)
			}
		})
	}
}

func TestIsInputFilePathExists(t *testing.T) {
	type ExecuteRegex struct {
		name   string
		input  model.InputFileObj
		result bool
	}
	tt := []ExecuteRegex{
		{"Invalid options should return false", model.InputFileObj{Path: "", Name: ""}, false},
		{"Valid options should return true", model.InputFileObj{Path: "testpath", Name: "testName"}, true},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res := IsInputFilePathExists(ti.input)
			if ti.result != res {
				t.Errorf("Expected %t actual %t \n", ti.result, res)
			}
		})
	}
}

func TestGetFilePathAndNameFromOptions(t *testing.T) {
	type ExecuteRegex struct {
		name   string
		input  model.Options
		result model.InputFileObj
	}
	tt := []ExecuteRegex{
		{"Construct input file options", model.Options{OutputFormat: "", FilePath: ""}, model.InputFileObj{Path: "", Name: ""}},
		{"Construct input file options", model.Options{OutputFormat: "html", FilePath: "./examples/sample.csv"}, model.InputFileObj{Path: `./examples/sample.csv`, Name: "sample"}},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res := GetFilePathAndNameFromOptions(ti.input)
			if ti.result != res {
				t.Errorf("Expected %#v actual %#v \n", ti.result, res)
			}
		})
	}
}

func TestGetFilePathAndNameFromCli(t *testing.T) {
	type ExecuteRegex struct {
		name   string
		input  []string
		result model.InputFileObj
	}
	tt := []ExecuteRegex{
		{"Construct input file options", []string{`filepath`, `./examples/MW-NIFTY-BANK-05-Aug-2021.csv`},
			model.InputFileObj{Path: `./examples/MW-NIFTY-BANK-05-Aug-2021.csv`, Name: "MW-NIFTY-BANK-05-Aug-2021"}},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res := GetFilePathAndNameFromCli(ti.input)
			if ti.result != res {
				t.Errorf("Expected %#v actual %#v \n", ti.result, res)
			}
		})
	}
}

func TestGetFormatedHeader(t *testing.T) {
	type ExecuteRegex struct {
		name   string
		input  []string
		result []string
	}
	tt := []ExecuteRegex{
		{"Construct input file options", []string{`SYMBOL`, `","OPEN`},
			[]string{"SYMBOL", "OPEN"}},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res := GetFormatedHeader(ti.input)
			if !reflect.DeepEqual(ti.result, res) {
				t.Errorf("Expected %#v actual %#v \n", ti.result, res)
			}
		})
	}
}

func TestGetFormatedBody(t *testing.T) {
	type ExecuteRegex struct {
		name   string
		input  []string
		result []string
	}
	tt := []ExecuteRegex{
		{"Construct input file options", []string{"NIFTY BANK", "36,093.95", "36,115.45"},
			[]string{"NIFTY BANK", "36,093.95", "36,115.45"}},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res := GetFormatedBody(ti.input, -1)
			if !reflect.DeepEqual(ti.result, res) {
				t.Errorf("Expected %#v actual %#v \n", ti.result, res)
			}
		})
	}
}

func TestConstructJson(t *testing.T) {
	type ExecuteRegex struct {
		name    string
		headers []string
		body    []string
		result  []model.Keyvalue
	}
	tt := []ExecuteRegex{
		{"Construct input file options", []string{`SYMBOL`, `OPEN`, `HIGH`}, []string{"NIFTY BANK", "36,093.95", "36,115.45"},
			[]model.Keyvalue{{"HIGH": "36,115.45", "OPEN": "36,093.95", "SYMBOL": "NIFTY BANK"}}},
	}
	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			res, _ := ConstructJson(ti.headers, ti.body)
			if !reflect.DeepEqual(ti.result, res) {
				t.Errorf("Expected %#v actual %#v \n", ti.result, res)
			}
		})
	}
}
