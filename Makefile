include .env
export
BINARY=engine
## Run the application
run-api:
	@echo "Running the application"
	@go run cmd/api/main.go

## Run the consumer
run-consumer:
	@echo "Running the consumer"
	@go run cmd/consumer/main.go

## Run the scheduler
run-scheduler:
	@echo "Running the scheduler"
	@go run cmd/scheduler/main.go

swagger:
	@echo "Generating swagger for application"
	@swag init -g cmd/api/main.go


engine:
	go build -o ${BINARY} app/*.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

build:
	docker-compose build

run:
	docker-compose up -d --force-recreate

stop:
	docker-compose down


.PHONY: clean install unittest build docker run stop vendor lint-prepare lint