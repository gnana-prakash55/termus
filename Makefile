BIN=termus

test:
	@echo "Testing Application..."
	@go test

build:
	@echo "Building Application..."
	@go build -o ./bin/${BIN}

run:
	@echo "Running Application..."
	@./bin/${BIN}