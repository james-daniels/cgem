package order

const (
	newOrderEndpoint  = "/v1/order/new"
	priceFeedEndpoint = "/v1/pricefeed"
)

var (
	oType   = "exchange limit"
	options = "immediate-or-cancel"
)
