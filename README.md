# Go Concurrent Port Scanner
A high-performance, concurrent TCP port scanner built with Go. 
This tool demonstrates advanced concurrency patterns including Pipelines, Fan-out (distributing work across multiple CPU cores), and Fan-in (multiplexing results back into a single stream).

# 🚀 Features

* **Generator Pattern**: Efficiently streams port numbers without pre-allocating large slices.

* **Adaptive Concurrency**: Automatically detects the number of available CPU cores to spin up an optimal number of worker goroutines.

* **Graceful Shutdown**: Uses a done channel to prevent goroutine leaks and ensure clean exits.

* **Pipeline Architecture**: Decouples port generation, scanning logic, and result aggregation.


# 🛠 Project Structure
* **port/port.go**: Contains the core logic for scanning individual ports and the pipeline stage functions.

* **main.go**: Orchestrates the scanner, manages the worker pool, and aggregates results.

# 📋 How it Works

1. Generation: GeneratePortNumber creates a stream of integers representing the port range.

2. Fan-out: The main function spins up multiple ScanPortStream workers (one per CPU core). These workers all consume from the same portNumbersStream.

3. Scanning: Each worker attempts a TCP connection with a 10-second timeout.

4. Fan-in: The fanIn function collects the "Open Port" results from all workers and merges them into a single output channel for printing.


# Running the scanner

1. Clone the repository or navigate to the project directory.

2. Run the following command:`go run main.go`