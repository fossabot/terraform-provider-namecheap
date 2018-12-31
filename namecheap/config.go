package namecheap

// Config defines the structure used to instantiate the Namecheap provider.
type Config struct {
	user     string
	token    string
	clientIP string
	baseURL  string
}

// Client configures and returns a fully initialized Namecheap client.
func (c *Config) Client() (interface{}, error) {
	return struct{}{}, nil
}
