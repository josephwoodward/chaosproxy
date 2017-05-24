package configuration

type ConfigurationOptions struct {
	Config    Config
	Endpoints []Endpoint
}

type Config struct {
	Port string
}

type Endpoint struct {
	Host               string
	Url                string
	Delay              int
	ResponseStatusCode int "responseStatusCode"
}
