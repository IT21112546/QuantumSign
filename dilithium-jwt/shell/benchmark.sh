#!/bin/bash

echo ""
echo "Benchmarking memory usage..."
./dilithium_keygen && ./ed25519_keygen && ./rsa_keygen && ./dilithium_verify && ./ed25519_verify && ./rsa_verify

echo ""
echo "Benchmarking performance..."
./benchmark
