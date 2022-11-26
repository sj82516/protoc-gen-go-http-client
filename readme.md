# protoc-go-gen-http-client
only for practice. Gen http client api call from protobuf.

## what can it do
refer to example/test.proto.   
Gen http client call in golang by proto tags.

```protobuf

```

```go
package main

func GetClient() {
    res, err := http.Get("api.com")
    
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()


}
```

## How to dev
### 1. install at local
rebuild and install bin
> $ go install

test if works.
> $ protoc --go-http-client_out=. --go-http-client_opt="paths=source_relative" --go_out=. --go_opt=paths=source_relative example/proto/*.proto

add opts from command.
- base_url: the base url of http request
> $ protoc --go-http-client_out=. --go-http-client_opt=paths=source_relative,base_url=api.com example/*.proto
## Ref
1. https://rotemtam.com/2021/03/22/creating-a-protoc-plugin-to-gen-go-code/