# Dockerfile

# Step 1: build

FROM golang:1.24.3 AS build

WORKDIR /backend
COPY . .

RUN go mod tidy
RUN go build -o focusio-back main.go

# Step 2: execution
FROM debian:ubuntu:bulleye-slim

RUN apt update && apt install -y build-essential

WORKDIR /backend
COPY --from=build /backend/focusio-back .

EXPOSE 8080

CMD ["./focusio-back"]

