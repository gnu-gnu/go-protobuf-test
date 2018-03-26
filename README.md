# go-protobuf-test

simple message marshal and unmarshal example by go and protobuf

### to make your own protobuf message type

1) install go - refer https://golang.org/doc/install (set GOROOT and GOPATH environment variables)

2) download protocol buffers - refer https://github.com/google/protobuf/releases/tag/v3.5.1

3) install proto go library - ***go get -u github.com/golang/protobuf/protoc-gen-go*** (refer https://github.com/golang/protobuf)

4) if protoc cannot find protoc-gen-go, check environment path, %GOPATH%\bin

5) make your own message .proto file (refer https://developers.google.com/protocol-buffers/docs/proto3)

6) ***protoc --go_out=. <file_name>.proto*** and .go source file will be generated

7) modify source file's package declare following your project directory structure





