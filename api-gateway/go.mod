module api-gateway

go 1.21.6

require (
	github.com/go-playground/validator/v10 v10.22.0
	github.com/gorilla/mux v1.8.1
	github.com/rs/cors v1.11.0
	google.golang.org/grpc v1.65.0
	user-service v1.0.1
)

require (
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace user-service v1.0.1 => ../user-service
