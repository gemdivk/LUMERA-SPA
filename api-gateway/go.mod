module github.com/gemdivk/LUMERA-SPA/api-gateway

go 1.24.3

require (
	github.com/gemdivk/LUMERA-SPA/booking-service v0.0.0
	github.com/gemdivk/LUMERA-SPA/notification-service v0.0.0
	github.com/gemdivk/LUMERA-SPA/review-service v0.0.0
	github.com/gemdivk/LUMERA-SPA/user-service v0.0.0
	github.com/gin-gonic/gin v1.10.0
	google.golang.org/grpc v1.72.2
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/bytedance/sonic v1.13.2 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/gabriel-vasile/mimetype v1.4.9 // indirect
	github.com/gemdivk/LUMERA-SPA/salon-management-service v0.0.0-20250530081857-25ba8d5d4421
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.26.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.16.0 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gemdivk/LUMERA-SPA/review-service => ../review-service

replace github.com/gemdivk/LUMERA-SPA/user-service => ../user-service

replace github.com/gemdivk/LUMERA-SPA/booking-service => ../booking-service

replace github.com/gemdivk/LUMERA-SPA/notification-service => ../notification-service
