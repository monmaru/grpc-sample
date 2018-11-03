
all:
	echo "## Go stubs and gateway ##"
	protoc \
		-I /usr/local/include -I . \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		proto/ght.proto
	protoc \
		-I /usr/local/include -I . \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		proto/ght.proto
	protoc \
		-I /usr/local/include -I . \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		proto/ght.proto

clean:
	rm proto/*.pb.go
	rm proto/*.pb.gw.go
	rm proto/*.swagger.json