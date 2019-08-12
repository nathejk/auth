module api

go 1.12

replace local/token => ./token

require (
	github.com/SermoDigital/jose v0.9.2-0.20180104203859-803625baeddc
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	local/token v0.0.0-00010101000000-000000000000
)
