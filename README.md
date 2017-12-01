# rest-api-golang

This branch contains a simple REST API using golang.

## RUN API

    $ go run api.go 

## Build and Run API
### BUILD

    $ go build api.go
    
### RUN

    $ ./api
    
## Testing this API

### Root called

    $ curl -i -X GET http://localhost:9000/

### Stats for API

    $ curl -i -X GET http://localhost:9000/stats
 