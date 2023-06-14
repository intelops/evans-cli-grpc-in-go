path:
	export PATH="$PATH:$(go env GOPATH)/bin"

protoc:
	rm -rf gen && mkdir gen && protoc --go_out=./gen --go-grpc_out=./gen api/v1/*.proto