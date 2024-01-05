gentool:
	go build -o ./tools ./cmd/protoc-gen-chat

genpb:
	rm -rf pkg/pb/message_rpc/*.ext.go pkg/pb/message_rpc/*.pb.go
	rm -rf internal/api/*.cli.go
	./tools/protoc --proto_path=protos/ --plugin=protoc-gen-go=./tools/protoc-gen-go --go_out=. protos/rpc/*
	./tools/protoc --proto_path=protos/ --plugin=protoc-gen-chat=./tools/protoc-gen-chat --chat_out=rpcext=true,rpccli="internal/api":. protos/rpc/*
	