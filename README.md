# MiniHttpReMux

MiniHttpReMux is an HTTP proxy server written in Go that allows you to listen on multiple ports, modify HTTP headers, and forward requests to specified target servers. This project is designed to be easily configurable via a YAML file and can be deployed using Docker.

## Features

- Listen on multiple ports with individual configurations.
- Modify the Host header and other HTTP headers as specified in the configuration.
- Forward requests to designated target servers.
- Configurable using a YAML file for easy management.
- Docker support for easy deployment.

## Project Structure

```
MiniHttpReMux
├── main.go               # Entry point of the application     
├── internal
│   └── handlers
│       └── proxy.go      # Proxy logic and configuration
├── config
│   └── config.yaml       # Configuration file for the proxy
├── Dockerfile            # Dockerfile for building the image
├── go.mod                # Module definition and dependencies
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/dotnetautor/MiniHttpReMux.git
   cd MiniHttpReMux
   ```

2. Build the project:
   ```
   go build -o minihttremux
   ```

3. Run the application:
   ```
   ./minihttremux
   ```

4. Run with config file:
   ```
   ./minihttremux -config config/config.user.yaml
   ```` 

## Configuration

The proxy default configuration is defined in the `config/config.yaml` file. Each port can have its own target server and header modifications. An example configuration might look like this:

```yaml
ports:
  - port: 8080
    target: "http://example.com"
    headers:
      Host: "example.com"
      X-Custom-Header: "value"
  - port: 8081
    target: "http://another-example.com"
    headers:
      Host: "another-example.com"
```

## Docker

To build and run the Docker container, use the following commands:

1. Build the Docker image:
   ```
   docker build -t minihttremux .
   ```

2. Run the Docker container:
   ```
   docker run -p 8080:8080 -p 8081:8081 -v ./config.yaml:/app/config/config.yaml dotnetautor/mini-http-re-mux
   ```

## Docker Compose

You can also use Docker Compose to manage the container. Here is an example `docker-compose.yml` file:

```yaml
services:
  minihttremux:
    image: dotnetautor/mini-http-re-mux
    build: .
    ports:
      - "8080:8080"
      - "8081:8081"
    volumes:
      - ./config.yaml:/app/config/config.yaml
```

To build and run the Docker container using Docker Compose, use the following commands:

1. Build and start the Docker container:
```sh
docker-compose up --build
```

2. Stop the Docker container:
```sh
docker-compose down
```

## Usage

Once the server is running, you can send HTTP requests to the configured ports, and the proxy will handle the requests according to the specified configuration.

## License

This project is open source and available under the MIT License.