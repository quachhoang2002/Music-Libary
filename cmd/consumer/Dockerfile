# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.21.8-alpine as builder

# create a working directory inside the image
WORKDIR /app

# copy directory files i.e all files ending with .go
COPY . .

# download Go modules and dependencies
# RUN go mod download

# compile application
# /api: directory stores binaries file
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/api/main.go
RUN go build -o consumer ./cmd/consumer

##
## STEP 2 - DEPLOY
##
FROM alpine:latest
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /
COPY --from=builder /app/consumer .

EXPOSE 80

ENTRYPOINT ["./consumer"]