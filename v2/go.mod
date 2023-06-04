module github.com/yourusername/inventory-service

go 1.16

require (
	github.com/go-kratos/kratos/v2 v2.1.1
	github.com/google/wire v0.5.0
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/go-kratos/kratos/v2 => github.com/go-kratos/kratos/v2 v2.1.1
)
