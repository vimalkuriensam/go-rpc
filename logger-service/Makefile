PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

log:
	protoc -Ipkg/proto --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. ./pkg/proto/logs.proto