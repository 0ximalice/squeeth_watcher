# Squeeth Watcher & Simulator

Spawn a Squeeth Watcher, watch the pool prices and provide a simulation tool that simulate trades.

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/doc/install)
- [Python](https://www.python.org/downloads/)

### Building and Running

1. Clone the repository:

   ```bash
   git clone https://github.com/0ximalice/squeeth_watcher.git
   cd squeeth_watcher
   ```

2. Build and run the Docker container:

   ```bash
   docker build -t squeeth-watcher .
   docker run -p 8501:8501 -v ./data:/app/data -name squeeth-watcher squeeth-watcher
   ```

This command will spawn a Squeeth watcher, and Steamlit app to simulate the pool.
The pool data needed for the simulation is stored in the `data` folder. It will be taked some time to generate the data.

## Usage

Wait for the container to start and open the following URL in your browser:

```ssh
http://localhost:8501/
```
