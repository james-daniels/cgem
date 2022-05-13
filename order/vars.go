package order

const (
	newOrderEndpoint  = "/v1/order/new"
	priceFeedEndpoint = "/v1/pricefeed"

	configFile = "config.ini"
)

var (
	apiKey    string
	apiSecret string

	oType   = "exchange limit"
	options = "immediate-or-cancel"
)
