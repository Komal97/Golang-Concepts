gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:proto

clean:
	rm proto/*.go

run-server:
	go run server/main.go

run-client:
	go run client/main.go