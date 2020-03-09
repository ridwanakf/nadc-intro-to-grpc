# command for generating protobuf code
proto:
	@echo " >> generating protobuf code"
	@protoc --proto_path=protos/ --proto_path=/usr/local/include/ --go_out=plugins=grpc:protos protos/*.proto

# go build command for rpc server
build:
	@echo " >> building binaries for nadc-intro-to-grpc"
	@go build -v -o nadc-intro-to-grpc app.go

# go run command for rpc server
run: build
	./nadc-intro-to-grpc
