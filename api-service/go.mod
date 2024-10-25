module github.com/meeranh/register_service

go 1.23.0

require (
	github.com/cloudflare/circl v1.4.0
	github.com/google/uuid v1.6.0
	github.com/meeranh/register_service/base v0.0.0-00010101000000-000000000000
	github.com/meeranh/register_service/utils v0.0.0-00010101000000-000000000000
	gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5 v5.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/qdrant/go-client v1.11.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.16.1 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240311173647-c811ad7063a7 // indirect
	google.golang.org/grpc v1.62.1 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/meeranh/register_service/utils => ./utils

replace github.com/meeranh/register_service/base => ./algorithms

replace gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5 => ./algorithms/jwt
