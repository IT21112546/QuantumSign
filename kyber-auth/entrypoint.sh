#!/bin/bash

# Print initial options
echo "Do you want to run the client or server?"
echo "1. Client"
echo "2. Server"

# Read user input
read -p "Enter your choice (1 or 2): " choice

# Process the user input
case $choice in
    1)
        printf "Starting the client...\n\n"
        ./run_client
        ;;
    2)
        printf "Starting the server...\n\n"
        ./run_server
        ;;
    *)
        printf "Invalid choice\n\n"
        exit 1
        ;;
esac
