# Go Microservice

## REST  apis for common CRUD operations

### Tools
> stretchr/testify : For unit testing \
> gin-gonic/gin : Gin is a HTTP web framework written in Go for writing controller layer.

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
> curl -v localhost:9000/users?user_id=1234 