package configuration

type ConfigurationOptions struct {
	Config    Config
	Endpoints []Endpoint
}

type Config struct {
}

type Endpoint struct {
	Url                string
	Host               string
	Delay              int
	ResponseStatusCode int
}
