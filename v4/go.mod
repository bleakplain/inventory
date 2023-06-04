module github.com/yourusername/inventory-service

go 1.16

require (
	github.com/go-kratos/kratos/v2 v2.1.1
	github.com/go-kratos/kratos-layout v0.0.0-20211010065956-4c6f06b3aaf8
	github.com/google/wire v0.6.0
	github.com/jinzhu/gorm v1.9.16
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/go-kratos/kratos-layout => github.com/go-kratos/kratos-layout v0.0.0-20211010065956-4c6f06b3aaf8
