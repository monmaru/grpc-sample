# grpc-sample

## Generate gRPC code
```
protoc -I proto/ proto/ght.proto --go_out=plugins=grpc:proto
```