# ENV
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-linux-x86_64.zip
unzip protoc-3.14.0-linux-x86_64.zip
mv bin/protoc /usr/local/bin/

# Step
goctl api -o hermes.api
cd api
goctl api go -api hermes.api -dir .
cd rpc/transform
goctl rpc template -o transform.proto
goctl rpc protoc transform.proto --go_out=. --go-grpc_out=. --zrpc_out=.
cd rpc/transform/model
goctl model mysql ddl -c -src shorturl.sql -dir .