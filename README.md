# rest-api-golang

This branch contains a PRO REST API using golang.

## Build API
### Format code golang
    $ make format

### Build 

    $ make

### Binaries

    You can find binaries in folder bin for platform Linux, Mac OS X and Windows

### Run Unit Tests

    $ make test

### Run Integration tests

    $ make integrationtest

## Docker-compose
### Build containers

    $ docker-compose build

### Run infrastructure

    $ docker-compose up -d
    
    
### Down infrastructure

    $ docker-compose down
    
    
## Testing this API

### Root called

    $ curl -i -X GET http://localhost:9000/

### Stats for API

    $ curl -i -X GET http://localhost:9000/stats


### Numbers for API

    $ curl -i -X GET http://localhost:9000/numbers/<number>
    
 