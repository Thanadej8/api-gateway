module github.com/Thanadej8/api-gateway

require github.com/Thanadej8/api-gateway/graph v0.0.1
replace github.com/Thanadej8/api-gateway/graph => ./graph

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/gin-gonic/gin v1.6.3
	github.com/vektah/gqlparser/v2 v2.1.0
)
