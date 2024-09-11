module github.com/DataDog/datadog-api-client-go/v2

go 1.19

retract (
	// Version used to retract v2.0.0 and v2.0.1. DO NOT USE.
	v2.0.1
	// Premature major version v2 release. DO NOT USE.
	v2.0.0
)

require (
	github.com/DataDog/zstd v1.5.2
	github.com/bytedance/sonic v1.12.2
	github.com/goccy/go-json v0.10.2
	github.com/google/uuid v1.5.0
	golang.org/x/oauth2 v0.10.0
)

require (
	github.com/bytedance/sonic/loader v0.2.0 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/net v0.17.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
