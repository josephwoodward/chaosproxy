package configuration

type ConfigurationOptions struct {
	Config    Config
	Endpoints []Endpoint
}

type Config struct {
}

type Endpoint struct {
	Path  string
	Delay int
}
