# Folder Structure
```
keys            *(Example public/private keys used for benchmarking)*
├── dilithium
├── ed25519
├── jwt
└── rsa

base            *(The folder containing my JWT library & helper functions)
└── jwt         *(My custom fork of the JWT library with my novel Dilithium signature)*

.
├── benchmarks  *(This folder contains the benchmarking code to prove my Dilithium signature's novelty)*
└── shell       *(Some shell files that help me with creating a Docker image/container)*
```

# Usage Instructions
+ The only pre-requisite to run this fork and my benchmarks is `Docker`.
+ This documentation will only cover how to run this using Docker CLI, but you should also be able to run it using Docker Desktop.
+ This Docker image is hosted on Docker Hub at: 
