.PHONY: all basic intermediate advanced clean fmt test

# Basic Examples
basic: hello vars control funcs arrays maps structs pointers methods

hello:
	go run 01_HelloWorld/main.go

vars:
	go run 02_Variables/main.go

control:
	go run 03_ControlStructures/main.go

funcs:
	go run 04_Functions/main.go

arrays:
	go run 05_Arrays_Slices/main.go

maps:
	go run 06_Maps/main.go

structs:
	go run 07_Structs/main.go

pointers:
	go run 08_Pointers/main.go

methods:
	go run 09_Methods_Interfaces/main.go

# Intermediate Examples
intermediate: packages errors fileio

packages:
	cd 10_Packages_Modules && go run main.go

errors:
	go run 11_ErrorHandling/main.go

fileio:
	go run 12_FileIO/main.go

# Advanced Concurrency Examples
advanced: concurrency channels fanin ratelimit context semaphores

concurrency:
	go run 13_Concurrency/main.go

channels:
	go run 14_Channels/01_basic/main.go
	fanin:
	go run 15_AdvancedConcurrency/01_FanInFanOut/main.go

ratelimit:
	go run 15_AdvancedConcurrency/02_RateLimiting/main.go

context:
	go run 15_AdvancedConcurrency/03_ContextCancellation/main.go

semaphores:
	go run 15_AdvancedConcurrency/04_Semaphores/main.go

# Common Operations
fmt:
	go fmt ./...

test:
	go test ./...

clean:
	go clean

# Default target
all: basic intermediate advanced