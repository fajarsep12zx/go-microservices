module zebrax.id/product/dmaa/services/user

go 1.13

require (
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	zebrax.id/product/dmaa/core v0.0.0
)

replace zebrax.id/product/dmaa/core => ../../core

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
