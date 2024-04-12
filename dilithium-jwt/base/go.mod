module gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base

go 1.22.2

replace gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5 => ./jwt

require (
	github.com/cloudflare/circl v1.3.7
	gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5 v5.0.0-00010101000000-000000000000
)

require golang.org/x/sys v0.15.0 // indirect
