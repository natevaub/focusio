# Step 1: Build statically linked Go binary
FROM golang:1.24.3 AS build

WORKDIR /backend

COPY . .

# Static build without glib c
# CGO_ENABLED=0: disables CGO (C language interoperability) which forces Go to compile a fully static binary without any external C library dependencies like glibc
# GOOS=linux: tells Go to build the binary for Linux OS
# GOARCH=amd64: tells Go to build the binary for 64-bit x86 architecture

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o focusio-back main.go

FROM scratch

COPY --from=build /backend/focusio-back .

EXPOSE 8080

CMD ["./focusio-back"]


