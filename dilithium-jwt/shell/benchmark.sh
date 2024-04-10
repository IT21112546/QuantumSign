#!/bin/bash

echo ""
echo "Benchmarking memory usage..."
./dilithium_verify && ./ed25519_verify && ./rsa_verify && ./dilithium_keygen && ./ed25519_keygen && ./rsa_keygen

echo ""
echo "Benchmarking performance..."
./benchmark
