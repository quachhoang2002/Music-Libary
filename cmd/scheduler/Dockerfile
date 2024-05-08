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
RUN go build -o scheduler ./cmd/scheduler

##
## STEP 2 - DEPLOY
##
FROM alpine:latest
RUN apk add --no-cache ca-certificates && update-ca-certificates

# Set the timezone.
ENV TZ=Asia/Ho_Chi_Minh
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /
COPY --from=builder /app/scheduler .

EXPOSE 80

ENTRYPOINT ["./scheduler"]
