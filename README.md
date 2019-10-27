#Go Microservice

## REST  apis for common CRUD operations

###Tools
> stretchr/testify : For unit testing 

### Steps to create build
> cd mvc/ \
Build: go build -o mvc \
Run: ./mvc 

### To execute unit tests
>cd mvc/domains \
 go test -cover 

### Curl to execute GET api
> curl -v localhost:9000/users?user_id=1234 