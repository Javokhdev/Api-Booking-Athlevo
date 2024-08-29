
swag-init:
	swag init -g api/router.go -o api/docs 

prot-exp:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"

gen-proto:
	protoc --go_out=./ \
    --go-grpc_out=./ \
	submodule-for-timecapsule/timeline_service/*.proto

	protoc --go_out=./ \
    --go-grpc_out=./ \
	submodule-for-timecapsule/memory_service/*.proto