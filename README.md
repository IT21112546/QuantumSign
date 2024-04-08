# Repository Structure
**benchmarking branch**     *For PP1*
```
benchmarking
   ├── dilithium-jwt        *(Handled by IT21145766)*
   ├── kyber-auth           *(Handled by IT21025808)*
   ├── ml                   *(Handled by IT21146824)*
   └── vectordb             *(Handled by IT21112546)*
```

**qsign branch**            *For PP2 & Final Viva*
```
qsign
   ├── anomaly-detection    *(Handled by IT21146824)*
   ├── federated-sso        *(Handled by IT21145766)*
   ├── key-retrieval        *(Handled by IT21112546)*
   └── oauth2               *(Handled by IT21025808)*
```

**NOTE:** Before PP1, all our code changes will be done on the *benchmarking* branch and a PR will be made to *master*. In this branch, our primary focus is to prove that the implementations we built perform better than the existing implementations, and therefore, we primarily focus on benchmarking our implementations before PP1. These benchmarks are what will be included in our research paper.
After PP1, we will create a new branch called *qsign* in which we will be building our commercials software that will be presented in PP2 and the final viva. PRs will be made to *master* from the *qsign* branch.
Each branch in the *qsign* branch will have four micro-services that will be built by each individual, and deployed to a K8s cluster.

# Introduction
+ The most widely used public-key cryptosystem known as RSA is deprecated as it is vulnerable to quantum computer brute-force attempts. Post-quantum algorithms exist, but they are based on lattice-based cryptography, which is different from Elliptic Curve Cryptography (ECC) and RSA which the world was used to. Crystals-Kyber/Dilithium is a promising cryptosystem based on lattice-based cryptography, which is the core of our research.

### Research Core Objectives
+ Our research at it's core tries to accomplish is two things:
    1. Discover the cost of transitioning a serverless application from ECC to lattice-based cryptography.
    2. Improve the speed of authentication & authorization services via novel methods.

### Individual Research Questions
##### IT21025808's Question
+ **How much more (or less) would it cost to migrate a serverless application using RSA to Crystals-Kyber?**
    + We believe this is an important question because companies may be reluctant to migrate to algorithms like Crystals-Kyber as there is no research done demonstrating the infrastructure cost differences.

##### IT21145766's Question
+ **How much infrastructure resource difference would transitioning from RS256 digital signatures (JWT) to Dilithium digital signatures make?**
    + Since Dilithium is quantum-safe and is potentially the future of token signatures, it makes sense to do a research and identify the infrastructure cost differences.

##### IT21112546's Question
+ **How faster can key retrieval be if we migrated the key storage mechanism to a vector database?**
    + Crystals-Kyber's public and private keys are vectors. Therefore, we have an opportunity to speed up key storage and retrieval speeds by using vector similarity searches.

##### IT21146824's Question
+ **How much of a performance drawback would integrating an anomaly detection model into a federated SSO system have?**
    + Anomaly detection models are computationally expensive. We need to research how much of a performance hit an SSO system will take if a real-time model is integrated.

### Research Objectives Per Individual

| IT Number  | Research Objective | Completed? |
|------------|--------------------|------------|
| IT21145766 | Calculate the infrastructure cost difference between RSA and Dilithium. | No |
|            | Implement a JWT signing scheme using Crystals-Dilithium. | No |
|            | Perform a benchmark to measure the performance difference between RSA and Dilithium. | No |
|            | Implement a commercial token-granting server using Dilithium signed JWTs. | No |
|            | Develop a commercial federated SSO system that supports concurrent users. | No |
| IT21025808 | Implement an authentication system using Crystals-Kyber. | No |
|            | Benchmark the infrastructure resource usage of RSA and Crystals-Kyber. | No |
|            | Connect the authentication service to the token-granting server. | No |
| IT21112546 | Implement a post-quantum key storage system using a vector database. | No |
|            | Benchmark the key retrieval speed of the vector database. | No |
|            | Compare it to speeds of traditional key storage systems. | No |
|            | Combine the key storage system with the authentication service. | No |
| IT21146824 | Implement an anomaly detection model with an increased accuracy score. | No |
|            | Attach the model to the authentication service and compare performance hits. | No |
