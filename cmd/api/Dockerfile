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

# compile application
# /api: directory stores binaries file
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/api/main.go
RUN go build -o api ./cmd/api
RUN mkdir music

##
## STEP 2 - DEPLOY
##
FROM alpine:latest

# Install ca-certificates and timezone
RUN apk update
RUN apk add --no-cache ca-certificates tzdata && update-ca-certificates

# Set the timezone.
ENV TZ=Asia/Ho_Chi_Minh
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
COPY --from=builder /app/api .

EXPOSE 80

ENTRYPOINT ["./api"]
