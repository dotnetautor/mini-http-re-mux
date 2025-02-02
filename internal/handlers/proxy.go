package proxy

import (
	"io"
	"net/http"
)

// ProxyConfig holds the configuration for the proxy server.
type ProxyConfig struct {
	Port       int               `yaml:"port"`
	Target     string            `yaml:"target"`
	Headers    map[string]string `yaml:"headers"`
}

// Proxy represents the HTTP proxy.
type Proxy struct {
	config ProxyConfig
}

// NewProxy creates a new Proxy instance with the given configuration.
func NewProxy(config ProxyConfig) *Proxy {
	return &Proxy{config: config}
}

// ServeHTTP handles incoming HTTP requests and forwards them to the target server.
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Modify the Host header and other headers as per the configuration
	r.Host = p.config.Target
	for key, value := range p.config.Headers {
		r.Header.Set(key, value)
	}

	// Forward the request to the target server
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, p.config.Target+r.URL.String(), r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers and status code to the original response
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}