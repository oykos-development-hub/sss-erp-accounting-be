BINARY_NAME=accountingAPP

build:
	@echo "Building Accounting APP..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Accounting APP built!"

run: build
	@echo "Starting Accounting APP..."
	@./tmp/${BINARY_NAME}
	@echo "Accounting APP started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Accounting APP..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Accounting APP!"

restart: stop start