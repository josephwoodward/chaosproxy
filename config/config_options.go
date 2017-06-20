package config

const DefaultTrafficRange = 100

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
	ResponseStatusCode int  "responseStatusCode"
	Range              *int `yaml:"range,omitempty"`
}

func (d Endpoint) RangeOrDefault() int {
	if d.Range != nil {
		return *d.Range
	}

	return DefaultTrafficRange
}
