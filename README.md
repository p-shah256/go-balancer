# Rust Reverse Proxy with Load Balancing

Welcome to the Rust Reverse Proxy with Load Balancing project! This is a Rust implementation of a reverse proxy designed to efficiently handle HTTP requests and distribute them among multiple backend servers for load balancing.

## Key Components

### 1. Reverse Proxy Engine

- **HTTP Server:** Utilizes Rust's `hyper` crate to implement a performant HTTP server, capable of handling incoming requests.

- **Request Parsing:** Parses incoming HTTP requests, extracting relevant information like method, headers, and target URL.

- **Response Handling:** Processes backend server responses and forwards them back to the client.

### 2. Load Balancing

- **Server Pool Management:** Maintains a pool of backend servers and efficiently distributes incoming requests among them.

- **Dynamic Weighting:** Allows dynamic adjustment of backend server weights based on configuration, enabling fine-grained control over load balancing.

- **Health Checks:** Implements a health-check mechanism to ensure only healthy backend servers receive traffic.

### 3. Configuration System

- **TOML Configuration:** Utilizes a TOML configuration file (`config.toml`) for easy customization of proxy settings, including backend server details and load balancing parameters.

### 4. Logging and Monitoring

- **Request Logging:** Logs detailed information about incoming requests, including method, URL, and response status.

- **Monitoring Endpoints:** Exposes monitoring endpoints for metrics retrieval, providing insights into proxy performance and health.

## Building and Running

### Prerequisites

- [Rust](https://www.rust-lang.org/tools/install): Ensure Rust is installed on your system.

### Build and Execution

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/reverse-proxy-rust.git
    cd reverse-proxy-rust
    ```

2. Build the project:

    ```bash
    cargo build --release
    ```

3. Customize the configuration file in `config.toml` to suit your environment and backend servers.

4. Run the reverse proxy application:

    ```bash
    ./target/release/reverse-proxy-rust
    ```

### Contributing

Contributions are welcome! If you encounter issues, have suggestions, or want to contribute new features, please open an issue or submit a pull request.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

