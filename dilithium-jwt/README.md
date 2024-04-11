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

# Research Findings
+ My goal of this research was to find out create a JWT signature scheme for Dilithium & measure the following data:
    1. Memory usage difference between Dilithium, RSA, & ED25519
    2. Performance difference between Dilithium, RSA, & ED25519

### Token Generation Performance
| Algorithm | Time (µs) |
|-----------|-----------|
| RSA (2048)| 820.98    |
| Dilithium | 310.616   |
| ED25519   | 23.718    |

+ Dilithium performed 90% faster than RSA. Dilithium may be slower than ED25519 at token generation, but Dilithium beats ED25519 at token verification performance.

### Token Verification Performance
| Algorithm | Time (µs) |
|-----------|-----------|
| RSA       | 28.651    |
| Dilithium | 32.49     |
| ED25519   | 50.698    |

+ Dilithium performs 60% faster than ED25519 at token verification. Dilithium is slower than RSA at token verification, but it is only 13% slower. Having a faster token verification time is more important than token generation time because if you have 1 million concurrent users accessing a resource, the server will perform 1 million token verifications. Token generation only happens when the user's token expires.

### Memory Usage
| Algorithm | Verification (KB) | Generation (KB) | Mean (KB) |
|-----------|-------------------|-----------------|-----------|
| Dilithium | 283               | 301             | 292       |
| RSA       | 169               | 172             | 171       |
| ED25519   | 153               | 156             | 155       |

+ A memory usage measurement of the token generation and verification process was also done.
+ Dilithium uses 52% more heap memory compared to RSA and 61% more heap memory compared to ED25519 on average.
+ This is largely due to the fact that the binary length of of a Dilithium private key is very large compared to RSA or ED25519, so, the generated signature is also large.
+ Therefore, the CPU has to allocate memory on the heap both for the large signature and the large private key.
+ Even with a memory optimized implementation where we pass the private key and token by reference/pointers, the memory usage remains unchanged.

# Setup Instructions
+ The only pre-requisite to run my benchmark is `Docker`.
+ This documentation will only cover how to run this using Docker CLI, but you should also be able to run it using Docker Desktop.
+ This Docker image is hosted on Docker Hub at: [`meeranh/dilithium_jwt`](https://hub.docker.com/r/meeranh/dilithium_jwt)
+ Therefore, you can either pull this image straight from Docker Hub, or you can build the image manually.
1. **Pulling the Docker Image**
    ```bash
    docker pull meeranh/dilithium_jwt
    docker container run meeranh/dilithium_jwt
    ```
2. **Building the Docker Image**
    ```bash
    git clone http://gitlab.sliit.lk/r24-055/r24-055.git
    cd dilithium-jwt
    docker build -t dilithium_jwt .
    docker container run dilithium_jwt
    ```

# Output
+ Once you have run the container, you will be able to see the benchmarks to prove the novelty of my Dilithium JWT signature.
+ A sample output can be seen below:

```
Benchmarking memory usage...
Memory used for token generation (Dilithium)    : 301 KB
Memory used for token generation (ED25519)      : 156 KB
Memory used for token generation (RSA)          : 172 KB
Memory used for verification (Dilithium)        : 283 KB
Memory used for verification (ED25519)          : 153 KB
Memory used for verification (RS256)            : 169 KB
```

```
Benchmarking performance...
Time to generate a Dilithium JWT token  : 310.616µs
Time to verify a Dilithium JWT token    : 32.49µs

Time to generate an RSA JWT token       : 820.98µs
Time to verify an RSA JWT token         : 28.651µs

Time to generate a Ed25519 JWT token    : 23.718µs
Time to verify an Ed25519 JWT token     : 50.698µs
```

```
JWT TOKEN:
eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTI4ODM3NTcsImlhdCI6MTcxMjc5NzM1NywidXNlcm5hbWUiOiJIYXNzYW4gTS5NIn0.kRCp8RByxO61yzp3m8aGLt3qwiiHezNyeXTALpjf8_DNVSfCV6jnmmkZCJJo3QahYYWCk_7WcwUHq7hXrvmYAg
```
```json
{
  "exp": 1712770384,
  "iat": 1712683984,
  "username": "Hassan M.M"
}
```
