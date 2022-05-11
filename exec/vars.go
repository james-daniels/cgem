package exec

const (
	production = "https://api.gemini.com"
	sandbox    = "https://api.sandbox.gemini.com"

	configFile = "config.ini"
)

var (
	apikey    string
	apisecret string
	env       string
	freq      int
	iOffset   int
	logfile  string
	pretty    bool
	repeat    bool
)

