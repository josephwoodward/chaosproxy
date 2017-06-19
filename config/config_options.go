package config

type ConfigurationOptions struct {
	Config    Config
	Endpoints []Endpoint
}

type Config struct {
	Port    string
	Enabled bool
}

type Endpoint struct {
	Host               string
	Url                string
	Delay              int
	ResponseStatusCode int "responseStatusCode"
	Range              int
}
