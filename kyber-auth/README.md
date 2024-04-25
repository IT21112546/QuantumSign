# Folder Structure

```
.
├── client      (The client program to interact with the server)
├── server      (The server program to interact with the client)
└── proto       (The protobuf definitions for both client and server)
```

# Usage Instructions
+ Build and run the Docker image using the following command:
```bash
docker image build -t kyber-auth .
docker container run -it --network host kyber-auth /bin/bash
```
+ First run the server, which will then listen for connections infinitely until you close the connection.
+ In a separate terminal, run the same container with the same command:
```bash
docker container run -it --network host kyber-auth /bin/bash
```
+ In this container, select the client option if you want to run a demo mutual authenticated challenge-response instance.
+ Else, if you want to benchmark the speeds, select the benchmark option.
+ For both options, specify `public.key` as the path to the public key. This is a demo key located in this folder.

# Methodology
+ The goal of this component was to come up with a faster mutual authentication challenge-response mechanism which is faster than modern mutual authentication methods based on RSA.
+ The current methodologies revolve around the following logic:
    1. Client initiates a connection to the server with a challenge to the server (CHALL1).
    2. Server responds to the challenge (CHALL1) with a new challenge to the client (CHALL2).
    3. Because the server responded to CHALL1, the user can be sure that the server is authentic.
    4. The client now responds to CHALL2 and sends the response to the server.
    5. Based on the client's response, the server can be sure that the client is authentic.
    6. Then, a Diffie-Hellman key exchange is performed to establish a shared symmetric key for encryption and decryption.

+ In total, the above process requires 3 messages + 1 Diffie-Hellman key exchange. This is a total of 5 messages back-and-forth for both mutual authentication and key exchange.
+ Also, Diffie-Hellman key exchange is not a post-quantum secure algorithm. So, it is not a good idea to use it in the future.
+ The implementation in this component is based on the following logic:
    1. Client sends his public key `PUBKEY` + a challenge to the server `CHALL1`.
        + The challenge sent to the server `CHALL1` is a random string encrypted with the server's public key.
    2. The server decrypts `CHALL1` using his private key, which reveals the shared secret `SS`.
    3. The server has now verified that the client is authentic.
    4. The server then sends data encrypted using the shared secret `SS`, which the client decrypts using the shared secret `SS`.
    5. At this point, both the client and the server have verified each other's authenticity.

+ The methodology we have proposed requires only 2 messages back-and-forth for both mutual authentication and key exchange, which is an 85% reduction of total communication.
+ Our authentication is also based on Google's Remote Procedure Call (RPC) system.
+ gRPC is one of the fastest RPC systems available today, it directly invokes a function on the server from the client.

# Research Findings (CPU & RAM Usage)
+ We tested our methodology on a local network and found that our methodology completes a mutual authentication challenge-response in 142 microseconds with a standard deviation of 20 microseconds.
+ For perspective, [a paper from Sami A. Nagar and Saad Alshamma](www.setit.rnu.tn/final/P86951046.pdf) demonstrates that their fastest implementation of RSA has a decryption time of ~425ms.
+ Another research paper from [Priyadarshini Patil et al.](https://www.sciencedirect.com/science/article/pii/S1877050916001101/pdf?md5=3ba721d629b6a4c93cc91abb712ee32c&pid=1-s2.0-S1877050916001101-main.pdf) also studied the performance effects of RSA and concluded that fastest decryption time of RSA was ~390ms and the fastest encryption time of RSA falls within ~500ms, which goes well with Sami A. Nagar and Saad Alshamma's findings.
+ Another interesting metric is that Priyadarshini Patil et al. demonstrated that RSA uses 14KB of memory.
+ Our implementation allocated a total of 77MB over 10,000 iterations.
+ Which means that on average, our implementation uses 7.7KB of memory per iteration.

| Algorithm | CPU Time | Memory Usage |
|-----------|----------|--------------|
| RSA       | 425ms    | 14KB         |
| Kyber     | 142μs    | 7.7KB        |

+ This is a further 58% reduction in memory usage compared to RSA.
+ All of these performance tests are reproducible via our benchmark programs containerized using Docker.


+ **Note**: The benchmarks were generated on an HP Victus 15 with an AMD Ryzen 7 7535HS (8 Cores, 16 Threads, 3.3 GHz Base), 16GB DDR4 memory, Arch Linux (6.8.4-arch1-1 kernel).
