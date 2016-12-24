# grpc-sample

## Generate gRPC code
```
protoc -I proto/ proto/ght.proto --go_out=plugins=grpc:proto
```
## Run Server
```
go run server/main.go
```
## Run Client
```
go run client/main.go -l Go
```