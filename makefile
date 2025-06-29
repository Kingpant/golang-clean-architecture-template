.PHONY: swag swag-fmt swag-init

# Format all swag comments
swag-fmt:
	swag fmt

# Generate Swagger docs
swag-init:
	swag init --generalInfo cmd/api/main.go --parseDependency --parseInternal

# Run both
swag: swag-fmt swag-init
