package utility

import (
	"csv-to-html-converter/model"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Find given value and return its index from array
func IndexOf(word string, data []string) int {
	for idx, val := range data {
		if word == val {
			return idx
		}
	}
	return -1
}

// Match given value with given Regex pattern
func ExceuteRegexPattern(val string, reg string) bool {
	match, _ := regexp.MatchString(reg, val)
	return match
}

// Match regex in slice and return matched value
func FindSliceElementByRegex(arr []string, reg string) (val string) {
	val = ""
	for _, i := range arr {
		m, _ := regexp.MatchString(reg, i)
		if m {
			val = i
		}
	}
	return val
}

// Format given string by removing its unwanted quotes, comma's using regex
func FormatString(str string) string {
	firstStr := regexp.MustCompile(`"`).ReplaceAllString(str, "")
	fomattedStr := regexp.MustCompile(`,`).ReplaceAllString(firstStr, "")
	return fomattedStr
}

func IsInputFilePathExists(inputFileObj model.InputFileObj) bool {
	return inputFileObj.Path != "" && inputFileObj.Name != ""
}

// Get file path and file name from input options
func GetFilePathAndNameFromOptions(options model.Options) model.InputFileObj {
	if options.FilePath != "" && ExceuteRegexPattern(options.FilePath, CSV_FILE_TYPE) {
		strings.Split(options.FilePath, "/")
		splitInputFileName := strings.Split(options.FilePath, "/")
		var name string
		if len(splitInputFileName) > 0 {
			name = strings.ReplaceAll(splitInputFileName[len(splitInputFileName)-1], ".csv", "")
		}

		return model.InputFileObj{Path: options.FilePath, Name: name}
	} else {
		return model.InputFileObj{Path: "", Name: ""}
	}
}

// Get seperated file path and file name from cli input arguments
func GetFilePathAndNameFromCli(inputArgs []string) model.InputFileObj {
	filePathIdx := IndexOf(FILEPATH_PATH, inputArgs)
	path := inputArgs[filePathIdx+1]
	if filePathIdx >= 0 && ExceuteRegexPattern(path, CSV_FILE_TYPE) && inputArgs[filePathIdx+1] != "" {
		splitInputFileName := strings.Split(inputArgs[filePathIdx+1], "/")
		var name string
		if len(splitInputFileName) > 0 {
			name = strings.ReplaceAll(splitInputFileName[len(splitInputFileName)-1], ".csv", "")
		}
		return model.InputFileObj{Path: path, Name: name}
	} else {
		return model.InputFileObj{Path: "", Name: ""}
	}
}

// Get formated table headers
func GetFormatedHeader(headers []string) []string {
	formatedHeadVal := make([]string, 0)
	if len(headers) > 0 {
		formatedHeadVal = append(formatedHeadVal, FormatString(headers[0]))
	}
	for headIdx, headVal := range headers {
		if ExceuteRegexPattern(headVal, FORMAT_TABLE_VAL_REGEX) {
			formatedHeadVal = append(formatedHeadVal, FormatString(headVal))
		} else {
			subHeader := strings.Split(headVal, ",")
			if len(formatedHeadVal) > 0 && len(subHeader) > 0 && headIdx > 0 {
				formatedHeadVal[len(formatedHeadVal)-1] = formatedHeadVal[len(formatedHeadVal)-1] + " " + subHeader[0]
				for idx, sub := range subHeader {
					if idx > 0 {
						formatedHeadVal = append(formatedHeadVal, FormatString(sub))
					}
				}
			}
		}

	}
	return formatedHeadVal
}

// Get formated table body
func GetFormatedBody(lines []string, headerEndIdex int) []string {
	formatedBodyVal := make([]string, 0)
	for idx, val := range lines {
		if idx > headerEndIdex {
			splitedRowData := strings.Split(val, `",`)
			for _, splitVal := range splitedRowData {
				if ExceuteRegexPattern(splitVal, FORMAT_TABLE_VAL_REGEX) {
					formatedBodyVal = append(formatedBodyVal, regexp.MustCompile(`"`).ReplaceAllString(splitVal, ""))
				} else {
					tempSplitVal := strings.Split(splitVal, `,"`)
					for _, t := range tempSplitVal {
						formatedBodyVal = append(formatedBodyVal, t)
					}
				}
			}
		}
	}
	return formatedBodyVal
}

// Construct JSON from headears and body array value
func ConstructJson(headers []string, body []string) ([]model.Keyvalue, []string) {
	jsonBuild := make(model.Keyvalue)
	jsonArr := make([]model.Keyvalue, 0)
	marshalJsonArr := make([]string, 0)
	for bodyIdx, bodyVal := range body {
		headIdx := 0
		if bodyIdx < len(headers) {
			headIdx = bodyIdx
		} else {
			headIdx = (bodyIdx % len(headers))
		}
		if bodyVal != "" {
			jsonBuild[headers[headIdx]] = fmt.Sprint(bodyVal)
		}
		if headIdx == len(headers)-1 {
			copyJsonBuild := make(model.Keyvalue)
			for k, v := range jsonBuild {
				copyJsonBuild[k] = v
			}
			temp := model.StockMarketData{
				Symbol:                fmt.Sprint(body[bodyIdx-13]),
				Open:                  fmt.Sprint(copyJsonBuild["OPEN "]),
				High:                  fmt.Sprint(copyJsonBuild["HIGH "]),
				Low:                   fmt.Sprint(copyJsonBuild["LOW "]),
				PreviousClose:         fmt.Sprint(copyJsonBuild["PREV. CLOSE "]),
				Ltp:                   fmt.Sprint(copyJsonBuild["LTP "]),
				Change:                fmt.Sprint(copyJsonBuild["CHNG "]),
				PercentageChange:      fmt.Sprint(copyJsonBuild["%CHNG "]),
				Volume:                fmt.Sprint(copyJsonBuild[`VOLUME  (shares)"`]),
				Value:                 fmt.Sprint(copyJsonBuild["VALUE "]),
				YearHigh:              fmt.Sprint(copyJsonBuild["52W H "]),
				YearLow:               fmt.Sprint(copyJsonBuild["52W L "]),
				YearPercentageChange:  fmt.Sprint(copyJsonBuild[FindSliceElementByRegex(headers, YEAR_PERCENTAGE_REG)]),
				YonthPercentageChange: fmt.Sprint(copyJsonBuild[FindSliceElementByRegex(headers, MONTH_PERCENTAGE_REG)]),
			}
			bytes, _ := json.Marshal(temp)
			marshalJsonArr = append(marshalJsonArr, string(bytes))
			jsonArr = append(jsonArr, copyJsonBuild)
			for key, _ := range jsonBuild {
				delete(jsonBuild, key)
			}
		}
	}
	return jsonArr, marshalJsonArr
}

func constructTableHeader(tHeaders []string) string {
	var th string
	for _, item := range tHeaders {
		th = th + `<th>` + item + `</th>`
	}
	return th
}

func constructedTableData(mapData model.Keyvalue, tHeaders []string) string {
	var td string
	for _, v := range tHeaders {
		td = td + `<td>` + mapData[v].(string) + `</td>`
	}
	return td
}

func ConstructHtml(csvToJSON []model.Keyvalue, tHeaders []string) string {
	var tr string
	for _, mapData := range csvToJSON {
		tr = tr + `<tr>` + constructedTableData(mapData, tHeaders) + `</tr>`
	}
	return `<html>
        <head>
        <h3>CSV TO HTML TABLE</h3>
        <style>
        table {
            font-family: arial, sans-serif;
            /* border-collapse: collapse; */
            width: 100%;
        }
        td,
        th {
            border: 1px solid #dddddd;
            text-align: left;
            padding: 10px;
        }
        tr:nth-child(even) {
            background-color: #dddddd;
        }
        </style>
        </head>
        <body>
        <table>
        <tr>` + constructTableHeader(tHeaders) + `</tr>` + tr +
		`</table>
        </body>
        </html>`
}

func CreateOutputFolder() {
	_, err := os.Stat(OUTPUT_DIR_NAME)
	if err != nil {
		os.Mkdir(OUTPUT_DIR_NAME, os.ModePerm)
	}
}
