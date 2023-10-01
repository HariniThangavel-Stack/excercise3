package model

type Options struct {
	OutputFormat string
	FilePath     string
}

type InputFileObj struct {
	Path string
	Name string
}

type Keyvalue map[string]interface{}

type StockMarketData struct {
	Symbol                string
	Open                  string
	High                  string
	Low                   string
	PreviousClose         string
	Ltp                   string
	Change                string
	PercentageChange      string
	Volume                string
	Value                 string
	YearHigh              string
	YearLow               string
	YearPercentageChange  string
	YonthPercentageChange string
}
