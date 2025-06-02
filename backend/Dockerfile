# Dockerfile

# Step 1: build

FROM golang:1.22 AS build

WORKDIR /backend
COPY . .

RUN go mod tidy
RUN go build -o focusio-back main.go

# Step 2: execution
FROM debian:bullseye-slim

WORKDIR /backend
COPY --from=build /backend/focusio-back .

EXPOSE 8080

CMD ["./focusio-back"]

