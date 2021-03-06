package types

import (
	"fmt"
	"strings"
)

const (
	APIVersion = "1.0"
	QuoteStr   = "quote"
	NewsStr    = "news"
	ChartStr   = "chart"
	StockStr   = "stock"
	PriceStr   = "price"
	BatchStr   = "batch"
	LastStr    = "last"
	MrktStr    = "market"
	APIURL     = "https://api.iextrading.com/"
)

// Quote repesents the format returned for a quote from IEX(https://iextrading.com)
type Quote struct {
	Symbol           string  `json:symbol`
	CompanyName      string  `json:companyName`
	PrimaryExchange  string  `json:primaryExchange`
	CalculationPrice string  `json:calculationPrice`
	IexRealtimePrice float64 `json:iexRealtimePrice`
	IexRealtimeSize  float64 `json:iexRealtimeSize`
	IexLastUpdated   float64 `json:iexLastUpdated`
	DelayedPrice     float64 `json:delayedPrice`
	DelayedPriceTime float64 `json:delayedPriceTime`
	PreviousClose    float64 `json:previousClose`
	Change           float64 `json:change`
	ChangePercent    float64 `json:changePercent`
	IexMarketPercent float64 `json:iexMarketPercent`
	IexVolume        float64 `json:iexVolume`
	AvgTotalVolume   float64 `json:avgTotalVolume`
	IexBidPrice      float64 `json:iexBidPrice`
	IexBidSize       float64 `json:iexBidSize`
	IexAskPrice      float64 `json:iexAskPrice`
	IexAskSize       float64 `json:iexAskSize`
	MarketCap        float64 `json:marketCap`
	//PeRatio          float64 `json:peRatio`
	Week52High float64 `json:week52High`
	Week52Low  float64 `json:week52Low`
}

// News is the news structure returned from IEX
type News struct {
	DateTime string `json:datetime`
	Headline string `json:headline`
	Source   string `json:source`
	URL      string `json:url`
	Summary  string `json:summar`
	Related  string `json:related`
}

// Batch is a []Quote
type Batch map[string]map[string]Quote

// Quote returns the quote in a iex batch for a specific ticker
// returns error if symbol does not exist
func (i Batch) Quote(ticker string) (Quote, error) {
	ticker = strings.ToUpper(ticker)
	t, ok := i[ticker]
	if !ok {
		return Quote{}, fmt.Errorf("Failed to find %v in batch request", ticker)
	}

	q, ok := t[QuoteStr]
	if !ok {
		return Quote{}, fmt.Errorf("Failed to find quote for %v in batch request", ticker)
	}

	return q, nil
}
