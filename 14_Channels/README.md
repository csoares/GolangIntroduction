# Go Channels Examples

This directory contains examples demonstrating different patterns and use cases for Go channels.

## Directory Structure

1. `01_basic/`: Basic channel operations including:
   - Synchronous communication
   - Channel closing
   - Range over channels

2. `02_buffered/`: Buffered channels demonstrating:
   - Asynchronous communication
   - Channel capacity
   - Producer-consumer pattern

3. `03_pipeline/`: Channel pipeline pattern showing:
   - Data transformation flow
   - Generator pattern
   - Multiple stage processing

4. `04_worker_pool/`: Worker pool pattern illustrating:
   - Concurrent task processing
   - Load distribution
   - Result collection

## Running the Examples

To run any example, navigate to its directory and use:

```bash
go run main.go
```

Each example is self-contained and demonstrates a specific channel concept.