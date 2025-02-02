package main

import (
    "flag"
    "fmt"
    "os"
    "log"
    "net/http"
    
    "gopkg.in/yaml.v2"
    "MiniHttpReMux/internal/handlers"
)

type Config struct {
    Ports []proxy.ProxyConfig `yaml:"ports"`
}

func main() {
    // Define a command-line flag for the config file path
    configFile := flag.String("config", "config/config.yaml", "Path to the config file")
    flag.Parse()
    
    // Read the config file
    data, err := os.ReadFile(*configFile)
    if err != nil {
        log.Fatalf("Failed to read config file: %v", err)
    }

    // Parse the config file
    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("Failed to parse config file: %v", err)
    }

    // Use the parsed configuration
    for _, proxyConfig := range config.Ports {
        go func(proxyConfig proxy.ProxyConfig) {
            log.Printf("Starting proxy on port %d targeting %s", proxyConfig.Port, proxyConfig.Target)
            // Erstellen einer neuen Proxy-Instanz
            p := proxy.NewProxy(proxyConfig)

            // Erstellen eines neuen ServeMux für jeden Proxy
            mux := http.NewServeMux()
            mux.HandleFunc("/", p.ServeHTTP)

            // Starten eines HTTP-Servers mit dem Proxy als Handler
            server := &http.Server{
                Addr:    fmt.Sprintf(":%d", proxyConfig.Port),
                Handler: mux,
            }
            if err := server.ListenAndServe(); err != nil {
                log.Fatalf("Failed to start server on port %d: %v", proxyConfig.Port, err)
            }
        }(proxyConfig)
    }

    // Keep the main goroutine alive
    select {}
}