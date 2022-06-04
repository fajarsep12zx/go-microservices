module go-microservices/core

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	google.golang.org/protobuf v1.28.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
