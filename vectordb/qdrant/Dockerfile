FROM golang
WORKDIR /app
COPY . .
RUN go build -o demo main.go

ENTRYPOINT ["./demo"]
