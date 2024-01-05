gentool:
	go build -a -o ./tools ./cmd/protoc-gen-chat

genpb:
	rm -rf pkg/pb/message_cli/*.pb.go pkg/pb/message_cli/*.ext.go
	rm -rf pkg/pb/message_rpc/*.pb.go pkg/pb/message_rpc/*.ext.go
	rm -rf internal/api/*.cli.go
	./tools/protoc --proto_path=protos/ --plugin=protoc-gen-go=./tools/protoc-gen-go --go_out=. protos/cli/* protos/rpc/*
	./tools/protoc --proto_path=protos/ --plugin=protoc-gen-chat=./tools/protoc-gen-chat --chat_out=cli=true:. protos/cli/*
	./tools/protoc --proto_path=protos/ --plugin=protoc-gen-chat=./tools/protoc-gen-chat --chat_out=cli=true,rpcext=true,rpccli="internal/api",rpcinter="pkg/inter/rpcinter":. protos/rpc/*
	