package exec

const (
	prodEnv    = "https://api.gemini.com"
	sandboxEnv = "https://api.sandbox.gemini.com"

	configFile = "config.ini"
)

var (
	apiKey    string
	apiSecret string
	env       string
	freq      int
	iOffset   int
	logFile   string
	pretty    bool
	repeat    bool
	baseURL   string
)
