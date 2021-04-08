package traefik_routing_plugin

// Config the plugin configuration.
type Config struct {
	Routes map[string]string `json:"routes,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Routes: make(map[string]string),
	}
}
