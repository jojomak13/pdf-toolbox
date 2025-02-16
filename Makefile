serve:
	@air

# Run on production mode
start:
	@go build -o bin/app && ./bin/app

# Install required packages
install:
	@go mod tidy
