# go microservices Backend

> Backend API service for DMAA using go-micro, rpc

## Prerequisites

* [Go](https://golang.org/doc/install) v1.16
* [Protobuf](https://developers.google.com/protocol-buffers/docs/downloads) v3.x.x (or via `brew install protobuf`)
* [protoc-gen-go](https://pkg.go.dev/mod/github.com/golang/protobuf@v1.4.2) v1.4.2
```
  GIT_TAG="v1.4.2"
  go get -d -u github.com/golang/protobuf/protoc-gen-go
  go install github.com/golang/protobuf/protoc-gen-go
```

## Quick Config

* Enter directory core and generate proto file
```
  cd core && ./protoc.sh
```
* Install all package go mod using mod tidy and vendor


* Generate graphql schema for apps
```
  cd api \
  && go run github.com/99designs/gqlgen generate
```

* Generate graphql schema for login
```
  cd api/login \
  && go run github.com/99designs/gqlgen generate
```

* Copy and setup env config
```
  cp api/.env.example api/.env \
  && cp core/config/.app.config.json.example core/config/.app.config.json
```

* Copy all config .env to all services
```
  cp services/.env.example services/<ServiceName>/.env
```

## Run server graphql services
```
  cd api \
  && go run server/server.go
```

## Run example services
```
    cd services/<ServiceName> \
    && go run server/server.go 
```
