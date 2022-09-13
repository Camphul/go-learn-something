FROM golang:1-alpine as build
ARG TARGETARCH
WORKDIR /app
COPY go.* .
COPY main.go .
RUN GOARCH=${TARGETARCH} go build main.go
FROM alpine:latest
ARG TARGETARCH=amd64
WORKDIR /app
COPY --from=build /app/main /app/main

EXPOSE 8008
ENTRYPOINT ["./main"]