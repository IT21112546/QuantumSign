module gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt

go 1.22.2

replace (
	gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base => ./base
	gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5 => ./base/jwt
)

require (
	gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base v0.0.0-00010101000000-000000000000
	gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5 v5.0.0-00010101000000-000000000000
)

require (
	github.com/cloudflare/circl v1.3.7 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
