module github.com/hw1513/ecomm

go 1.24.2

replace (
	github.com/hw1513/ecomm/internal/common => ./internal/common
	github.com/hw1513/ecomm/internal/order => ./internal/order
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
)
