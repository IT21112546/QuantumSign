# Folder Structure
```
.
├── qdrant  (Our novel public key storage system)
└── mysql   (The benchmarks for storing public keys in MySQL)
```

# Usage Instructions
+ `Qdrant` and `MySQL` should be launched through `docker-compose` to run the benchmarks.
```bash
# Execute in this folder
docker compose up --detach
```

+ To run a demo of our vector database based public-key storage/retrieval functions, navigate to the `qdrant` folder, build the Docker image, and run the Docker container interactively.
```bash
# Execute in qdrant folder
docker image build -t vectordb/qdrant .
docker container run -it --network host vectordb/qdrant /bin/bash
```

+ To replicate the MySQL benchmarks, navigate to the `mysql` folder, build the Docker image, and run the Docker container similar to the previous command.
```bash
docker image build . -t vectordb/mysql
docker container run --network host vectordb/mysql
```

# Methodology
+ Our goal of this research was to come up with a convenient public key storage system using vector databases.
+ Our hypotheses was that since Crystal-Kyber public keys are vectors, we can leverage the high speeds of vector databases to store and retrieve them.
+ Our methodology was the following:
    1. Build the needed functions to store vectors on [Qdrant](https://github.com/qdrant/qdrant).
    2. Simulate storing 100,000 database entries and measure their speed. A single entry contained:
        1. A randomly generated Crystal-Kyber public key vector.
        2. A randomly generated username.
        3. A randomly generated email.
    3. Contrast their speeds with [this](https://ieeexplore.ieee.org/document/8898035) research paper.
    4. Finally, do the same benchmark on MySQL for 100,000 database entries and compare the speed.
        + Since MySQL does not support vector storage (which our public keys are), we have a separate methodology to measure MySQL performance.

### MySQL Vector Storage Performance
+ MySQL does not support storing vectors as a data type.
+ In our case, the public keys were arrays of unsigned 32-bit integers.
+ Therefore, we had to convert these unsigned 32-bit integers to a single base64 string.
+ A `Benchmark` table was created with the following columns:
    1. `id` (Primary Key)
    2. `username` (VARCHAR)
    3. `email` (VARCHAR)
    4. `public_key` (TEXT)
+ The base64 string was stored in the `public_key` column.
+ This was iterated 100,000 times and the average time taken to write was measured.
+ Also, for every 10,000 entries, one entry was randomly selected and its time taken to read was measured.

# Research Findings
### Read Operation Speeds
+ We will label our implementation as `Qdrant` and the MySQL implementation as `MySQL`.
+ The read operation speeds for `Qdrant` were significantly faster than `MySQL`.
+ This is because vector similarity search is a core feature of Qdrant, which is not present in MySQL.

| Database Entry Count | Qdrant (Indexed) Read Speed (ms) | MySQL Read Speed (ms) |
|----------------------|----------------------------------|-----------------------|
| 10,000               | 0.638                            | 6.3393                |
| 20,000               | 1.162                            | 13.0808               |
| 30,000               | 0.626                            | 41.3334               |
| 40,000               | 2.963                            | 203.8643              |
| 50,000               | 0.716                            | 79.7105               |
| 60,000               | 3.972                            | 348.3405              |
| 70,000               | 0.792                            | 119.2560              |
| 80,000               | 3.584                            | 156.8944              |
| 90,000               | 0.986                            | 159.5426              |

+ We also performed a non-indexed benchmark of our implementation of Qdrant versus MySQL.

| Database Entry Count | Qdrant (Without Indexing) Read Speed (ms) | MySQL Read Speed (ms) |
|----------------------|-------------------------------------------|-----------------------|
| 10,000               | 0.663                                     | 6.3393                |
| 20,000               | 1.656                                     | 13.0808               |
| 30,000               | 0.682                                     | 41.3334               |
| 40,000               | 3.070                                     | 203.8643              |
| 50,000               | 0.747                                     | 79.7105               |
| 60,000               | 4.118                                     | 348.3405              |
| 70,000               | 0.813                                     | 119.2560              |
| 80,000               | 4.919                                     | 156.8944              |
| 90,000               | 0.796                                     | 159.5426              |

+ There does not seem to be a big speed increase when indexing is used for our vector database.
+ However, our public key storage system is still significantly faster than MySQL.
+ The non-indexed version does not show any speed improvement because we indexed the following keys:
    1. `username`
    2. `email`
+ There was no specific way to index the public key, as it is already like the primary key of the database.
+ Therefore, the reason why there was no speed improvement is likely because the public key was not indexed.

### Write Operation Speeds
+ The write operation speeds for `Qdrant` were significantly faster than `MySQL`.
+ The following is the time taken to store data on our implementation of Qdrant vs MySQL

| Database Entry Count | Qdrant (Indexed) Write Speed (ms) | MySQL Write Speed (ms) |
|----------------------|----------------------------------|-------------------------|
| 10,000               | 0.062                            | 0.3384                  |
| 20,000               | 0.079                            | 0.3334                  |
| 30,000               | 0.084                            | 0.3338                  |
| 40,000               | 0.071                            | 0.3336                  |
| 50,000               | 0.091                            | 0.3329                  |
| 60,000               | 0.088                            | 0.3314                  |
| 70,000               | 0.093                            | 0.3305                  |
| 80,000               | 0.095                            | 0.3311                  |
| 90,000               | 0.105                            | 0.3317                  |

| Database Entry Count | Qdrant (Without Indexing) Write Speed (ms) | MySQL Write Speed (ms) |
|----------------------|-------------------------------------------|-------------------------|
| 10,000               | 0.060                                     | 0.3384                  |
| 20,000               | 0.077                                     | 0.3334                  |
| 30,000               | 0.071                                     | 0.3338                  |
| 40,000               | 0.080                                     | 0.3336                  |
| 50,000               | 0.087                                     | 0.3329                  |
| 60,000               | 0.085                                     | 0.3314                  |
| 70,000               | 0.089                                     | 0.3305                  |
| 80,000               | 0.094                                     | 0.3311                  |
| 90,000               | 0.100                                     | 0.3317                  |

## Memory & CPU Differences
+ The methodology used to measure RAM and CPU usage is as follows:
    + We containerized MySQL and Qdrant into Docker images.
    + We performed the read and write operations on both databases.
    + Using `docker stats`, we were able to retrieve the precise memory and CPU usage of the containers.

### Memory Usage
#### Starting Memory Usage
1. **MySQL**                : 2.45% ± 0%    (371.7 MiB / 14.83 GiB)
2. **Qdrant (Indexed)**     : 1.51% ± 0.03% (230.5 MiB / 14.83 GiB)
3. **Qdrant (Non-Indexed)** : 1.5% ± 0.00%  (228.5 MiB / 14.83 GiB)

#### Maximum Memory Usage
1. **MySQL**                : 3.35% ± 0%    (508.9 MiB / 14.83 GiB)
2. **Qdrant (Indexed)**     : 8.35% ± 0.15% (1.783 GiB / 14.83 GiB)
3. **Qdrant (Non-Indexed)** : 6.34% ± 1.5%  (1.089 GiB / 14.83 GiB)

+ It is important to note the following facts:
    1. **MySQL** stayed at a constant memory usage of 3.35% after benchmark execution (no reduction).
    2. **Qdrant (Indexed)** stayed at a constant memory usage of 8.35% after benchmark execution.
    3. **Qdrant (Non-Indexed)** dropped down to 3.34% after benchmark execution.

+ The above facts suggest that:
    1. Qdrant (Indexed) uses 113% more memory than MySQL.
    2. Qdrant (Non-Indexed) uses 74% more memory than MySQL.
    3. MySQL does not free up memory after benchmark execution
    4. Qdrant (Non-Indexed) frees up memory after benchmark execution.

### CPU Usage
#### Starting CPU Usage
1. **MySQL**                : 0.39% ± 0.05%
2. **Qdrant (Indexed)**     : 0.10% ± 0.04%
3. **Qdrant (Non-Indexed)** : 0.10% ± 0.04%

#### Maximum CPU Usage
1. **MySQL**                : 50.6% ± 5.2%
2. **Qdrant (Indexed)**     : 993% ± 31%
3. **Qdrant (Non-Indexed)** : 988% ± 23%

+ **Note**: The benchmarks were generated on an HP Victus 15 with an AMD Ryzen 7 7535HS (8 Cores, 16 Threads, 3.3 GHz Base), 16GB DDR4 memory, Arch Linux (6.8.4-arch1-1 kernel).
