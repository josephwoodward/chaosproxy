package config

const DefaultTrafficRange = 100
const DefaultEnabled = true

type ConfigurationOptions struct {
	Config    Config
	Endpoints []Endpoint
}

type Config struct {
	Port    string
	Enabled *bool `yaml:"enabled,omitempty"`
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

func (d Config) IsEnabled() bool {
	if d.Enabled != nil {
		return *d.Enabled
	}

	return DefaultEnabled
}
