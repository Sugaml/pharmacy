proto:
	protoc --go_out=. --go-grpc_out=. proto/patient/patient.proto

run:
	go run main.go

