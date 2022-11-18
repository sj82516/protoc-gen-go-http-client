# protoc-go-gen-http-client
only for practice. Gen http client api call from protobuf.

## How to dev
### 1. install at local
rebuild and install bin
> $ go install

test if works.
> $ protoc --go-http-client_out=. --go-http-client_opt="paths=source_relative" example/*.proto

## Ref
1. https://rotemtam.com/2021/03/22/creating-a-protoc-plugin-to-gen-go-code/