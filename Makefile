.PHONY: generate

# Generate Go code from spec
generate:
	oapi-codegen -generate types -o internal/gen/types.gen.go -package gen api/openapi.yaml
	oapi-codegen -generate chi-server -o internal/gen/server.gen.go -package gen api/openapi.yaml

# Convert OpenAPI YAML → JSON for Insomnia
openapi-json:
	yq -o=json eval api/openapi.yaml > api/openapi.json
