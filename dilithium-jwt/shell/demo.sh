#!/bin/bash

echo ""
echo "JWT TOKEN:"
cat ./keys/jwt/ed25519

echo ""
echo ""

# Check if jq is installed and use if installed
if ! [ -x "$(command -v jq)" ]; then
	cat ./keys/jwt/decoded.json
else
	cat ./keys/jwt/decoded.json | jq
fi
