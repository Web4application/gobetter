cd $(go list -f '{{ .Dir }}' github.com/google/certificate-transparency-go); \
go install github.com/golang/mock/mockgen; \
go install google.golang.org/protobuf/proto; \
go install google.golang.org/protobuf/cmd/protoc-gen-go; \
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc; \
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc; \
go install golang.org/x/tools/cmd/stringer
go install github.com/web4application/gobetter@latest
