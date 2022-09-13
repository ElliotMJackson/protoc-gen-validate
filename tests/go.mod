module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	golang.org/x/net v0.0.0-20220907135653-1e95f45603a7
	google.golang.org/protobuf v1.28.1
)

replace github.com/envoyproxy/protoc-gen-validate => ../
