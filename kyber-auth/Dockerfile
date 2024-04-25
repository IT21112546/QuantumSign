FROM golang
WORKDIR /app
COPY . .
RUN go build -o run_server server/main.go
RUN go build -o run_client client/main.go

ENTRYPOINT ["./entrypoint.sh"]
