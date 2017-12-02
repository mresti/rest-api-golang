# rest-api-golang
This repo contains code to learn "Build a REST API with Golang".

## Code

The code is splitted in two branch to learn little by little.

### Simple Example

In the branch [#simple_api](https://github.com/mresti/rest-api-golang/tree/simple_api) you can find a simple file in golang.

### Complex Example

In the branch [#professional_api](https://github.com/mresti/rest-api-golang/tree/professional_api)you can find the API with below contains:

- Scaffolding this project as Golang project.
- Use a Makefile to tasks for this project as build, format code, unit tests, integration tests, etc. (Warning: This makefile set Go env as GOPATH...)
- Use a dependency: [Gorilla mux](https://github.com/gorilla/mux) instead of standard http router.
- Travis.yml as CI tool.
- Dockerfile to build this API in a container using Docker.
- docker-compose file to use the Dockerfile.
