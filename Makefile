.PHONY: gen

gen:
	protoc --proto_path=. --twirp_out=. --go_out=. api/crisp.proto

test:
	go test -v ./...
