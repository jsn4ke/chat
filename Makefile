gentool:
	go build -a -o ./tools ./cmd/protoc-gen-chat

genpb:
	rm -rf pbk/pb/message_obj/*.pb.go
	rm -rf pkg/pb/message_cli/*.pb.go pkg/pb/message_cli/*.ext.go
	rm -rf pkg/pb/message_rpc/*.pb.go pkg/pb/message_rpc/*.ext.go
	rm -rf internal/inter/rpcinter/*.inter.go
	rm -rf internal/api/*.cli.go
	rm -rf internal/rpc/*.svr.go
	./tools/protoc --proto_path=protos/ --plugin=protoc-gen-go=./tools/protoc-gen-go \
		--go_out=paths=source_relative:./pkg/pb/ \
		protos/message_cli/* protos/message_rpc/* protos/message_obj/*
	./tools/protoc --proto_path=protos/ --plugin=protoc-gen-chat=./tools/protoc-gen-chat \
		--chat_out=cli="pkg/pb":. protos/message_cli/*
	./tools/protoc --proto_path=protos/ --plugin=protoc-gen-chat=./tools/protoc-gen-chat \
		--chat_out=rpcext="pkg/pb",rpccli="internal/api",rpcinter="internal/inter/rpcinter",rpcsvr=internal/rpc:. protos/message_rpc/*
	