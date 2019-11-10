# Go Microservice

## REST  apis for common CRUD operations

### Tools
> stretchr/testify : For unit testing \
> gin-gonic/gin : Gin is a HTTP web framework written in Go for writing controller layer. \
> h2non/gock : HTTP traffic mocking for writing integration tests \
> logrus: Logging for Go \
> uber-go/zap: Blazing fast, structured, leveled logging in Go.

### Steps to create build
> cd mvc/ \
Build: go build -o mvc \
Run: ./mvc 

### To execute unit tests
>cd mvc/domains \
 go test -cover 

### To execute benchmarks tests
>cd mvc/utils \
 go test -bench=.

### Curl to execute GET api
> curl -v localhost:9000/users/1234 \
> curl -v localhost:9000/users/1234 -H "Accept:application/xml"

### Curl to Create and GET OAuth Tokens
> curl localhost:8080/oauth/access_token/ABCD1I \
> curl -X POST "localhost:8080/oauth/access_token" -d '{"username":"fede", "password":"testing"}' \
> curl localhost:8080/oauth/access_token/USR_123