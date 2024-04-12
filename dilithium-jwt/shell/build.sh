#!/bin/bash

# Build the benchmarks
echo "Building the benchmarks..."
go build benchmarks/dilithium_keygen.go && \
	go build benchmarks/dilithium_verify.go && \
	go build benchmarks/rsa_keygen.go && \
	go build benchmarks/rsa_verify.go && \
	go build benchmarks/ed25519_verify.go && \
	go build benchmarks/ed25519_keygen.go && \
	go build benchmarks/benchmark.go
