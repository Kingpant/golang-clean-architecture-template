.PHONY: swag swag-fmt swag-init mocks

# Format all swag comments
swag-fmt:
	swag fmt

# Generate Swagger docs
swag-init:
	swag init --generalInfo cmd/api/main.go --parseDependency --parseInternal

# Run both
swag: swag-fmt swag-init

# Generate mocks for user repository
mocks: 
	go generate ./...
