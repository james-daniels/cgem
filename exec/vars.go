package exec

const (
	production = "https://api.gemini.com"
	sandbox    = "https://api.sandbox.gemini.com"
)

var (
	apikey    string
	apisecret string
	env       string
	freq      int
	iOffset    int
	pretty    bool
	repeat    bool
)