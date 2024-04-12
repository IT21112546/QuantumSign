# A multi-stage build to create the binaries
FROM golang
WORKDIR /usr/src/app
COPY . .
RUN chmod +x shell/*.sh && ./shell/build.sh

# We wil be running the benchmarks in an Ubuntu container
FROM ubuntu
WORKDIR /usr/src/app
COPY --from=0 /usr/src/app .

ENTRYPOINT ["./shell/entrypoint.sh"]
