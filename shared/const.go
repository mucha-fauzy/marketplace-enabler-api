package shared

import "time"

const (
	MarketPathField     = "market"
	StartDateQueryField = "dateFrom"
	EndDateQueryField   = "dateTo"
	BrandNameQueryField = "brand"
	StoreNameQueryField = "store"
	PageQueryField      = "page"
	SizeQueryField      = "size"
	DefaultDateFormat   = time.RFC3339
	DefaultPage         = 1
	DefaultSize         = 10
	PageLowerLimit      = 1
	SizeLowerLimit      = 1
	SizeUpperLimit      = 20
	BLIBLI              = "blibli"
	SHOPEE              = "shopee"
)
